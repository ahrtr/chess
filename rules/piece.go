package rules

import "fmt"

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

func (p Piece) String() string {
	return fmt.Sprintf("(%s, %s)", p.color, p.role)
}
