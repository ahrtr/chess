package rules

import (
	"fmt"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ahrtr/chess/utils"
)

type Board struct {
	// What's your color: red or black?
	selfColor PieceColor
	// Is it the Red's turn to move? Always defaults to true in the beginning.
	isRedTurn bool
	// Is the MouseButtonLeft pressed?
	mouseDown bool
	// The point clicked by mouse
	// Note the point.X is the row number (0-9), and point.Y is the column number (0-8).
	targetPt *image.Point
	// The piece on the selected point should be displayed in dash circle.
	// Note the point.X is the row number (0-9), and point.Y is the column number (0-8).
	selectedFromPoint *image.Point
	// What's the start time that the player is allowed to move?
	startTime time.Time
	// the time that generates the winner. Once the field is set,
	// then the game is over.
	finalTime time.Time
	// the hint from the AI
	hintFromAI string
	// whether the AI is working
	isAIWorking bool
	// the time when the AI starts to work
	aiStartTime time.Time
	aiStopTime  time.Time
	// The board has 10 rows, and 9 columns
	pieceMatrix [10][9]*Piece
}

func NewBoard(selfRole PieceColor) (*Board, error) {
	if err := initializePieceImageMap(); err != nil {
		return nil, err
	}

	return newBoard(selfRole), nil
}

func (b *Board) Clone() *Board {
	clone := &Board{
		selfColor:   b.selfColor,
		isRedTurn:   b.isRedTurn,
		mouseDown:   b.mouseDown,
		startTime:   b.startTime,
		pieceMatrix: b.pieceMatrix,
	}
	if b.selectedFromPoint != nil {
		clone.selectedFromPoint = &image.Point{X: b.selectedFromPoint.X, Y: b.selectedFromPoint.Y}
	}
	if b.targetPt != nil {
		clone.targetPt = &image.Point{X: b.targetPt.X, Y: b.targetPt.Y}
	}
	return clone
}

func (b *Board) ResetTimer() {
	b.startTime = time.Now()
}

func newBoard(selfRole PieceColor) *Board {
	// Self is red.
	// Black pieces are on top and red pieces are at the bottom area.
	if selfRole == Red {
		return &Board{
			selfColor:         Red,
			isRedTurn:         true,
			selectedFromPoint: nil,
			startTime:         time.Now(),
			pieceMatrix: [10][9]*Piece{
				// Black pieces (rows 0~4, top-->down)
				{
					&Piece{Black, RoleRook},
					&Piece{Black, RoleHorse},
					&Piece{Black, RoleBishop},
					&Piece{Black, RoleGuard},
					&Piece{Black, RoleKing},
					&Piece{Black, RoleGuard},
					&Piece{Black, RoleBishop},
					&Piece{Black, RoleHorse},
					&Piece{Black, RoleRook},
				},
				{
					nil, nil, nil, nil, nil, nil, nil, nil, nil,
				},
				{
					nil, &Piece{Black, RoleCannon}, nil, nil, nil, nil, nil, &Piece{Black, RoleCannon}, nil,
				},
				{
					&Piece{Black, RoleSolder},
					nil,
					&Piece{Black, RoleSolder},
					nil,
					&Piece{Black, RoleSolder},
					nil,
					&Piece{Black, RoleSolder},
					nil,
					&Piece{Black, RoleSolder},
				},
				{
					nil, nil, nil, nil, nil, nil, nil, nil, nil,
				},

				// Red pieces (rows 5~9, top-->down)
				{
					nil, nil, nil, nil, nil, nil, nil, nil, nil,
				},
				{
					&Piece{Red, RoleSolder},
					nil,
					&Piece{Red, RoleSolder},
					nil,
					&Piece{Red, RoleSolder},
					nil,
					&Piece{Red, RoleSolder},
					nil,
					&Piece{Red, RoleSolder},
				},
				{
					nil, &Piece{Red, RoleCannon}, nil, nil, nil, nil, nil, &Piece{Red, RoleCannon}, nil,
				},
				{
					nil, nil, nil, nil, nil, nil, nil, nil, nil,
				},
				{
					&Piece{Red, RoleRook},
					&Piece{Red, RoleHorse},
					&Piece{Red, RoleBishop},
					&Piece{Red, RoleGuard},
					&Piece{Red, RoleKing},
					&Piece{Red, RoleGuard},
					&Piece{Red, RoleBishop},
					&Piece{Red, RoleHorse},
					&Piece{Red, RoleRook},
				},
			},
		}
	}

	// Self is black.
	// Red pieces are on top and black pieces are at the bottom area.
	return &Board{
		selfColor:         Black,
		isRedTurn:         true,
		selectedFromPoint: nil,
		startTime:         time.Now(),
		pieceMatrix: [10][9]*Piece{
			// Red pieces (rows 0~4, top-->down)
			{
				{Red, RoleRook},
				{Red, RoleHorse},
				{Red, RoleBishop},
				{Red, RoleGuard},
				{Red, RoleKing},
				{Red, RoleGuard},
				{Red, RoleBishop},
				{Red, RoleHorse},
				{Red, RoleRook},
			},
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
			{
				nil, {Red, RoleCannon}, nil, nil, nil, nil, nil, {Red, RoleCannon}, nil,
			},
			{
				{Red, RoleSolder},
				nil,
				{Red, RoleSolder},
				nil,
				{Red, RoleSolder},
				nil,
				{Red, RoleSolder},
				nil,
				{Red, RoleSolder},
			},
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},

			// Black pieces (rows 5~9, top-->down)
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
			{
				{Black, RoleSolder},
				nil,
				{Black, RoleSolder},
				nil,
				{Black, RoleSolder},
				nil,
				{Black, RoleSolder},
				nil,
				{Black, RoleSolder},
			},
			{
				nil, {Black, RoleCannon}, nil, nil, nil, nil, nil, {Black, RoleCannon}, nil,
			},
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
			{
				{Black, RoleRook},
				{Black, RoleHorse},
				{Black, RoleBishop},
				{Black, RoleGuard},
				{Black, RoleKing},
				{Black, RoleGuard},
				{Black, RoleBishop},
				{Black, RoleHorse},
				{Black, RoleRook},
			},
		},
	}
}

// Update returns true if any piece moves, returns false otherwise.
func (b *Board) Update() bool {
	if b.isGameOver() {
		return false
	}

	moved := false

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		pt := image.Pt(ebiten.CursorPosition())

		targetPt := b.findMouseClickedPoint(pt)
		if targetPt != nil {
			b.mouseDown = true
			b.targetPt = targetPt
		} else {
			b.mouseDown = false
			b.targetPt = nil
		}
	} else {
		if b.mouseDown {
			p := b.pieceMatrix[b.targetPt.X][b.targetPt.Y]
			if p == nil {
				if b.selectedFromPoint != nil {
					selectedPiece := b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y]
					// check whether it's a valid move.
					if selectedPiece.validatePieceMove(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, b) {
						b.move(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, true)
						moved = true
					}
					b.selectedFromPoint = nil
				}
			} else {
				if b.selectedFromPoint == nil {
					// You can't select a piece of the wrong color unless you are capturing the opponent's pieces.
					if (p.color == Red) == b.isRedTurn {
						b.selectedFromPoint = b.targetPt
					}
				} else {
					// You can't capture a piece of the same color.
					if (p.color == Red) != b.isRedTurn {
						// check whether it's a valid capture.
						selectedPiece := b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y]
						if selectedPiece.validatePieceMove(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, b) {
							b.move(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, true)
							moved = true
						}
					}
					b.selectedFromPoint = nil
				}
			}
		}
		b.mouseDown = false
	}

	return moved
}

// findMouseClickedPoint locates the point clicked by the mouse.
// Note the input parameter is the position of the cursor when
// mouse being clicked; while the return parameter is the position
// [row number(0-9): column number(0-8)] on the board.
func (b *Board) findMouseClickedPoint(pt image.Point) *image.Point {
	var (
		// step of rows and columns
		widthStep, heightStep = (WindowsWidth - leftMargin*2) / 8, (WindowsHeight - topMargin*2) / 9
	)

	for i := 0; i < 10; i++ { // 10 rows
		for j := 0; j < 9; j++ { // 9 columns
			targetPt := image.Pt(leftMargin+widthStep*j, topMargin+heightStep*i)
			rect := image.Rect(targetPt.X-imageWidth/2, targetPt.Y-imageHeight/2, targetPt.X+imageWidth/2, targetPt.Y+imageHeight/2)
			if !utils.IsPointInsideRect(pt, rect) {
				continue
			}
			return &image.Point{X: i, Y: j}
		}
	}

	return nil
}

// findKing returns the position on the board of the king of the specified color.
func (b *Board) findKing(color PieceColor) image.Point {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			p := b.pieceMatrix[i][j]
			if p != nil && p.color == color && p.role == RoleKing {
				return image.Point{X: i, Y: j}
			}
		}
	}
	panic(fmt.Sprintf("can't find the king: %s", color))
}

func (b *Board) move(fromX, fromY, toX, toY int, checkWinner bool) {
	b.pieceMatrix[toX][toY] = b.pieceMatrix[fromX][fromY]
	b.pieceMatrix[fromX][fromY] = nil
	b.resetAI()

	if checkWinner {
		cloneBoard := b.Clone()
		cloneBoard.switchPlayer()
		if cloneBoard.isWinner() {
			b.finalTime = time.Now()
			return
		}
	}
	b.switchPlayer()
}

func (b *Board) switchPlayer() {
	b.isRedTurn = !b.isRedTurn
	b.startTime = time.Now()
}

func (b *Board) validMoves() []Move {
	color := b.color()

	var allRoutes []Move
	// get all the valid moves of the pieces of the current active player
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 8; j++ {
			p := b.pieceMatrix[i][j]
			if p == nil || p.color != color {
				continue
			}
			routes := p.validMoves(b, image.Point{X: i, Y: j})
			for _, r := range routes {
				allRoutes = append(allRoutes, Move{Piece: *p, route: r})
			}
		}
	}

	return allRoutes
}

// isWinner should be called right after `move`, to check whether
// the `move` has resulted to a winner.
func (b *Board) isWinner() bool {
	allRoutes := b.validMoves()
	// No valid routes (困毙)
	if len(allRoutes) == 0 {
		return true
	}

	color := b.color()
	if !isKingInDanger(b, color) {
		return false
	}

	// try all the possible valid move, and check whether it can
	// resolve the danger of the king. If the king won't be in
	// danger anymore after the possible valid move, then no winner
	// generated yet.
	for _, r := range allRoutes {
		clonedBoard := b.Clone()
		// The last parameter must be `false`, otherwise it will lead to
		// indefinite recursion.
		clonedBoard.move(r.from.X, r.from.Y, r.to.X, r.to.Y, false)
		if !isKingInDanger(clonedBoard, color) {
			return false
		}
	}

	return true
}

func (b *Board) isGameOver() bool {
	return b.finalTime.After(b.startTime)
}

// `color` returns the color of the current active side.
func (b *Board) color() PieceColor {
	if b.isRedTurn {
		return Red
	}
	return Black
}
