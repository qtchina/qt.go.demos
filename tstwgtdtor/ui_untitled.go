package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"
import "github.com/kitech/qt.go/qtrt"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }
func init() { qtrt.KeepMe() }

//  header block end

//  struct block begin
func NewUi_MainWindow() *Ui_MainWindow {
	return &Ui_MainWindow{}
}

type Ui_MainWindow struct {
	Centralwidget            *qtwidgets.QWidget
	VerticalLayout           *qtwidgets.QVBoxLayout
	HorizontalLayout         *qtwidgets.QHBoxLayout
	PushButton               *qtwidgets.QPushButton
	PushButton_2             *qtwidgets.QPushButton
	PushButton_3             *qtwidgets.QPushButton
	PushButton_4             *qtwidgets.QPushButton
	ScrollArea               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents *qtwidgets.QWidget
	VerticalLayout_9         *qtwidgets.QVBoxLayout
	Menubar                  *qtwidgets.QMenuBar
	Statusbar                *qtwidgets.QStatusBar
	MainWindow               *qtwidgets.QMainWindow
	SizePolicy               *qtwidgets.QSizePolicy
	SizePolicy1              *qtwidgets.QSizePolicy
}

//  struct block end

//  setupUi block begin
// void setupUi(QMainWindow *MainWindow)
func (this *Ui_MainWindow) SetupUi(MainWindow *qtwidgets.QMainWindow) {
	this.MainWindow = MainWindow
	// { // 126
	if MainWindow.ObjectName() == "" {
		MainWindow.SetObjectName("MainWindow")
	}
	MainWindow.Resize(393, 458)
	this.Centralwidget = qtwidgets.NewQWidget(this.MainWindow, 0)        // 111
	this.Centralwidget.SetObjectName("Centralwidget")                    // 112
	this.VerticalLayout = qtwidgets.NewQVBoxLayout_1(this.Centralwidget) // 111
	this.VerticalLayout.SetObjectName("VerticalLayout")                  // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()                   // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout")              // 112
	this.PushButton = qtwidgets.NewQPushButton(this.Centralwidget)       // 111
	this.PushButton.SetObjectName("PushButton")                          // 112

	this.HorizontalLayout.Layout().AddWidget(this.PushButton) // 115

	this.PushButton_2 = qtwidgets.NewQPushButton(this.Centralwidget) // 111
	this.PushButton_2.SetObjectName("PushButton_2")                  // 112

	this.HorizontalLayout.Layout().AddWidget(this.PushButton_2) // 115

	this.PushButton_3 = qtwidgets.NewQPushButton(this.Centralwidget) // 111
	this.PushButton_3.SetObjectName("PushButton_3")                  // 112

	this.HorizontalLayout.Layout().AddWidget(this.PushButton_3) // 115

	this.PushButton_4 = qtwidgets.NewQPushButton(this.Centralwidget) // 111
	this.PushButton_4.SetObjectName("PushButton_4")                  // 112

	this.HorizontalLayout.Layout().AddWidget(this.PushButton_4) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0) // 115

	this.ScrollArea = qtwidgets.NewQScrollArea(this.Centralwidget) // 111
	this.ScrollArea.SetObjectName("ScrollArea")                    // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Expanding, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                            // 114
	this.SizePolicy.SetVerticalStretch(0)                                                              // 114
	this.SizePolicy.SetHeightForWidth(this.ScrollArea.SizePolicy().HasHeightForWidth())                // 114
	this.ScrollArea.SetSizePolicy(this.SizePolicy)                                                     // 114
	this.ScrollArea.SetFocusPolicy(qtcore.Qt__NoFocus)                                                 // 114
	this.ScrollArea.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)                        // 114
	this.ScrollArea.SetWidgetResizable(true)                                                           // 114
	this.ScrollArea.SetAlignment(qtcore.Qt__AlignLeading | qtcore.Qt__AlignLeft | qtcore.Qt__AlignTop) // 114
	this.ScrollAreaWidgetContents = qtwidgets.NewQWidget(nil, 0)                                       // 111
	this.ScrollAreaWidgetContents.SetObjectName("ScrollAreaWidgetContents")                            // 112
	this.ScrollAreaWidgetContents.SetGeometry(0, 0, 373, 16)                                           // 114
	this.SizePolicy1 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy1.SetHorizontalStretch(0)                                                           // 114
	this.SizePolicy1.SetVerticalStretch(0)                                                             // 114
	this.SizePolicy1.SetHeightForWidth(this.ScrollAreaWidgetContents.SizePolicy().HasHeightForWidth()) // 114
	this.ScrollAreaWidgetContents.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.VerticalLayout_9 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents)                  // 111
	this.VerticalLayout_9.SetSpacing(0)                                                                // 114
	this.VerticalLayout_9.SetObjectName("VerticalLayout_9")                                            // 112
	this.VerticalLayout_9.SetContentsMargins(0, 0, 0, 0)                                               // 114
	this.ScrollArea.SetWidget(this.ScrollAreaWidgetContents)                                           // 114

	this.VerticalLayout.Layout().AddWidget(this.ScrollArea) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget)      // 114
	this.Menubar = qtwidgets.NewQMenuBar(this.MainWindow)     // 111
	this.Menubar.SetObjectName("Menubar")                     // 112
	this.Menubar.SetGeometry(0, 0, 393, 21)                   // 114
	this.MainWindow.SetMenuBar(this.Menubar)                  // 114
	this.Statusbar = qtwidgets.NewQStatusBar(this.MainWindow) // 111
	this.Statusbar.SetObjectName("Statusbar")                 // 112
	this.MainWindow.SetStatusBar(this.Statusbar)              // 114

	this.RetranslateUi(MainWindow)

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	// noimpl: {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "MainWindow", "dummy123", 0))
	this.PushButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "PushButton", "dummy123", 0))
	this.PushButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "PushButton", "dummy123", 0))
	this.PushButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Add Items", "dummy123", 0))
	this.PushButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Clear Items", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_MainWindow2() *Ui_MainWindow {
	this := &Ui_MainWindow{}
	w := qtwidgets.NewQMainWindow(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_MainWindow) QWidget_PTR() *qtwidgets.QWidget {
	return this.MainWindow.QWidget_PTR()
}

//  done block end
