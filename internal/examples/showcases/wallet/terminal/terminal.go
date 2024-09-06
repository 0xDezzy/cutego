package terminal

import (
	"github.com/0xDezzy/cutego/quick"

	_ "github.com/0xDezzy/cutego/internal/examples/showcases/wallet/terminal/controller"
)

func init() { terminalTemplate_QmlRegisterType2("TerminalTemplate", 1, 0, "TerminalTemplate") }

type terminalTemplate struct {
	quick.QQuickItem

	_ func(cmd string) string `slot:"command,->(controller.NewTerminalController(nil))"`
}
