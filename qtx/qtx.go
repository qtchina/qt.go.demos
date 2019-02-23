package qtx

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtgui"
	"github.com/kitech/qt.go/qtwidgets"
)

func AppRunW(appfn func()) { AppRunW2(appfn) }
func AppRunG(appfn func()) { AppRunG2(appfn) }
func AppRunC(appfn func()) { AppRunC2(appfn) }

// hold appfn returned value's reference for not auto dtor
func AppRunW2(appfn interface{}) {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	rets, err := runappfn(appfn)
	if err != nil {
		log.Println(err, rets)
		return
	}
	app.Exec()
}

func AppRunG2(appfn interface{}) {
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)
	rets, err := runappfn(appfn)
	if err != nil {
		log.Println(err, rets)
		return
	}
	app.Exec()
}

func AppRunC2(appfn interface{}) {
	app := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)
	rets, err := runappfn(appfn)
	if err != nil {
		log.Println(err, rets)
		return
	}
	app.Exec()
}

// appfn any count return value(s), no argument func
func runappfn(appfn interface{}) ([]interface{}, error) {
	appfnty := reflect.TypeOf(appfn)
	if appfnty.Kind() != reflect.Func {
		return nil, fmt.Errorf("Invalid argument, not func")
	}

	appfnval := reflect.ValueOf(appfn)
	outs := appfnval.Call(nil)

	var rets []interface{}
	for _, out := range outs {
		rets = append(rets, out.Interface())
	}
	return rets, nil
}
