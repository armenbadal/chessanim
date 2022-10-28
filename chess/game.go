package chess

import (
	"image"
	"image/gif"
	"os"
)

type Step struct {
	from string
	to   string
}

func Sp(f, t string) Step {
	return Step{from: f, to: t}
}

type Game []Step

func (g *Game) Animation(filename string) {
	board := new(Board)
	board.Initial()

	frames := make([]*image.Paletted, len(*g))
	delays := make([]int, len(*g))
	for i, s := range *g {
		board.Move(s.from, s.to)
		frames[i] = board.Image()
		delays[i] = 100
	}

	w, _ := os.Create(filename)
	defer w.Close()
	gif.EncodeAll(w, &gif.GIF{Image: frames, Delay: delays, LoopCount: 1})
}
