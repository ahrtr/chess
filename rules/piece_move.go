package rules

import "fmt"

// validatePieceMove validates whether it's a valid move.
//  1. It should follow the rules of chinese chess;
//  2. The king shouldn't be in danger after the move.
func (p Piece) validatePieceMove(fromX, fromY, toX, toY int, b *Board) bool {
	if !p.canMove(fromX, fromY, toX, toY, b) {
		return false
	}
	clonedBoard := b.Clone()
	clonedBoard.move(fromX, fromY, toX, toY)
	return !isKingInDanger(clonedBoard, p.color)
}

// canMove checks whether it's a valid move from [fromX, fromY] to [toX, toY].
// It doesn't need to care about the color (soldier is an exception, because
// it can only move forward, not backward). It's already guaranteed that the
// Point{toX, toY} is
//   - either empty
//   - or there is an opponent's piece
func (p Piece) canMove(fromX, fromY, toX, toY int, b *Board) bool {
	switch p.role {
	case RoleRook:
		return canRookMove(fromX, fromY, toX, toY, b)
	case RoleHorse:
		return canHorseMove(fromX, fromY, toX, toY, b)
	case RoleBishop:
		return canBishopMove(fromX, fromY, toX, toY, b)
	case RoleGuard:
		return canGuardMove(fromX, fromY, toX, toY, b)
	case RoleKing:
		return canKingMove(fromX, fromY, toX, toY, b)
	case RoleCannon:
		return canCannonMove(fromX, fromY, toX, toY, b)
	case RoleSolder:
		return canSoldierMove(fromX, fromY, toX, toY, b)
	default:
		panic(fmt.Sprintf("invalid %v move from [%d: %d] to [%d: %d]", p, fromX, fromY, toX, toY))
	}
}

// Rules:
//  1. Must be in a horizontal or vertical line;
//  2. There is no any obstacle along the route.
func canRookMove(fromX, fromY, toX, toY int, b *Board) bool {
	if fromX != toX && fromY != toY {
		return false
	}
	if fromX == toX {
		y1, y2 := min(fromY, toY), max(fromY, toY)
		for y := y1 + 1; y < y2; y++ {
			if b.pieceMatrix[fromX][y] != nil {
				return false
			}
		}
	}
	if fromY == toY {
		x1, x2 := min(fromX, toX), max(fromX, toX)
		for x := x1 + 1; x < x2; x++ {
			if b.pieceMatrix[x][fromY] != nil {
				return false
			}
		}
	}
	return true
}

// Rules:
//  1. The route must be an 'L' shape (or 日);
//  2. The horse's leg is not blocked (no neighbour along the long side).
func canHorseMove(fromX, fromY, toX, toY int, b *Board) bool {
	if !(abs(toX-fromX) == 1 && abs(toY-fromY) == 2) &&
		!(abs(toX-fromX) == 2 && abs(toY-fromY) == 1) {
		return false
	}
	if abs(toX-fromX) == 2 {
		if b.pieceMatrix[(fromX+toX)/2][fromY] != nil {
			return false
		}
	}
	if abs(toY-fromY) == 2 {
		if b.pieceMatrix[fromX][(fromY+toY)/2] != nil {
			return false
		}
	}
	return true
}

// Rules:
//  1. The route must be in a "field" (or 田) pattern (diagonal, two steps each time);
//  2. The "eye" must not be blocked (the center point of the 田);
//  3. It can't across the country boarder.
func canBishopMove(fromX, fromY, toX, toY int, b *Board) bool {
	if abs(toX-fromX) != 2 || abs(toY-fromY) != 2 {
		return false
	}
	if b.pieceMatrix[(fromX+toX)/2][(fromY+toY)/2] != nil {
		return false
	}
	if (fromX <= 4 && toX > 4) || (fromX > 4 && toX <= 4) {
		return false
	}
	return true
}

// Rules:
//  1. The guard can only move within its own 3*3 grid (九宫格) near the King;
//  2. It moves diagonally by one step at a time.
func canGuardMove(fromX, fromY, toX, toY int, b *Board) bool {
	if abs(toX-fromX) != 1 || abs(toY-fromY) != 1 {
		return false
	}
	if fromY < 3 || fromY > 5 || toY < 3 || toY > 5 {
		return false
	}
	if !(fromX >= 0 && fromX <= 2 && toX >= 0 && toX <= 2) && !(fromX >= 7 && fromX <= 9 && toX >= 7 && toX <= 9) {
		return false
	}
	return true
}

// Rules:
//  1. The king can only move within its own 3*3 grid (九宫格);
//  2. It moves horizontally or vertically by one step at a time.
func canKingMove(fromX, fromY, toX, toY int, b *Board) bool {
	if !(fromX == toX && abs(toY-fromY) == 1) && !(fromY == toY && abs(toX-fromX) == 1) {
		return false
	}
	if fromY < 3 || fromY > 5 || toY < 3 || toY > 5 {
		return false
	}
	if !(fromX >= 0 && fromX <= 2 && toX >= 0 && toX <= 2) && !(fromX >= 7 && fromX <= 9 && toX >= 7 && toX <= 9) {
		return false
	}
	return true
}

// Rules:
//  1. Must be in a horizontal or vertical line;
//  2. Move or capture:
//     2.1 There is no any obstacle along the route for move case.
//     2.2 There must be one and only one platform for capture case.
func canCannonMove(fromX, fromY, toX, toY int, b *Board) bool {
	if fromX != toX && fromY != toY {
		return false
	}
	// Move case: no any obstacle along the route.
	if b.pieceMatrix[toX][toY] == nil {
		if fromX == toX {
			y1, y2 := min(fromY, toY), max(fromY, toY)
			for y := y1 + 1; y < y2; y++ {
				if b.pieceMatrix[fromX][y] != nil {
					return false
				}
			}
		}
		if fromY == toY {
			x1, x2 := min(fromX, toX), max(fromX, toX)
			for x := x1 + 1; x < x2; x++ {
				if b.pieceMatrix[x][fromY] != nil {
					return false
				}
			}
		}
	} else {
		// capture case: there must be one and only one platform.
		if fromX == toX {
			cnt := 0
			y1, y2 := min(fromY, toY), max(fromY, toY)
			for y := y1 + 1; y < y2; y++ {
				if b.pieceMatrix[fromX][y] != nil {
					cnt++
				}
			}
			if cnt != 1 {
				return false
			}
		}
		if fromY == toY {
			cnt := 0
			x1, x2 := min(fromX, toX), max(fromX, toX)
			for x := x1 + 1; x < x2; x++ {
				if b.pieceMatrix[x][fromY] != nil {
					cnt++
				}
			}
			if cnt != 1 {
				return false
			}
		}
	}

	return true
}

// Rules:
//  1. Must be in a horizontal or vertical line by one step at a time; but
//     horizontal move isn't allowed inside its own country.
//  2. It can only move forward, not backward.
func canSoldierMove(fromX, fromY, toX, toY int, b *Board) bool {
	if !(fromX == toX && abs(toY-fromY) == 1) && !(fromY == toY && abs(toX-fromX) == 1) {
		return false
	}

	fromBottom := (b.selfColor == Red) == b.isRedTurn

	// backward not allowed
	if fromY == toY {
		if fromBottom {
			if fromX-toX != 1 {
				return false
			}
		} else {
			if toX-fromX != 1 {
				return false
			}
		}
	}

	// horizontal move isn't allowed inside its own country
	if fromX == toX {
		if fromBottom {
			if fromX >= 5 {
				return false
			}
		} else {
			if fromX <= 4 {
				return false
			}
		}

	}

	return true
}

// isKingInDanger checks whether the king of the specified color
// is in danger.
func isKingInDanger(b *Board, color PieceColor) bool {
	pt := b.findKing(color)

	fmt.Printf("location of king X: %d, Y: %d, color: %v\n", pt.X, pt.Y, color)

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			p := b.pieceMatrix[i][j]
			if p == nil || p.color == color {
				continue
			}
			// If anyone of the opponent's pieces can capture the king,
			// then it's in danger.
			if p.canMove(i, j, pt.X, pt.Y, b) {
				return true
			}
		}
	}
	return false
}
