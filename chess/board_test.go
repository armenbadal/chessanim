package chess

import (
	"image/gif"
	"image/png"
	"os"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := NewBoard()
	if board == nil {
		t.Fail()
	}

	k := createImage(whiteKingData)
	pk := &Piece{colour: White, name: King, image: k}
	board.Put(Cell{Column: 'b', Row: '2'}, pk)

	q := createImage(blackQueenData)
	pq := &Piece{colour: Black, name: Queen, image: q}
	board.Put(Cell{Column: 'd', Row: '4'}, pq)

	img := board.Image()

	f, _ := os.Create("test.png")
	defer f.Close()
	png.Encode(f, img)

	g, _ := os.Create("test.gif")
	defer g.Close()
	gif.Encode(g, img, nil)
}
