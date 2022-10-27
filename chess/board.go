package chess

import (
	"image"
	"image/draw"
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

type Board [8][8]Piece

func (b *Board) Put(cell Cell, piece Piece) {
	r, c := cell.indices()
	b[r][c] = piece
}

func (b *Board) Get(cell Cell) Piece {
	r, c := cell.indices()
	return b[r][c]
}

func (b *Board) Move(from, to Cell) {
	b.Put(to, b.Get(from))
	b.Put(from, nil)
}

func (b *Board) Image() image.Image {
	board := createImage(baordGrayData)

	img := image.NewNRGBA(board.Bounds())
	draw.Draw(img, board.Bounds(), board, image.Point{X: 0, Y: 0}, draw.Src)

	for _, ri := range rows {
		for _, ci := range columns {
			pic := b[ci][ri]
			if pic != nil {
				place := image.Point{X: ri * 60, Y: ci * 60}
				p2 := image.Rectangle{
					Min: place,
					Max: place.Add(pic.Bounds().Size()),
				}
				draw.Draw(img, p2, pic, image.Point{X: 0, Y: 0}, draw.Over)
			}
		}
	}

	return img
}

func (b *Board) Initial() {
	b.Put(Cell{Column: 'a', Row: '1'}, WhiteRook)
	b.Put(Cell{Column: 'b', Row: '1'}, WhiteKnight)
	b.Put(Cell{Column: 'c', Row: '1'}, WhiteBishop)
	b.Put(Cell{Column: 'd', Row: '1'}, WhiteQueen)
	b.Put(Cell{Column: 'e', Row: '1'}, WhiteKing)
	b.Put(Cell{Column: 'f', Row: '1'}, WhiteBishop)
	b.Put(Cell{Column: 'g', Row: '1'}, WhiteKnight)
	b.Put(Cell{Column: 'h', Row: '1'}, WhiteRook)

	b.Put(Cell{Column: 'a', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'b', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'c', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'd', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'e', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'f', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'g', Row: '2'}, WhitePawn)
	b.Put(Cell{Column: 'h', Row: '2'}, WhitePawn)

	b.Put(Cell{Column: 'a', Row: '8'}, BlackRook)
	b.Put(Cell{Column: 'b', Row: '8'}, BlackKnight)
	b.Put(Cell{Column: 'c', Row: '8'}, BlackBishop)
	b.Put(Cell{Column: 'd', Row: '8'}, BlackQueen)
	b.Put(Cell{Column: 'e', Row: '8'}, BlackKing)
	b.Put(Cell{Column: 'f', Row: '8'}, BlackBishop)
	b.Put(Cell{Column: 'g', Row: '8'}, BlackKnight)
	b.Put(Cell{Column: 'h', Row: '8'}, BlackRook)

	b.Put(Cell{Column: 'a', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'b', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'c', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'd', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'e', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'f', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'g', Row: '7'}, BlackPawn)
	b.Put(Cell{Column: 'h', Row: '7'}, BlackPawn)
}
