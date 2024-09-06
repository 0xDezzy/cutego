package main

import (
	"os"

	"github.com/0xDezzy/cutego/core"
	"github.com/0xDezzy/cutego/gui"
	"github.com/0xDezzy/cutego/quick"

	_ "github.com/0xDezzy/cutego/internal/examples/qml/extending/components/test_module_2/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.Engine().AddImportPath("qrc:/qml")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
