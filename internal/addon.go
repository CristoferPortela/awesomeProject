package internal

import (
	"awesomeProject/pkg"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"plugin"
)

type RouteFunc func(c *echo.Echo)

type Addon struct {
	Routes []RouteFunc
}

var (
	Routes []RouteFunc
)

type Export struct {
	Name string
	//Routes []RouteFunc
}

func ConfigureAddons() Addon {

	routes := Routes

	for _, addon := range pkg.GetAddons() {
		fmt.Println(addon)
		plug, err := plugin.Open(fmt.Sprintf("addon/%s/main.so", addon))

		if err != nil {
			fmt.Println(err)
			log.Fatal("Error when getting plugins")
		}
		symbol, err := plug.Lookup("Exporting")
		if err != nil {
			fmt.Println(err)
			log.Fatal("Error when looking up the plugins")
		}

		var exported Export
		exported, ok := symbol.(Export)
		if !ok {
			log.Fatal("Error on getting exported value")
		}
		fmt.Println(exported)
	}
	return Addon{
		Routes: routes,
	}

}
