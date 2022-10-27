package chess

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

var (
	//go:embed images/board-gray.png
	baordGrayData []byte

	//go:embed images/white-king.png
	whiteKingData []byte
	//go:embed images/white-queen.png
	whiteQueenData []byte
	//go:embed images/white-rook.png
	whiteRookData []byte
	//go:embed images/white-knight.png
	whiteKnightData []byte
	//go:embed images/white-bishop.png
	whiteBishopData []byte
	//go:embed images/white-pawn.png
	whitePawnData []byte

	//go:embed images/black-king.png
	blackKingData []byte
	//go:embed images/black-queen.png
	blackQueenData []byte
	//go:embed images/black-rook.png
	blackRookData []byte
	//go:embed images/black-knight.png
	blackKnightData []byte
	//go:embed images/black-bishop.png
	blackBishopData []byte
	//go:embed images/black-pawn.png
	blackPawnData []byte
)

func createImage(data []byte) image.Image {
	png, _ := png.Decode(bytes.NewReader(data))
	return png
}

type Piece image.Image

var (
	WhiteKing   Piece = createImage(whiteKingData)
	WhiteQueen  Piece = createImage(whiteQueenData)
	WhiteRook   Piece = createImage(whiteRookData)
	WhiteKnight Piece = createImage(whiteKnightData)
	WhiteBishop Piece = createImage(whiteBishopData)
	WhitePawn   Piece = createImage(whitePawnData)

	BlackKing   Piece = createImage(blackKingData)
	BlackQueen  Piece = createImage(blackQueenData)
	BlackRook   Piece = createImage(blackRookData)
	BlackKnight Piece = createImage(blackKnightData)
	BlackBishop Piece = createImage(blackBishopData)
	BlackPawn   Piece = createImage(blackPawnData)
)
