package tetris

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenu struct {
	ft *FlickeringText
	tl *TextLabel
	sc *SceneManager
}

func newMainMenu(sc *SceneManager) *MainMenu {
	c := color.RGBA{91, 239, 91, 255}

	ft := newFlickeringText(c, 32, 5, 64, 255)
	tl := newTextLabel(c, 64)

	m := &MainMenu{
		ft: ft,
		tl: tl,
		sc: sc,
	}

	return m
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	bounds := screen.Bounds()
	dx := bounds.Dx() / 2
	dy := bounds.Dy()

	m.tl.Draw(screen, "TETRIS", dx, dy/2)
	m.ft.Draw(screen, "Press SPACE to start", dx, dy-dy/10)
}
func (m *MainMenu) Update() error {
	m.ft.Tick()
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		m.sc.SwitchGameBoard()
	}

	return nil
}
