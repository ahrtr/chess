package ui

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	uiFontSize = 12
)

var (
	//go:embed ui.png
	UI_png []byte
)

var (
	uiImage *ebiten.Image
)

func init() {
	// initialize uiImage
	img, _, err := image.Decode(bytes.NewReader(UI_png))
	if err != nil {
		panic(err)
	}
	uiImage = ebiten.NewImageFromImage(img)
}
