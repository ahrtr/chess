package ui

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/ahrtr/chess/fonts"
)

const (
	uiFontSize = 12
)

var (
	//go:embed ui.png
	UI_png []byte
)

var (
	uiImage      *ebiten.Image
	uiFaceSource *text.GoTextFaceSource
)

func init() {
	// initialize uiImage
	img, _, err := image.Decode(bytes.NewReader(UI_png))
	if err != nil {
		panic(err)
	}
	uiImage = ebiten.NewImageFromImage(img)

	// initialize uiFaceSource
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.RegularFont))
	if err != nil {
		log.Fatal(err)
	}
	uiFaceSource = s
}
