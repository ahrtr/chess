package fonts

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	//go:embed font.ttf
	RegularFont []byte

	TextFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(RegularFont))
	if err != nil {
		panic(fmt.Sprintf("error loading font: %v", err))
	}
	TextFaceSource = s
}
