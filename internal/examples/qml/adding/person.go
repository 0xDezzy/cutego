package main

import (
	"github.com/0xDezzy/cutego/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
