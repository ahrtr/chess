package main

import (
	"log"

	"github.com/ahrtr/chess/rules"
	"github.com/hajimehoshi/ebiten/v2"
)

var chessBoard *rules.Board

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rules.DrawBoard(screen)
	rules.DrawPieces(screen, chessBoard)

	/*	drawBoard(screen)
		for _, p := range pieces {
			p.Draw(screen)
		}*/
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	var err error
	chessBoard, err = rules.Initialize(rules.Red)
	if err != nil {
		log.Fatalf("Failed to initialize the board: %v", err)
	}

	game := &Game{}
	ebiten.SetWindowSize(rules.WindowsWidth, rules.WindowsHeight)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
