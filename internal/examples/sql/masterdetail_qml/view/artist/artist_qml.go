// +build qml

package artist

import (
	"github.com/0xDezzy/cutego/core"
	"github.com/0xDezzy/cutego/quick"

	"github.com/0xDezzy/cutego/internal/examples/sql/masterdetail_qml/controller"
)

func init() {
	artistController_QmlRegisterType2("Artist", 1, 0, "ArtistController")
}

type artistController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//qml->controller
	_ func(row int) `signal:"changeArtist"`
}

func (a *artistController) init() {
	a.SetListModel(controller.Instance.ArtistModel())

	//qml<-controller
	a.ListModel().ConnectRowsInserted(func(*core.QModelIndex, int, int) {
		model := a.ListModel()
		a.SetListModel(nil)
		a.SetListModel(model)
	})

	//qml->controller
	a.ConnectChangeArtist(controller.Instance.ChangeArtist)
}
