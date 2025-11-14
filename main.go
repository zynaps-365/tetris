package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	gameName     = "Tetris"
	ScreenWidth  = 480
	ScreenHeight = 640
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(gameName)

	img, _, err := ebitenutil.NewImageFromFile("tetris.png")
	if err != nil {
		panic(err)
	}

	imgs := []image.Image{img}
	ebiten.SetWindowIcon(imgs)

	sm := NewSceneManager()
	sm.switchToMainMenu()
	g := &Game{sceneManager: sm}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
