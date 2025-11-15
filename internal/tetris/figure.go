package tetris

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func NewRandomFigure() *Figure {
	var f *Figure
	r := rand.Intn(6)

	switch r {
	case 0:
		f = newOShapeFigure()
	case 1:
		f = newTShapeFigure()
	case 2:
		f = newSShapeFigure()
	case 3:
		f = newZShapeFigure()
	case 4:
		f = newIShapeFigure()
	case 5:
		f = newJShapeFigure()
	case 6:
		f = newLShapeFigure()
	}
	return f
}

type Figure struct {
	content [][]uint32
	x       int
	y       int
}

func (f *Figure) Content() [][]uint32 {
	return f.content
}

func (f *Figure) RotateLeft() {
	return
}

func (f *Figure) RotateRight() {
	return
}

func (f *Figure) Draw(screen *ebiten.Image, c color.RGBA, deltaX, deltaY int) {
	dx := float32(f.x)*pieceWidth + float32(deltaX)
	dy := float32(f.y)*pieceHeight + float32(deltaY)

	for x := 0; x < len(f.content); x++ {
		for y := 0; y < len(f.content[x]); y++ {
			if f.content[x][y] == 1 {
				x1 := float32(x)*pieceWidth + dx
				y1 := float32(y)*pieceWidth + dy

				vector.FillRect(screen, x1, y1, pieceWidth, pieceHeight, c, false)
			}
		}
	}
}

func (f *Figure) GetX() int {
	return f.x
}

func (f *Figure) SetX(x int) {
	f.x = x
}

func (f *Figure) GetY() int {
	return f.y
}

func (f *Figure) Rotate() {
	if len(f.content) != len(f.content[0]) {
		panic("dimensions not equal")
	}
	n := len(f.content)
	rotated := make([][]uint32, n)
	for i := range rotated {
		rotated[i] = make([]uint32, n)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			rotated[c][n-1-r] = f.content[r][c]
		}
	}

	f.content = rotated
}

func (f *Figure) Move(mo MoveOffset) {
	f.y = f.y + mo.y
	f.x = f.x + mo.x
}

func (f *Figure) Width() int {
	return len(f.content)
}

func (f *Figure) Height() int {
	if len(f.content) > 0 {
		return len(f.content[0])
	}
	return 0
}

func newOShapeFigure() *Figure {
	content := [][]uint32{
		{1, 1},
		{1, 1},
	}
	return &Figure{content: content}
}

func newTShapeFigure() *Figure {
	content := [][]uint32{
		{0, 1, 0},
		{1, 1, 0},
		{1, 0, 0},
	}
	return &Figure{content: content}
}

func newZShapeFigure() *Figure {
	content := [][]uint32{
		{0, 1, 1},
		{1, 1, 0},
		{0, 0, 0},
	}
	return &Figure{content: content}
}

func newSShapeFigure() *Figure {
	content := [][]uint32{
		{1, 1, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	return &Figure{content: content}
}

func newIShapeFigure() *Figure {
	content := [][]uint32{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	return &Figure{content: content}
}

func newJShapeFigure() *Figure {
	content := [][]uint32{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	return &Figure{content: content}
}

func newLShapeFigure() *Figure {
	content := [][]uint32{
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	return &Figure{content: content}
}
