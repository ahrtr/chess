package rules

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	selectedFromPoint *image.Point
	// What's the start time that the player is allowed to move?
	startTime time.Time
	// The board has 10 rows, and 9 columns
	pieceMatrix [10][9]*Piece
}

func NewBoard(selfRole PieceColor) (*Board, error) {
	if err := initializePieceImageMap(); err != nil {
		return nil, err
	}

	return newBoard(selfRole), nil
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

func (b *Board) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		pt := image.Pt(ebiten.CursorPosition())

		targetPt := b.findTargetPoint(pt)
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
					if selectedPiece.canMove(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, b) {
						b.pieceMatrix[b.targetPt.X][b.targetPt.Y] = b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y]
						b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y] = nil

						b.isRedTurn = !b.isRedTurn
						b.startTime = time.Now()
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
						if selectedPiece.canMove(b.selectedFromPoint.X, b.selectedFromPoint.Y, b.targetPt.X, b.targetPt.Y, b) {
							b.pieceMatrix[b.targetPt.X][b.targetPt.Y] = b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y]
							b.pieceMatrix[b.selectedFromPoint.X][b.selectedFromPoint.Y] = nil

							b.isRedTurn = !b.isRedTurn
							b.startTime = time.Now()
						}
					}
					b.selectedFromPoint = nil
				}
			}
		}
		b.mouseDown = false
	}
}

func (b *Board) findTargetPoint(pt image.Point) *image.Point {
	var (
		// step of rows and columns
		widthStep, heightStep = (WindowsWidth - leftMargin*2) / 8, (WindowsHeight - topMargin*2) / 9
	)

	for i := 0; i < 10; i++ { // 10 rows
		for j := 0; j < 9; j++ { // 9 columns
			targetPt := image.Pt(leftMargin+widthStep*j, topMargin+heightStep*i)
			rect := image.Rect(targetPt.X-imageWidth/2, targetPt.Y-imageHeight/2, targetPt.X+imageWidth/2, targetPt.Y+imageHeight/2)
			if !isPointInsideRect(pt, rect) {
				continue
			}
			return &image.Point{X: i, Y: j}
		}
	}

	return nil
}
