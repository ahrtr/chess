package main

import (
	"image/color"
	"log"

	"github.com/ahrtr/chess/rules"
	"github.com/hajimehoshi/ebiten/v2"
)

var backgroundColor = color.RGBA{0xbb, 0xad, 0xa0, 0xff}

var chessBoard *rules.Board

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
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
	ebiten.SetWindowSize(640, 840)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
