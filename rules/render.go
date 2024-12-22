package rules

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	leftMargin = 40
	topMargin  = 60

	borderLineWidth = 2.0
	innerLineWidth  = 1.0
)

// Board has 10 rows, and 9 columns.
type Board [10][9]*Piece

func Initialize(selfRole PieceColor) (*Board, error) {
	if err := initializePieceImageMap(); err != nil {
		return nil, err
	}

	return initializeBoard(selfRole), nil
}

func initializeBoard(selfRole PieceColor) *Board {
	// Self is red.
	// Black pieces are on top and red pieces are at the bottom area.
	if selfRole == Red {
		return &Board{
			// Black pieces (rows 0~4, top-->down)
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
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
			{
				nil, {Black, RoleCannon}, nil, nil, nil, nil, nil, {Black, RoleCannon}, nil,
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
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},

			// Red pieces (rows 5~9, top-->down)
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
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
				nil, {Red, RoleCannon}, nil, nil, nil, nil, nil, {Red, RoleCannon}, nil,
			},
			{
				nil, nil, nil, nil, nil, nil, nil, nil, nil,
			},
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
		}
	}

	// Self is black.
	// Red pieces are on top and black pieces are at the bottom area.
	return &Board{
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
	}
}

func DrawBoard(screen *ebiten.Image) {
	bounds := screen.Bounds()

	var (
		// width & height of the windows
		windowsWidth, windowsHeight = bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
		// width & height of the game area
		boardWidth, boardHeight = windowsWidth - leftMargin*2, windowsHeight - topMargin*2

		// left top point
		minPoint = image.Point{X: bounds.Min.X + leftMargin, Y: bounds.Min.Y + topMargin}
		// right bottom point
		maxPoint = image.Point{X: minPoint.X + boardWidth, Y: minPoint.Y + boardHeight}
	)

	// Draw the border
	// top row line
	vector.StrokeLine(screen, float32(minPoint.X), float32(minPoint.Y), float32(maxPoint.X), float32(minPoint.Y), borderLineWidth, color.White, false)
	// bottom row line
	vector.StrokeLine(screen, float32(minPoint.X), float32(maxPoint.Y), float32(maxPoint.X), float32(maxPoint.Y), borderLineWidth, color.White, false)
	// left column line
	vector.StrokeLine(screen, float32(minPoint.X), float32(minPoint.Y), float32(minPoint.X), float32(maxPoint.Y), borderLineWidth, color.White, false)
	// right column line
	vector.StrokeLine(screen, float32(maxPoint.X), float32(minPoint.Y), float32(maxPoint.X), float32(maxPoint.Y), borderLineWidth, color.White, false)

	// Draw the internal 8 rows
	heightStep := boardHeight / 9
	for i := 1; i <= 8; i++ {
		vector.StrokeLine(screen, float32(minPoint.X), float32(minPoint.Y+heightStep*i), float32(maxPoint.X), float32(minPoint.Y+heightStep*i), innerLineWidth, color.White, false)
	}

	// Draw the internal 2 sets of columns, and each has 7 columns.
	widthStep := boardWidth / 8
	// the top set
	for i := 1; i <= 7; i++ {
		vector.StrokeLine(screen, float32(minPoint.X+widthStep*i), float32(minPoint.Y), float32(minPoint.X+widthStep*i), float32(minPoint.Y+heightStep*4), innerLineWidth, color.White, false)
	}
	// the bottom set
	for i := 1; i <= 7; i++ {
		vector.StrokeLine(screen, float32(minPoint.X+widthStep*i), float32(minPoint.Y+heightStep*5), float32(minPoint.X+widthStep*i), float32(maxPoint.Y), innerLineWidth, color.White, false)
	}
}

func DrawPieces(screen *ebiten.Image, board *Board) {
	bounds := screen.Bounds()
	var (
		// width & height of the windows
		windowsWidth, windowsHeight = bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
		// step of rows and columns
		widthStep, heightStep = (windowsWidth - leftMargin*2) / 8, (windowsHeight - topMargin*2) / 9
	)

	for i := 0; i < 10; i++ { // 10 rows
		for j := 0; j < 9; j++ { // 9 columns
			p := board[i][j]
			if p == nil {
				continue
			}

			img := pieceImageMap[*p]
			bound := img.Bounds()
			frameWidth, frameHeight := bound.Max.X-bound.Min.X, bound.Max.Y-bound.Min.Y
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-float64(frameWidth/2), -float64(frameHeight/2))
			op.GeoM.Translate(float64(leftMargin+widthStep*j), float64(topMargin+heightStep*i))

			screen.DrawImage(img, op)
		}
	}
}
