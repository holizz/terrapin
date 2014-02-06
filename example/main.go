package main

import (
	"github.com/codegangsta/martini"
	"github.com/holizz/terrapin"
	"github.com/martini-contrib/render"
	"image"
	"image/png"
	"math"
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

		t := terrapin.NewTerrapin(i, terrapin.Position{150.0, 150.0})

		t.Forward(20)
		t.Right(math.Pi * 1/2)
		t.Forward(20)
		t.Right(math.Pi * 3/4)
		t.Forward(math.Hypot(20, 20))

		png.Encode(w, i)
	})

	m.Run()
}
