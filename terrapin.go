package terrapin

import (
	"image"
	"image/color"
)

type Terrapin struct {
	image *image.RGBA
	pos image.Point
	color color.Color
}

func NewTerrapin(i *image.RGBA, starting image.Point) (t *Terrapin) {
	t = &Terrapin{
		image: i,
		pos: starting,
		color: color.Black,
	}

	return
}

func (t *Terrapin) Forward(dist int) {
	for i := 0; i < dist; i++ {
		t.image.Set(t.pos.X, t.pos.Y, t.color)
		t.pos = t.pos.Add(image.Pt(0, 1))
	}
}
