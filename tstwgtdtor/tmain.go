package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	log.Println("Enter main...")

	// Create application
	if runtime.GOOS == "android" {
		os.Setenv("QT_AUTO_SCREEN_SCALE_FACTOR ", "1.5")
		qtcore.QCoreApplication_SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true)
	} else {
		// qtrt.SetFinalizerObjectFilter(finalizerFilter)
	}
	qtrt.SetDebugFFICall(true)
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	if false {
		app.SetAttribute(qtcore.Qt__AA_EnableHighDpiScaling, true) // for android
	}
	uiqc := make(chan func(), 1)
	mech := NewNotifier(func() {
		for len(uiqc) > 0 {
			remfn := <-uiqc
			remfn()
		}
	})
	_ = mech
	qtrt.FinalProxyFn = func(f func()) {
		uiqc <- f
		mech.Trigger()
		// f()
	}

	items := []*Ui_ContactItemView{}
	mw := NewUi_MainWindow2()
	mw.QWidget_PTR().Show()

	qtrt.Connect(mw.PushButton_2, "clicked(bool)", func(bool) { runtime.GC() })
	qtrt.Connect(mw.PushButton_3, "clicked(bool)", func(bool) {
		for i := 0; i < 100; i++ {
			item := NewUi_ContactItemView2()
			mw.VerticalLayout_9.InsertWidget(0, item.QWidget_PTR(), 0, 0)
			items = append(items, item)
		}
	})
	qtrt.Connect(mw.PushButton_4, "clicked(bool)", func(bool) {
		for _, item := range items {
			mw.VerticalLayout_9.RemoveWidget(item.QWidget_PTR())
			// item.QWidget_PTR().DeleteLater()
			qtwidgets.DeleteQWidget(item.QWidget_PTR())
		}
		items = items[:0]
	})

	go func() {
		time.Sleep(2 * time.Second)
		// app.Exit(0)
	}()
	// Execute app
	app.Exec()
}
