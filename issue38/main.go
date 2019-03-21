package main

import (
	"github.com/kitech/qt.go/qtwidgets"
	"github.com/qtchina/qt.go.demos/qtx"
)

func main() {

	qtx.AppRunW2(func() interface{} {
		mw := qtwidgets.NewQMainWindowp()
		mw.Show()

		path := "."
		fileFialog := qtwidgets.NewQFileDialog1(mw, "openfile", path, "*.*")
		str := fileFialog.GetOpenFileNamep()

		return []interface{}{fileFialog, str}
	})
}
