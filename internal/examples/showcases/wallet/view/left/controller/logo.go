package controller

import (
	"github.com/0xDezzy/cutego/core"
	"github.com/0xDezzy/cutego/gui"
)

type logoController struct {
	core.QObject

	_ func() `signal:"clicked,auto"`
}

func (c *logoController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com/", 0))
}
