package chess

import (
	_ "embed"
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
