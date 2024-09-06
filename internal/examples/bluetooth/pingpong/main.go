package main

import (
	"os"

	"github.com/0xDezzy/cutego/core"
	"github.com/0xDezzy/cutego/gui"
	"github.com/0xDezzy/cutego/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	core.QLoggingCategory_SetFilterRules("qt.bluetooth* = true")
	gui.NewQGuiApplication(len(os.Args), os.Args)

	pingPong := NewPingPong(nil)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("pingPong", pingPong)
	engine.Load(core.NewQUrl3("qrc:/assets/main.qml", 0))

	gui.QGuiApplication_Exec()
}
