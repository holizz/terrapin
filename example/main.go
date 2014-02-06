package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"image"
	"image/png"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Get("/example1.png", func(w http.ResponseWriter) {
		i := image.NewRGBA(image.Rect(0, 0, 300, 300))
		png.Encode(w, i)
	})

	m.Run()
}
