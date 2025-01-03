package rules

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/ahrtr/chess/fonts"
)

const (
	WindowsWidth  = 640
	WindowsHeight = 840

	leftMargin = 40
	topMargin  = 80

	borderLineWidth = 2.0
	innerLineWidth  = 1.0

	msgFontSize     = 16
	msgBottomMargin = 35
)

var (
	boardBackgroundColor = color.RGBA{R: 0xbb, G: 0xad, B: 0xa0, A: 0xff}
)

func (b *Board) Draw(screen *ebiten.Image) {
	drawBoard(screen)
	b.drawPieces(screen)
	b.drawMessage(screen)
}

func drawBoard(screen *ebiten.Image) {
	screen.Fill(boardBackgroundColor)

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

	// Draw diagonal lines in both king areas
	// top
	vector.StrokeLine(screen, float32(minPoint.X+widthStep*3), float32(minPoint.Y), float32(minPoint.X+widthStep*5), float32(minPoint.Y+heightStep*2), borderLineWidth, color.White, false)
	vector.StrokeLine(screen, float32(minPoint.X+widthStep*5), float32(minPoint.Y), float32(minPoint.X+widthStep*3), float32(minPoint.Y+heightStep*2), borderLineWidth, color.White, false)
	// bottom
	vector.StrokeLine(screen, float32(minPoint.X+widthStep*3), float32(maxPoint.Y), float32(minPoint.X+widthStep*5), float32(maxPoint.Y-heightStep*2), borderLineWidth, color.White, false)
	vector.StrokeLine(screen, float32(minPoint.X+widthStep*5), float32(maxPoint.Y), float32(minPoint.X+widthStep*3), float32(maxPoint.Y-heightStep*2), borderLineWidth, color.White, false)
}

func (b *Board) drawPieces(screen *ebiten.Image) {
	bounds := screen.Bounds()
	var (
		// width & height of the windows
		windowsWidth, windowsHeight = bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
		// step of rows and columns
		widthStep, heightStep = (windowsWidth - leftMargin*2) / 8, (windowsHeight - topMargin*2) / 9
	)

	for i := 0; i < 10; i++ { // 10 rows
		for j := 0; j < 9; j++ { // 9 columns
			p := b.pieceMatrix[i][j]
			if p == nil {
				continue
			}

			var img *ebiten.Image
			if (b.selectedFromPoint != nil) && (*b.selectedFromPoint == image.Point{X: i, Y: j}) {
				img = pieceImageMap[*p][1]
			} else {
				img = pieceImageMap[*p][0]
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(-float64(imageWidth/2), -float64(imageHeight/2))
			op.GeoM.Translate(float64(leftMargin+widthStep*j), float64(topMargin+heightStep*i))

			screen.DrawImage(img, op)
		}
	}
}

func (b *Board) drawMessage(screen *ebiten.Image) {
	b.drawTimerAndWinner(screen)
	b.drawHintFromAI(screen)
}

func (b *Board) drawTimerAndWinner(screen *ebiten.Image) {
	timeElapsed := time.Since(b.startTime)
	if b.finalTime.After(b.startTime) {
		timeElapsed = b.finalTime.Sub(b.startTime)
	}
	msg := timeElapsed.Round(time.Second).String()
	if b.finalTime.After(b.startTime) {
		msg += "  winner!"
	}

	op := &text.DrawOptions{}

	bounds := screen.Bounds()
	windowsHeight := bounds.Max.Y - bounds.Min.Y
	if (b.selfColor == Red && b.isRedTurn) || (b.selfColor == Black && !b.isRedTurn) {
		// print the timer at the bottom
		op.GeoM.Translate(10, float64(windowsHeight-msgBottomMargin))

	} else {
		// print the timer at the top
		op.GeoM.Translate(10, 10)
	}

	text.Draw(screen, msg, &text.GoTextFace{
		Source: fonts.TextFaceSource,
		Size:   msgFontSize,
	}, op)
}

func (b *Board) drawHintFromAI(screen *ebiten.Image) {
	if !b.isAIWorking && len(b.hintFromAI) == 0 {
		return
	}

	var msg string
	if b.isAIWorking {
		timeElapsed := time.Since(b.aiStartTime).Round(time.Second)
		msg = fmt.Sprintf("The AI is thinking... %s", timeElapsed.String())
	} else {
		timeElapsed := b.aiStopTime.Sub(b.aiStartTime).Round(time.Second)
		msg = fmt.Sprintf("%s, took: %s", b.hintFromAI, timeElapsed.String())
	}

	bounds := screen.Bounds()
	windowsHeight := bounds.Max.Y - bounds.Min.Y

	op := &text.DrawOptions{}
	op.GeoM.Translate(128, float64(windowsHeight-msgBottomMargin))
	text.Draw(screen, msg, &text.GoTextFace{
		Source: fonts.TextFaceSource,
		Size:   msgFontSize,
	}, op)
}
