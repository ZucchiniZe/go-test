package main

import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
)

func main() {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Hello World"
		ctx.Data["Name"] = "World"
		ctx.HTML(200, "hello")
	})

	m.Get("/:name", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Hello " + ctx.Params(":name")
		ctx.Data["Name"] = ctx.Params(":name")
		ctx.HTML(200, "hello")
	})

	m.Run()
}
