package chess

import (
	"image"
	"image/color/palette"
	"image/draw"
	"regexp"
)

var BoardGray = createImage(baordGrayData)

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

func indices(cell string) (int, int) {
	pat := regexp.MustCompile("^([a-h])([1-9])$")
	sm := pat.FindStringSubmatch(cell)
	return columns[sm[1][0]], rows[sm[2][0]]
}

type Board [8][8]Piece

func (b *Board) Put(cell string, piece Piece) {
	c, r := indices(cell)
	b[r][c] = piece
}

func (b *Board) Get(cell string) Piece {
	c, r := indices(cell)
	return b[r][c]
}

func (b *Board) Move(from, to string) {
	b.Put(to, b.Get(from))
	b.Put(from, nil)
}

func (b *Board) Image() *image.Paletted {
	zero := image.Pt(0, 0)
	box := image.Rect(0, 0, 60, 60)

	img := image.NewPaletted(BoardGray.Bounds(), palette.Plan9)
	draw.Draw(img, BoardGray.Bounds(), BoardGray, zero, draw.Src)

	for _, ri := range rows {
		for _, ci := range columns {
			if pic := b[ci][ri]; pic != nil {
				place := box.Add(image.Pt(ri*60, ci*60))
				draw.Draw(img, place, pic, zero, draw.Over)
			}
		}
	}

	return img
}

func (b *Board) Initial() {
	b.Put("a1", WhiteRook)
	b.Put("b1", WhiteKnight)
	b.Put("c1", WhiteBishop)
	b.Put("d1", WhiteQueen)
	b.Put("e1", WhiteKing)
	b.Put("f1", WhiteBishop)
	b.Put("g1", WhiteKnight)
	b.Put("h1", WhiteRook)

	b.Put("a2", WhitePawn)
	b.Put("b2", WhitePawn)
	b.Put("c2", WhitePawn)
	b.Put("d2", WhitePawn)
	b.Put("e2", WhitePawn)
	b.Put("f2", WhitePawn)
	b.Put("g2", WhitePawn)
	b.Put("h2", WhitePawn)

	b.Put("a8", BlackRook)
	b.Put("b8", BlackKnight)
	b.Put("c8", BlackBishop)
	b.Put("d8", BlackQueen)
	b.Put("e8", BlackKing)
	b.Put("f8", BlackBishop)
	b.Put("g8", BlackKnight)
	b.Put("h8", BlackRook)

	b.Put("a7", BlackPawn)
	b.Put("b7", BlackPawn)
	b.Put("c7", BlackPawn)
	b.Put("d7", BlackPawn)
	b.Put("e7", BlackPawn)
	b.Put("f7", BlackPawn)
	b.Put("g7", BlackPawn)
	b.Put("h7", BlackPawn)
}
