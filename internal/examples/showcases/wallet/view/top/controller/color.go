package controller

import (
	"github.com/0xDezzy/cutego/core"

	"github.com/0xDezzy/cutego/internal/examples/showcases/wallet/theme/controller"
)

type colorController struct {
	core.QObject

	_ func() `signal:"change,auto"`
}

//lazy binding to the (qml singleton) theme controller
func (c *colorController) change() { controller.Controller.Change() }
