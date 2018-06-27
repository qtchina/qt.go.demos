package qtx

import (
	"os"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func AppRunW(appfn func()) {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	appfn()
	app.Exec()
}

func AppRunG(appfn func()) {
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)
	appfn()
	app.Exec()
}

func AppRunC(appfn func()) {
	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	appfn()
	app.Exec()
}
