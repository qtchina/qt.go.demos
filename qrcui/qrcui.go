package main

import (
	"image/png"

	"gopp"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/tuotoo/qrcode"
)

func createqrc(data string) string {
	qrco, err := qr.Encode(data, qr.M, qr.Auto)
	gopp.ErrPrint(err)
	log.Println(qrco)
	qrco, err = barcode.Scale(qrco, 200, 200)

	file, err := os.Create("qrc1.png")
	gopp.ErrPrint(err)
	log.Println(file.Name())
	defer file.Close()

	err = png.Encode(file, qrco)
	gopp.ErrPrint(err)

	return "qrc1.png"
}

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	// pm := qtgui.NewQPixmap_3_("/home/gzleo/toxrecvs/avatar.png")
	// lab := qtwidgets.NewQLabel__()
	// lab.SetPixmap(pm)
	// lab.Show()

	mw := NewUi_MainWindow2()
	mw.MainWindow.Show()

	qtrt.Connect(mw.PushButton, "clicked()", func() {
		log.Println("hehehe")
		mw.PlainTextEdit_2.Clear()

		dat := mw.PlainTextEdit.Document().ToPlainText()
		fname := createqrc(dat)

		pm := qtgui.NewQPixmap_3_(fname)
		mw.Label_3.SetPixmap(pm)
	})

	qtrt.Connect(mw.PushButton_2, "clicked()", func() {
		file, err := os.Open("qrc1.png")
		gopp.ErrPrint(err, file)
		qrcm, err := qrcode.Decode(file)
		gopp.ErrPrint(err)
		mw.PlainTextEdit_2.SetPlainText(qrcm.Content)
	})

	app.Exec()
}
