package chesspiece

type PType int

const (
	BISHOP PType = iota
	KNIGHT
	QUEEN
	KING
	ROOK
	PAWN
)

type Colour int

const (
	WHITE Colour = iota
	BLACK
)

type ChessPiece struct {
	pieceType PType
	colour    Colour
	row       int
	column    int
}

func (p ChessPiece) MoveAllowed(row, col int) bool {

	if p.pieceType == BISHOP {
		if row-col == p.row-p.column {
			return true
		}
		if row+col == p.row+p.column {
			return true
		}
	}

	if p.pieceType == KNIGHT {
		if (p.row == row+2 || p.row == row-2) && (p.column == col+1) || (p.column == col-1) {
			return true
		}

		if (p.row == row+1 || p.row == row-1) && (p.column == col+2) || (p.column == col-2) {
			return true
		}

	}

	if p.pieceType == ROOK {
		if p.row == row && p.column != col {
			return true
		}
		if p.column == col && p.row != row {
			return true
		}
	}

	if p.pieceType == QUEEN {
		if p.row == row && p.column != col {
			return true
		}
		if p.column == col && p.row != row {
			return true
		}
		if row-col == p.row-p.column {
			return true
		}
		if row+col == p.row+p.column {
			return true
		}

	}

	return false
}
