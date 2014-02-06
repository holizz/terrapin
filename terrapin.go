package terrapin

import (
	"image"
	"image/color"
	"math"
)

type Terrapin struct {
	Image       *image.RGBA
	Pos         Position
	Orientation float64
	Color       color.Color
}

type Position struct {
	X, Y float64
}

func NewTerrapin(i *image.RGBA, starting Position) (t *Terrapin) {
	t = &Terrapin{
		Image:       i,
		Pos:         starting,
		Orientation: 0.0,
		Color:       color.Black,
	}

	return
}

func (t *Terrapin) Forward(dist float64) {
	for i := 0; i < int(dist); i++ {
		t.Image.Set(int(t.Pos.X), int(t.Pos.Y), t.Color)

		x := 1.0 * math.Sin(t.Orientation)
		y := 1.0 * -math.Cos(t.Orientation)

		t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
	}
}

func (t *Terrapin) Right(radians float64) {
	t.Orientation += radians
}
