package main

import (
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

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(GameName)

	img, _, err := ebitenutil.NewImageFromFile("resources/tetris.png")
	if err != nil {
		panic(err)
	}

	imgs := []image.Image{img}
	ebiten.SetWindowIcon(imgs)

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
