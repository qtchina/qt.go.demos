package main

import (
	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/qtchina/qt.go.demos/qtx"
)

func main() {
	if true {
		test_q1()
	} else {
		test_q2()
	}
}

func test_q1() {
	var l *qtwidgets.QLabel
	qtx.AppRunW(func() {
		l = qtwidgets.NewQLabel1p("hehehe")

		l.SetTextFormat(qtcore.Qt__RichText)
		// strText1 := "<html><body> <img src=\"https://csdnimg.cn/pubfooter/images/csdn-kf.png\" width=\"100\" height=\"100\" /> </body> </html>"
		strText1 := "<html><body> <img src=\"https://upload-images.jianshu.io/upload_images/5310144-8cb3d72f8747c946.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/300/h/240\" width=\"100\" height=\"100\" /> </body> </html>"
		l.SetText(strText1)

		l.Show()
	})
}

func test_q2() {
	var l *qtwidgets.QLabel
	var pMovie *qtgui.QMovie

	qtx.AppRunW(func() {
		l = qtwidgets.NewQLabel1p("hehehe")

		// pMovie = pMovie.NewForInherit2p("d:\\zb.gif")
		pMovie = pMovie.NewForInherit2p("/home/gzleo/Downloads/loading.gif")
		l.SetMovie(pMovie)
		l.SetFixedSize1(125, 125)
		l.SetScaledContents(true)
		pMovie.Start()

		l.Show()
	})
}
