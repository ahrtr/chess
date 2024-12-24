package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ahrtr/chess/rules"
)

type Game struct {
	chessBoard *rules.Board
}

func NewGame() *Game {
	board, err := rules.NewBoard(rules.Red)
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

func main() {
	game := NewGame()
	ebiten.SetWindowSize(rules.WindowsWidth, rules.WindowsHeight)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
