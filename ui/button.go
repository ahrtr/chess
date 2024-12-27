package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/ahrtr/chess/utils"
)

// The button implementation is basically coming from the following link with minor change,
// https://github.com/hajimehoshi/ebiten/blob/db07dcfe9ffec469e3129e1c94504a129893cf84/examples/ui/main.go

type imageType int

const (
	imageTypeButton imageType = iota
	imageTypeButtonPressed
)

var imageSrcRects = map[imageType]image.Rectangle{
	imageTypeButton:        image.Rect(0, 0, 16, 16),
	imageTypeButtonPressed: image.Rect(16, 0, 32, 16),
}

type Button struct {
	rect image.Rectangle
	text string

	mouseDown bool

	onClick func(b *Button)
}

func NewButton(rect image.Rectangle, text string, onClick func(b *Button)) *Button {
	return &Button{
		rect:      rect,
		text:      text,
		mouseDown: false,
		onClick:   onClick,
	}
}

func (b *Button) SetOnClick(f func(b *Button)) {
	b.onClick = f
}

func (b *Button) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if utils.IsPointInsideRect(image.Pt(x, y), b.rect) {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onClick != nil {
				b.onClick(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	drawNinePatches(dst, b.rect, imageSrcRects[t])

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(b.rect.Min.X+b.rect.Max.X)/2, float64(b.rect.Min.Y+b.rect.Max.Y)/2)
	op.ColorScale.ScaleWithColor(color.Black)
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	text.Draw(dst, b.text, &text.GoTextFace{
		Source: uiFaceSource,
		Size:   uiFontSize,
	}, op)
}

func drawNinePatches(dst *ebiten.Image, dstRect image.Rectangle, srcRect image.Rectangle) {
	srcX := srcRect.Min.X
	srcY := srcRect.Min.Y
	srcW := srcRect.Dx()
	srcH := srcRect.Dy()

	dstX := dstRect.Min.X
	dstY := dstRect.Min.Y
	dstW := dstRect.Dx()
	dstH := dstRect.Dy()

	op := &ebiten.DrawImageOptions{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			op.GeoM.Reset()

			sx := srcX
			sy := srcY
			sw := srcW / 4
			sh := srcH / 4
			dx := 0
			dy := 0
			dw := sw
			dh := sh
			switch i {
			case 1:
				sx = srcX + srcW/4
				sw = srcW / 2
				dx = srcW / 4
				dw = dstW - 2*srcW/4
			case 2:
				sx = srcX + 3*srcW/4
				dx = dstW - srcW/4
			}
			switch j {
			case 1:
				sy = srcY + srcH/4
				sh = srcH / 2
				dy = srcH / 4
				dh = dstH - 2*srcH/4
			case 2:
				sy = srcY + 3*srcH/4
				dy = dstH - srcH/4
			}

			op.GeoM.Scale(float64(dw)/float64(sw), float64(dh)/float64(sh))
			op.GeoM.Translate(float64(dx), float64(dy))
			op.GeoM.Translate(float64(dstX), float64(dstY))
			dst.DrawImage(uiImage.SubImage(image.Rect(sx, sy, sx+sw, sy+sh)).(*ebiten.Image), op)
		}
	}
}
