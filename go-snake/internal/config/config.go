package config

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var (
	BlockWidth                   = 20
	FieldWidthInBlocks   float64 = 16
	ScreenWidth                  = BlockWidth*int(FieldWidthInBlocks) + 1
	ScreenHeight                 = BlockWidth*int(FieldWidthInBlocks) + 1
	PauseKey                     = ebiten.KeySpace
	MenuKey                      = ebiten.KeyEscape
	MaxDirQueueSize              = 3
	MoveIntervalInFrames float64 = 10
	BGColor                      = colornames.Aquamarine
)

var Map = map[float64]map[float64]bool{
	1: {
		1: true,
	},
	7: {
		4: true,
	},
}
