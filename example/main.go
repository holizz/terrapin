package main

import (
	"fmt"
	"github.com/holizz/terrapin"
	"image"
	"image/png"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<!doctype html>
			<title>Terrapin example</title>
			<h1>Terrapin example</h1>

			<img src="/example1.png">
			<img src="/example2.png">
		`))
	})

	http.HandleFunc("/example1.png", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/example2.png", func(w http.ResponseWriter, r *http.Request) {
		i := image.NewRGBA(image.Rect(0, 0, 300, 300))

		t := terrapin.NewTerrapin(i, terrapin.Position{150.0, 150.0})

		// Polygons!
		poly(t, 20.0, 3)
		poly(t, 40.0, 5)
		poly(t, 60.0, 7)

		png.Encode(w, i)
	})

	fmt.Println("Listening on http://localhost:3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
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
