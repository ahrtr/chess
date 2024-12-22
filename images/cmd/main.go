package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const (
	imageSize    = 60 // size of each image（both width and height share the same value）
	circleRadius = 22
	fontSize     = 20 // font size
	fontX        = 20 // X of the text, need to adjusted
	fontY        = 36 // Y of the text, need to adjusted
)

type pieceText struct {
	name string // used in file name
	text string // used in rendering
}

var pieces = map[string][]pieceText{
	"red": {
		{"rook", "車"},
		{"horse", "馬"},
		{"bishop", "相"},
		{"guard", "仕"},
		{"king", "帥"},
		{"cannon", "炮"},
		{"soldier", "兵"},
	},
	"black": {
		{"rook", "車"},
		{"horse", "馬"},
		{"bishop", "象"},
		{"guard", "士"},
		{"king", "将"},
		{"cannon", "炮"},
		{"soldier", "卒"},
	},
}

func addText(baseImage *image.RGBA, text string, point image.Point, col color.Color, fontSize float64) error {
	fontBytes, err := os.ReadFile("font.ttf")
	if err != nil {
		return err
	}

	ttf, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	drawer := &font.Drawer{
		Dst:  baseImage,
		Src:  image.NewUniform(col),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.I(point.X),
			Y: fixed.I(point.Y),
		},
	}

	drawer.DrawString(text)

	return nil
}

// drawCircle 使用中点画圆法在 img 上绘制圆形
func drawCircle(img *image.RGBA, centerX, centerY, radius int, col color.Color) {
	x, y := 0, radius
	d := 3 - 2*radius

	putPixel(img, centerX, centerY, x, y, col)

	for x <= y {
		x++
		if d > 0 {
			y--
			d += 4*(x-y) + 10
		} else {
			d += 4*x + 6
		}
		putPixel(img, centerX, centerY, x, y, col)
	}
}

// putPixel 设置圆的八个对称点
func putPixel(img *image.RGBA, cx, cy, x, y int, col color.Color) {
	img.Set(cx+x, cy+y, col)
	img.Set(cx-x, cy+y, col)
	img.Set(cx+x, cy-y, col)
	img.Set(cx-x, cy-y, col)
	img.Set(cx+y, cy+x, col)
	img.Set(cx-y, cy+x, col)
	img.Set(cx+y, cy-x, col)
	img.Set(cx-y, cy-x, col)
}

func createChessPieceImage(piece pieceText, pieceColor color.Color) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	// fill transparent background
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Transparent}, image.Point{}, draw.Src)

	drawCircle(img, imageSize/2, imageSize/2, circleRadius, color.Black)

	if err := addText(img, piece.text, image.Point{X: fontX, Y: fontY}, pieceColor, fontSize); err != nil {
		return nil, fmt.Errorf("error adding text: %w", err)
	}

	return img, nil
}

func saveImage(img *image.RGBA, filename string) error {
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	for colorName, piecesList := range pieces {
		for _, piece := range piecesList {
			var pieceColor color.Color
			if colorName == "red" {
				pieceColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // red
			} else {
				pieceColor = color.RGBA{R: 0, G: 0, B: 0, A: 255} // black
			}

			// create an image for each piece
			img, err := createChessPieceImage(piece, pieceColor)
			if err != nil {
				log.Fatalf("Failed to create image for %v, error: %v\n", piece, err)
			}

			// save the image to a PNG fle
			filename := fmt.Sprintf("../%s_%s.png", colorName, piece.name)
			err = saveImage(img, filename)
			if err != nil {
				log.Fatalf("Failed to save image %s, error: %v\n", filename, err)
			}
			fmt.Printf("Saved image: %s\n", filename)
		}
	}
}
