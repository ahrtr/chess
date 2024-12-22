package rules

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/ahrtr/chess/images"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	pieceImageMap = map[Piece]*ebiten.Image{
		Piece{Red, RoleRook}:   nil,
		Piece{Red, RoleHorse}:  nil,
		Piece{Red, RoleBishop}: nil,
		Piece{Red, RoleGuard}:  nil,
		Piece{Red, RoleKing}:   nil,
		Piece{Red, RoleCannon}: nil,
		Piece{Red, RoleSolder}: nil,

		Piece{Black, RoleRook}:   nil,
		Piece{Black, RoleHorse}:  nil,
		Piece{Black, RoleBishop}: nil,
		Piece{Black, RoleGuard}:  nil,
		Piece{Black, RoleKing}:   nil,
		Piece{Black, RoleCannon}: nil,
		Piece{Black, RoleSolder}: nil,
	}
)

func (p Piece) imageData() []byte {
	if p.color == Red {
		switch p.role {
		case RoleRook:
			return images.RedRookPng
		case RoleHorse:
			return images.RedHorsePng
		case RoleBishop:
			return images.RedBishopPng
		case RoleGuard:
			return images.RedGuardPng
		case RoleKing:
			return images.RedKingPng
		case RoleCannon:
			return images.RedCannonPng
		case RoleSolder:
			return images.RedSoldierPng
		default:
			panic("unexpected red piece role")
		}
	}

	if p.color == Black {
		switch p.role {
		case RoleRook:
			return images.BlackRookPng
		case RoleHorse:
			return images.BlackHorsePng
		case RoleBishop:
			return images.BlackBishopPng
		case RoleGuard:
			return images.BlackGuardPng
		case RoleKing:
			return images.BlackKingPng
		case RoleCannon:
			return images.BlackCannonPng
		case RoleSolder:
			return images.BlackSoldierPng
		default:
			panic("unexpected black piece role")
		}
	}

	panic("unexpected piece color")
}

func initializePieceImageMap() error {
	for p, _ := range pieceImageMap {
		rawImg, _, err := image.Decode(bytes.NewReader(p.imageData()))
		if err != nil {
			return fmt.Errorf("error loading piece (%v): %w", p, err)
		}
		pieceImageMap[p] = ebiten.NewImageFromImage(rawImg)
	}

	return nil
}
