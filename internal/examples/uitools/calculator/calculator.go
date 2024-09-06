package main

import (
	"os"

	"github.com/0xDezzy/cutego/widgets"

	"github.com/0xDezzy/cutego/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
