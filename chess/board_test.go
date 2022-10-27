package chess

import (
	"image/gif"
	"image/png"
	"os"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := new(Board)
	board.Initial()
	img := board.Image()

	f, _ := os.Create("test.png")
	defer f.Close()
	png.Encode(f, img)

	g, _ := os.Create("test.gif")
	defer g.Close()
	gif.Encode(g, img, nil)
}
