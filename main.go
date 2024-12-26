package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ahrtr/chess/rules"
)

type Game struct {
	chessBoard *rules.Board
}

func NewGame(selfColor rules.PieceColor) *Game {
	board, err := rules.NewBoard(selfColor)
	if err != nil {
		log.Fatalf("Failed to create the board: %v", err)
	}

	return &Game{
		chessBoard: board,
	}
}

func (g *Game) Update() error {
	g.chessBoard.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chessBoard.Draw(screen)
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
