package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tinne26/etxt"
	"github.com/tinne26/fonts/liberation/lbrtserif"
)

var (
	arialFaceSource *text.GoTextFaceSource

	//go:embed arial.ttf
	arialFont []byte
)

func newFlickeringText(c color.RGBA, size, speed, minValue, maxValue int) *FlickeringText {
	renderer := etxt.NewRenderer()
	renderer.SetFont(lbrtserif.Font())
	renderer.Utils().SetCache8MiB()
	renderer.SetAlign(etxt.HorzCenter)
	renderer.SetSize(float64(size))

	return &FlickeringText{
		renderer: renderer,
		color:    c,
		speed:    speed,
		minValue: minValue,
		maxValue: maxValue,
	}
}

func newTextLabel(c color.RGBA, size int) *TextLabel {
	renderer := etxt.NewRenderer()
	renderer.SetFont(lbrtserif.Font())
	renderer.Utils().SetCache8MiB()
	renderer.SetAlign(etxt.HorzCenter)
	renderer.SetColor(c)
	renderer.SetSize(float64(size))

	return &TextLabel{
		renderer: renderer,
	}
}

type FlickeringText struct {
	renderer *etxt.Renderer
	color    color.RGBA
	counter  int
	speed    int
	minValue int
	maxValue int
}

type TextLabel struct {
	renderer *etxt.Renderer
}

func (tl *TextLabel) Draw(screen *ebiten.Image, txt string, x, y int) {
	tl.renderer.Draw(screen, txt, x, y)
}

func (ft *FlickeringText) Tick() {
	ft.counter = ft.counter + ft.speed

	if ft.counter < ft.minValue {
		ft.counter = ft.minValue
		ft.speed = ft.speed * -1
	}

	if ft.counter > ft.maxValue {
		ft.counter = ft.maxValue
		ft.speed = ft.speed * -1
	}
}

func (ft *FlickeringText) Draw(screen *ebiten.Image, txt string, x, y int) {
	p := float64(ft.counter) / float64(ft.maxValue)
	c := color.RGBA{
		R: uint8(float64(ft.color.R) * p),
		G: uint8(float64(ft.color.G) * p),
		B: uint8(float64(ft.color.B) * p),
		A: uint8(float64(ft.color.A) * p),
	}
	ft.renderer.SetColor(c)
	ft.renderer.Draw(screen, txt, x, y)
}
