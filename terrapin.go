package terrapin

import (
	"image"
)

type Terrapin struct {
	image *image.Image
}

func NewTerrapin(i *image.Image) (t *Terrapin) {
	t = &Terrapin{
		image: i,
	}

	return
}
