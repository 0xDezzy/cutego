package c

import (
	"github.com/0xDezzy/cutego/core"

	_ "github.com/0xDezzy/cutego/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
