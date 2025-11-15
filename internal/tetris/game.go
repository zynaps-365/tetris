package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SceneManager *SceneManager
	counter      int
	ScreenWidth  int
	ScreenHeight int
	GameName     string
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Update() error {
	return g.SceneManager.Update()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}
