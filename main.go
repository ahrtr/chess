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

type Game struct {
	chessBoard *rules.Board
	undoButton *ui.Button
	redoButton *ui.Button

	hintButton   *ui.Button
	isAIThinking bool

	history        []*rules.Board
	historyPointer int
}

func NewGame(selfColor rules.PieceColor) *Game {
	board, err := rules.NewBoard(selfColor)
	if err != nil {
		log.Fatalf("Failed to create the board: %v", err)
	}

	g := &Game{
		chessBoard: board,
		undoButton: ui.NewButton(image.Rect(buttonX0, buttonY0, buttonX0+buttonWidth, buttonY1), "Undo", nil),
		redoButton: ui.NewButton(image.Rect(buttonX0+(buttonWidth+buttonGap)*1, buttonY0, buttonX0+(buttonWidth+buttonGap)*1+buttonWidth, buttonY1), "Redo", nil),

		hintButton: ui.NewButton(image.Rect(buttonX0+(buttonWidth+buttonGap)*2, buttonY0, buttonX0+(buttonWidth+buttonGap)*2+buttonWidth, buttonY1), "Hint", nil),

		history:        nil,
		historyPointer: -1,
	}
	g.backup()
	g.undoButton.SetOnClick(func(_ *ui.Button) {
		g.undo()
	})
	g.redoButton.SetOnClick(func(_ *ui.Button) {
		g.redo()
	})

	g.hintButton.SetOnClick(func(_ *ui.Button) {
		g.aiRun()
	})

	return g
}

func (g *Game) backup() {
	clone := g.chessBoard.Clone()
	if len(g.history) != g.historyPointer-1 {
		g.history = g.history[:(g.historyPointer + 1)]
	}
	g.history = append(g.history, clone)
	g.historyPointer = len(g.history) - 1
}

func (g *Game) undo() {
	if g.historyPointer > 0 {
		g.historyPointer--
		g.historyOperation()
	}
}

func (g *Game) redo() {
	if g.historyPointer < len(g.history)-1 {
		g.historyPointer++
		g.historyOperation()
	}
}

func (g *Game) historyOperation() {
	clone := g.history[g.historyPointer].Clone()
	clone.ResetTimer()
	g.chessBoard = clone
}

func (g *Game) aiRun() {
	if g.isAIThinking {
		return
	}

	g.isAIThinking = true
	g.chessBoard.StartAI()

	// let's do it async
	go func() {
		defer func() {
			g.isAIThinking = false
		}()

		cloneBoard := g.chessBoard.Clone()
		// Performance history:
		//   1. 2024-12-31 depth = 4, took 1m30s
		//      Very basic minimax algorithm with alpha-beta pruning improvement.
		bestMove := cloneBoard.GetBestMove(4)

		g.chessBoard.StopAI(fmt.Sprintf("Best move: %s", bestMove.String()))
	}()
}

func (g *Game) Update() error {
	// do nothing when the AI is thinking
	if g.isAIThinking {
		return nil
	}

	if g.chessBoard.Update() {
		g.backup()
	}
	g.undoButton.Update()
	g.redoButton.Update()
	g.hintButton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.chessBoard.Draw(screen)
	g.undoButton.Draw(screen)
	g.redoButton.Draw(screen)
	g.hintButton.Draw(screen)
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
