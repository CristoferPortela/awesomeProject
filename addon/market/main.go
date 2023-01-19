package main

//import "github.com/labstack/echo/v4"

type RouteFunc func()

type Export struct {
	Name string
	//Routes []RouteFunc
}

var Exporting = Export{
	Name: "Market",
}
