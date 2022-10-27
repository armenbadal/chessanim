package chess

import (
	"image"
)

type Colour byte

const (
	White byte = 'W'
	Black byte = 'B'
)

type Name byte

const (
	King Name = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

type Cell struct {
	Row    byte
	Column byte
}

var rows = map[byte]int{
	'8': 0,
	'7': 1,
	'6': 2,
	'5': 3,
	'4': 4,
	'3': 5,
	'2': 6,
	'1': 7,
}

var columns = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}

func (c *Cell) indices() (int, int) {
	return rows[c.Row], columns[c.Column]
}

type Piece struct {
	colour Colour
	name   Name
}

type Board [8][8]struct {
	colour Colour
	piece  *Piece
}

func (b *Board) Put(cell Cell, piece *Piece) {
	r, c := cell.indices()
	b[r][c].piece = piece
}

func (b *Board) Get(cell Cell) *Piece {
	r, c := cell.indices()
	return b[r][c].piece
}

func (b *Board) Move(from, to Cell) {
	b.Put(to, b.Get(from))
	b.Put(from, nil)
}

func (b *Board) Image() image.Image {
	return nil
}
