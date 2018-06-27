package main

import (
	"gopp"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"unsafe"

	"github.com/3d0c/gmf"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/qtchina/qt.go.demos/qtx"
)

var mw *Ui_MainWindow

var meth *qtx.Notifier
var stopped = true
var recfname string

func main() {
	qtx.AppRunW(func() {
		mw = NewUi_MainWindow2()
		recfname = mw.ComboBox.CurrentText()

		qtrt.Connect(mw.PushButton, "clicked(bool)", func(bool) {
			if !stopped {
				log.Println("Is running")
			} else {
				stopped = false
				go recproc(recfname)
			}
		})
		qtrt.Connect(mw.PushButton_2, "clicked(bool)", func(bool) {
			stopped = true
		})
		qtrt.Connect(mw.PushButton_3, "clicked(bool)", func(bool) {
			flt := "MPEG (*.mpg *.mpeg *.mkv *.mp4);; WEBM (*.webm);;"
			fname := qtwidgets.QFileDialog_GetOpenFileName(mw.QWidget_PTR(), "Select a video file/dev:", os.Getenv("HOME"), flt, flt, 0)
			if fname != "" {
				recfname = fname
				mw.ComboBox.SetCurrentText(fname)
			}
		})
		qtrt.Connect(mw.ComboBox, "currentTextChanged(const QString&)", func(s string) {
			recfname = s
		})
		mw.MainWindow.Show()

		pointerStep := func(p unsafe.Pointer, offset uintptr) unsafe.Pointer {
			return unsafe.Pointer(uintptr(p) + offset)
		}
		// for both Qt4 and Qt5
		NewQPainter := func(w qtwidgets.QWidget_ITF) *qtgui.QPainter {
			ptr := pointerStep(w.QWidget_PTR().GetCthis(), 2*unsafe.Sizeof(uintptr(0)))
			ptdev := qtgui.NewQPaintDeviceFromPointer(ptr)
			return qtgui.NewQPainter_1(ptdev)
		}

		mw.Widget.InheritPaintEvent(func(event *qtgui.QPaintEvent) {
			p := NewQPainter(mw.Widget)
			// p.FillRect_2(p.Viewport(), qtgui.NewQBrush_1(0))
			for len(imgdc) > 0 {
				d := <-imgdc
				idptr := unsafe.Pointer(&d.b[0])
				// imgo := qtgui.NewQImage_5_(idptr, d.w, d.h, d.ls, qtgui.QImage__Format_RGB32) // why black?
				imgo := qtgui.NewQImage_3_(idptr, d.w, d.h, qtgui.QImage__Format_RGB888) // why ok?
				p.DrawImage_5(p.Viewport(), imgo)
			}
			qtgui.DeleteQPainter(p)
		})

		meth = qtx.NewNotifier(func() {
			mw.Widget.Update()
		})

	})
}

var imgdc = make(chan *imginfo, 128)

type imginfo struct {
	b  []byte
	w  int
	h  int
	ls int
}

func recproc(fname string) {
	iptctx := gmf.NewCtx()
	// err := iptctx.OpenInput("/usr/share/texmf-dist/tex/latex/mwe/example-movie.mp4")
	// err := iptctx.OpenInput("/dev/video0")
	err := iptctx.OpenInput(fname)
	gopp.ErrPrint(err)
	defer iptctx.CloseInputAndRelease()
	log.Println(iptctx)
	log.Println(iptctx.StreamsCnt())
	// stm, err := iptctx.GetStream(0)
	stm, err := iptctx.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	gopp.ErrPrint(err)
	log.Println(stm)
	log.Println(stm.Id(), stm.Index(), stm.IsAudio(), stm.IsVideo(), stm.NbFrames(), stm.Type(), stm.IsCodecCtxSet())

	iptcctx := stm.CodecCtx()
	log.Println(iptcctx.Codec().Id(), gmf.AV_CODEC_ID_RAWVIDEO)

	dstcctx, swsctx, dstFrame := makeRawFrame(stm)
	_ = dstcctx

	// tmpcctx := makeRawCodec(stm) == dstcctx
	srcpkti := 0
	for !stopped {
		srcpkt := <-iptctx.GetNewPackets()
		srcpkti++
		frmi := 0
		ok := true
		for frm := range srcpkt.Frames(iptcctx) {
			swsctx.Scale(frm, dstFrame)
			pkt, err := dstFrame.Encode(dstcctx)
			gopp.ErrPrint(err, ok, frm)
			log.Println(srcpkti, frmi, ok, pkt.StreamIndex(), stm.Index())
			if pkt.StreamIndex() != stm.Index() {
				log.Println("steam index not match,", pkt.StreamIndex(), stm.Index())
				pkt.Free()
				pkt.Release()
				continue
			}
			log.Println(srcpkti, frmi, ok, err, "channels:", frm.Channels(), "format:", frm.Format(),
				"keyframe:", frm.KeyFrame(), "width:", frm.Width(), "height:", frm.Height(),
				"timestamp:", frm.TimeStamp(), "size:", pkt.Size())
			log.Println(srcpkti, frmi, ok, err, "datalen:", len(pkt.Data()), "linesize:", frm.LineSize(0), frm.LineSize(1))

			if true {
				dat := pkt.Data()
				// dat = dstFrame.RawData()
				// dat = frm.RawData()
				imgdc <- &imginfo{dat, dstFrame.Width(), dstFrame.Height(), dstFrame.LineSize(0)}
				meth.Trigger()
			}

			frm.Free()
			frm.Release()
			pkt.Free()
			pkt.Release()
		}
		srcpkt.Free()
		srcpkt.Release()
		// time.Sleep(10 * time.Millisecond)
	}
	log.Println("done")
}

func makeRawImage(data []byte, width, height int) {
	// width, height := frame.Width(), frame.Height()
	img := new(image.RGBA)
	img.Pix = data
	img.Stride = 4 * width // 4 bytes per pixel (RGBA), width times per row
	img.Rect = image.Rect(0, 0, width, height)

	writeFile(img)
}

var imgi int = 0

func writeFile(b image.Image) {
	name := "./tmp/virec" + strconv.Itoa(imgi) + ".jpg"

	fp, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(fp, b, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatal(err)
	}

	if err := fp.Close(); err != nil {
		log.Fatal(err)
	}
	imgi++
}

func makeRawCodec(srcvstm *gmf.Stream) *gmf.CodecCtx {
	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		log.Fatal(err)
	}

	cc := gmf.NewCodecCtx(codec)
	// defer gmf.Release(cc)

	cc.SetTimeBase(gmf.AVR{Num: 1, Den: 1})

	// This should really be AV_PIX_FMT_RGB32, but then red and blue channels are switched
	cc.SetPixFmt(gmf.AV_PIX_FMT_RGB24).SetWidth(srcvstm.CodecCtx().Width()).SetHeight(srcvstm.CodecCtx().Height())
	if codec.IsExperimental() {
		cc.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := cc.Open(nil); err != nil {
		log.Fatal(err)
	}
	return cc
}

func makeRawFrame(srcvstm *gmf.Stream) (*gmf.CodecCtx, *gmf.SwsCtx, *gmf.Frame) {
	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		log.Fatal(err)
	}

	cc := gmf.NewCodecCtx(codec)
	// defer gmf.Release(cc)

	cc.SetTimeBase(gmf.AVR{Num: 1, Den: 1})

	// This should really be AV_PIX_FMT_RGB32, but then red and blue channels are switched
	cc.SetPixFmt(gmf.AV_PIX_FMT_RGB24).SetWidth(srcvstm.CodecCtx().Width()).SetHeight(srcvstm.CodecCtx().Height())
	if codec.IsExperimental() {
		cc.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := cc.Open(nil); err != nil {
		log.Fatal(err)
	}

	swsCtx := gmf.NewSwsCtx(srcvstm.CodecCtx(), cc, gmf.SWS_BICUBIC)
	// defer gmf.Release(swsCtx)

	dstFrame := gmf.NewFrame().
		SetWidth(srcvstm.CodecCtx().Width()).
		SetHeight(srcvstm.CodecCtx().Height()).
		SetFormat(gmf.AV_PIX_FMT_RGB24) // see above

	if err := dstFrame.ImgAlloc(); err != nil {
		log.Fatal(err)
	}
	return cc, swsCtx, dstFrame
}
