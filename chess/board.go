package chess

import (
	"image"
	"image/draw"
)

type Colour byte

const (
	White Colour = 'W'
	Black Colour = 'B'
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
	image  image.Image
}

var (
	WhiteKing  *Piece = &Piece{colour: White, name: King, image: createImage(whiteKingData)}
	WhiteQueen *Piece = &Piece{colour: White, name: Queen, image: createImage(whiteQueenData)}
)

type Board [8][8]struct {
	colour Colour
	piece  *Piece
}

func NewBoard() *Board {
	board := new(Board)
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if (r+c)%2 == 0 {
				board[r][c].colour = Black
			} else {
				board[r][c].colour = White
			}
			board[r][c].piece = nil
		}
	}
	return board
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
	board := createImage(baordGrayData)

	img := image.NewNRGBA(board.Bounds())
	draw.Draw(img, board.Bounds(), board, image.Point{X: 0, Y: 0}, draw.Src)

	for _, ri := range rows {
		for _, ci := range columns {
			pic := b[ci][ri].piece
			if pic != nil {
				place := image.Point{X: ri * 60, Y: ci * 60}
				p2 := image.Rectangle{
					Min: place,
					Max: place.Add(pic.image.Bounds().Size()),
				}
				draw.Draw(img, p2, pic.image, image.Point{X: 0, Y: 0}, draw.Over)
			}
		}
	}

	return img
}
