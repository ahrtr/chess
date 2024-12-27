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

var history []*rules.Board

type Game struct {
	chessBoard *rules.Board
	undoButton *ui.Button
}

func NewGame(selfColor rules.PieceColor) *Game {
	board, err := rules.NewBoard(selfColor)
	if err != nil {
		log.Fatalf("Failed to create the board: %v", err)
	}

	g := &Game{
		chessBoard: board,
		undoButton: ui.NewButton(image.Rect(128, 12, 224, 48), "Undo", nil),
	}
	g.backup()
	g.undoButton.SetOnClick(func(_ *ui.Button) {
		g.undo()
	})

	return g
}

func (g *Game) backup() {
	clone := g.chessBoard.Clone()
	history = append(history, clone)
}

func (g *Game) undo() {
	if len(history) > 1 {
		history = history[:(len(history) - 1)]
		clone := history[len(history)-1].Clone()
		clone.ResetTimer()
		g.chessBoard = clone
	}
}

func (g *Game) Update() error {
	if g.chessBoard.Update() {
		g.backup()
	}
	g.undoButton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chessBoard.Draw(screen)
	g.undoButton.Draw(screen)
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
