package main

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/zynaps-365/tetris/internal/tetris"
)

const (
	GameName     = "Tetris"
	ScreenWidth  = 480
	ScreenHeight = 640
)

var (
	//go:embed tetris.png
	tetrisPng []byte
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(GameName)

	icon, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(tetrisPng))
	if err != nil {
		panic(err)
	}

	icons := []image.Image{icon}
	ebiten.SetWindowIcon(icons)

	sm := tetris.NewSceneManager()
	sm.SwitchToMainMenu()
	g := &tetris.Game{
		SceneManager: sm,
		ScreenWidth:  ScreenWidth,
		ScreenHeight: ScreenHeight,
		GameName:     GameName,
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
