# Terrapin

[Turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics) for Golang using image.RGBA

## Usage

    i := image.NewRGBA(image.Rect(0, 0, 300, 300))
    t := terrapin.NewTerrapin(i, terrapin.Position{150.0, 150.0})
    
    t.Forward(100)
    t.Right(math.Pi * 1 / 2)
    t.Forward(100)
    t.Right(math.Pi * 3 / 4)
    t.Forward(math.Hypot(100, 100))
    
    png.Encode(writer, i)

## Examples

To see a couple of little examples:

    go get github.com/holizz/terrapin/...
    go run $GOPATH/src/github.com/holizz/terrapin/example/main.go

## License

MIT (see LICENSE.txt)
