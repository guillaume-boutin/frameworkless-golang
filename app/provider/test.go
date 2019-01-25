package provider

import (
	"github.com/sarulabs/di"
)

var TestProvider = di.Def{
	Name: "test",
	Build: func(ctn di.Container) (interface{}, error) {
		return "It's Alive!!!", nil
	},
}
