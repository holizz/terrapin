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

		s := 20.0

		// Triangle
		t.Forward(s)
		t.Right(math.Pi * 1 / 2)
		t.Forward(s)
		t.Right(math.Pi * 3 / 4)
		t.Forward(math.Hypot(s, s))

		png.Encode(w, i)
	})

	m.Get("/example2.png", func(w http.ResponseWriter) {
		i := image.NewRGBA(image.Rect(0, 0, 300, 300))

		t := terrapin.NewTerrapin(i, terrapin.Position{150.0, 150.0})

		// Polygons!
		poly(t, 20.0, 3)
		poly(t, 40.0, 5)
		poly(t, 60.0, 7)

		png.Encode(w, i)
	})

	m.Run()
}

func poly(t *terrapin.Terrapin, size float64, sides int) {
	totalInterior := math.Pi * float64(sides-2)
	interior := totalInterior / float64(sides)
	exterior := math.Pi - interior

	// Get into position
	t.PenUp()
	t.Forward(size / 2)
	t.Right(math.Pi/2 + exterior/2)
	t.PenDown()

	// Draw it
	for i := 0; i < sides; i++ {
		t.Forward(size)
		t.Right(exterior)
	}

	// Get back to where we started
	t.PenUp()
	t.Left(math.Pi/2 + exterior/2)
	t.Backward(size / 2)
	t.PenDown()
}
