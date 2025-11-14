package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *SceneManager
	counter      int
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Update() error {
	return g.sceneManager.Update()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
