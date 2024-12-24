package rules

type PieceRole string

const (
	RoleKing   = PieceRole("king")
	RoleGuard  = PieceRole("guard")
	RoleBishop = PieceRole("bishop")
	RoleHorse  = PieceRole("horse")
	RoleRook   = PieceRole("rook")
	RoleCannon = PieceRole("cannon")
	RoleSolder = PieceRole("soldier")
)

type PieceColor string

const (
	Red   = PieceColor("red")
	Black = PieceColor("black")
)

type Piece struct {
	color PieceColor
	role  PieceRole
}

func (p Piece) imageData(isDash bool) []byte {
	if isDash {
		return pieceDashImageDataMap[p]
	}
	return pieceSolidImageDataMap[p]
}

func (p Piece) canMove(fromX, fromY, toX, toY int, b *Board) bool {

	switch p.role {
	case RoleKing:
	}

	return false
}
