package main

import (
	"flag"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ahrtr/chess/rules"
	"github.com/ahrtr/chess/ui"
)

const (
	buttonX0    = 128
	buttonWidth = 96
	buttonGap   = 12
	buttonY0    = 12
	buttonY1    = 48
)

var (
	history        []*rules.Board
	historyPointer = -1
)

type Game struct {
	chessBoard *rules.Board
	undoButton *ui.Button
	redoButton *ui.Button
}

func NewGame(selfColor rules.PieceColor) *Game {
	board, err := rules.NewBoard(selfColor)
	if err != nil {
		log.Fatalf("Failed to create the board: %v", err)
	}

	g := &Game{
		chessBoard: board,
		undoButton: ui.NewButton(image.Rect(buttonX0, buttonY0, buttonX0+buttonWidth, buttonY1), "Undo", nil),
		redoButton: ui.NewButton(image.Rect(buttonX0+buttonWidth+buttonGap, buttonY0, buttonX0+buttonWidth*2+buttonGap, buttonY1), "Redo", nil),
	}
	g.backup()
	g.undoButton.SetOnClick(func(_ *ui.Button) {
		g.undo()
	})
	g.redoButton.SetOnClick(func(_ *ui.Button) {
		g.redo()
	})

	return g
}

func (g *Game) backup() {
	clone := g.chessBoard.Clone()
	if len(history) != historyPointer-1 {
		history = history[:(historyPointer + 1)]
	}
	history = append(history, clone)
	historyPointer = len(history) - 1
}

func (g *Game) undo() {
	if historyPointer > 0 {
		historyPointer--
		g.historyOperation()
	}
}

func (g *Game) redo() {
	if historyPointer < len(history)-1 {
		historyPointer++
		g.historyOperation()
	}
}

func (g *Game) historyOperation() {
	clone := history[historyPointer].Clone()
	clone.ResetTimer()
	g.chessBoard = clone
}

func (g *Game) Update() error {
	if g.chessBoard.Update() {
		g.backup()
	}
	g.undoButton.Update()
	g.redoButton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chessBoard.Draw(screen)
	g.undoButton.Draw(screen)
	g.redoButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func selfColor() rules.PieceColor {
	color := flag.String("color", "red", "decide self color(defaults to red).")
	flag.Parse()
	if *color == string(rules.Red) {
		return rules.Red
	}
	if *color == string(rules.Black) {
		return rules.Black
	}
	panic(fmt.Sprintf("invalid color: %s", *color))
}

func main() {
	color := selfColor()
	game := NewGame(color)

	ebiten.SetWindowSize(rules.WindowsWidth, rules.WindowsHeight)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
