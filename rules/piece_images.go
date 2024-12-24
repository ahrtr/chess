package rules

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ahrtr/chess/images"
)

var (
	pieceSolidImageDataMap = map[Piece][]byte{
		Piece{Red, RoleRook}:   images.RedRookPng,
		Piece{Red, RoleHorse}:  images.RedHorsePng,
		Piece{Red, RoleBishop}: images.RedBishopPng,
		Piece{Red, RoleGuard}:  images.RedGuardPng,
		Piece{Red, RoleKing}:   images.RedKingPng,
		Piece{Red, RoleCannon}: images.RedCannonPng,
		Piece{Red, RoleSolder}: images.RedSoldierPng,

		Piece{Black, RoleRook}:   images.BlackRookPng,
		Piece{Black, RoleHorse}:  images.BlackHorsePng,
		Piece{Black, RoleBishop}: images.BlackBishopPng,
		Piece{Black, RoleGuard}:  images.BlackGuardPng,
		Piece{Black, RoleKing}:   images.BlackKingPng,
		Piece{Black, RoleCannon}: images.BlackCannonPng,
		Piece{Black, RoleSolder}: images.BlackSoldierPng,
	}

	pieceDashImageDataMap = map[Piece][]byte{
		Piece{Red, RoleRook}:   images.RedRookDashPng,
		Piece{Red, RoleHorse}:  images.RedHorseDashPng,
		Piece{Red, RoleBishop}: images.RedBishopDashPng,
		Piece{Red, RoleGuard}:  images.RedGuardDashPng,
		Piece{Red, RoleKing}:   images.RedKingDashPng,
		Piece{Red, RoleCannon}: images.RedCannonDashPng,
		Piece{Red, RoleSolder}: images.RedSoldierDashPng,

		Piece{Black, RoleRook}:   images.BlackRookDashPng,
		Piece{Black, RoleHorse}:  images.BlackHorseDashPng,
		Piece{Black, RoleBishop}: images.BlackBishopDashPng,
		Piece{Black, RoleGuard}:  images.BlackGuardDashPng,
		Piece{Black, RoleKing}:   images.BlackKingDashPng,
		Piece{Black, RoleCannon}: images.BlackCannonDashPng,
		Piece{Black, RoleSolder}: images.BlackSoldierDashPng,
	}

	pieceImageMap = map[Piece][2]*ebiten.Image{
		Piece{Red, RoleRook}:   {nil, nil},
		Piece{Red, RoleHorse}:  {nil, nil},
		Piece{Red, RoleBishop}: {nil, nil},
		Piece{Red, RoleGuard}:  {nil, nil},
		Piece{Red, RoleKing}:   {nil, nil},
		Piece{Red, RoleCannon}: {nil, nil},
		Piece{Red, RoleSolder}: {nil, nil},

		Piece{Black, RoleRook}:   {nil, nil},
		Piece{Black, RoleHorse}:  {nil, nil},
		Piece{Black, RoleBishop}: {nil, nil},
		Piece{Black, RoleGuard}:  {nil, nil},
		Piece{Black, RoleKing}:   {nil, nil},
		Piece{Black, RoleCannon}: {nil, nil},
		Piece{Black, RoleSolder}: {nil, nil},
	}

	// All images are supposed to have the same size.
	imageWidth  int
	imageHeight int
)

func initializePieceImageMap() error {
	for p, _ := range pieceImageMap {
		solidImg, _, err := image.Decode(bytes.NewReader(p.imageData(false)))
		if err != nil {
			return fmt.Errorf("error loading solid image for piece (%v): %w", p, err)
		}
		dashImg, _, err := image.Decode(bytes.NewReader(p.imageData(true)))
		if err != nil {
			return fmt.Errorf("error loading dash image for piece (%v): %w", p, err)
		}
		pieceImageMap[p] = [2]*ebiten.Image{ebiten.NewImageFromImage(solidImg), ebiten.NewImageFromImage(dashImg)}
	}

	img := pieceImageMap[Piece{Red, RoleRook}][0]
	imgBound := img.Bounds()
	imageWidth, imageHeight = imgBound.Max.X-imgBound.Min.X, imgBound.Max.Y-imgBound.Min.Y

	return nil
}
