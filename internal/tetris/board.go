package tetris

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	BoardWidth   = 10
	BoardHeight  = 20
	pieceWidth   = float32(30)
	pieceHeight  = float32(30)
	ticksPerMove = 30
)

var (
	glassBorderColor = color.RGBA{91, 239, 91, 255}
	glassBgColor     = color.RGBA{0, 0, 0, 255}
	mainMenuBgColor  = color.RGBA{128, 128, 128, 255}
)

type Board struct {
	field   [BoardWidth][BoardHeight]uint32
	sm      *SceneManager
	figure  *Figure
	next    *Figure
	ticks   int64
	level   int
	lastKey ebiten.Key
	score   int

	tlScoreTxt *TextLabel
	tlScoreVal *TextLabel
}

func newBoard(sm *SceneManager) *Board {
	b := &Board{
		sm:         sm,
		field:      [BoardWidth][BoardHeight]uint32{},
		level:      1,
		tlScoreTxt: newTextLabel(glassBgColor, 32),
		tlScoreVal: newTextLabel(glassBgColor, 32),
	}
	return b
}

func (b *Board) DrawScore(screen *ebiten.Image, dx, dy int) {
	x := int(float32(dx) + BoardWidth*pieceWidth + 90)
	y := dy + 20
	b.tlScoreTxt.Draw(screen, "SCORE", x, y)
	b.tlScoreVal.Draw(screen, strconv.Itoa(b.score), x, y+32)
}

func (b *Board) DrawCup(screen *ebiten.Image, dx, dy int) {
	for x := 0; x < BoardWidth; x++ {
		for y := 0; y < BoardHeight; y++ {
			if b.field[x][y] == 1 {
				x1 := float32(x)*pieceWidth + float32(dx)
				y1 := float32(y)*pieceHeight + float32(dy)
				vector.FillRect(screen, x1, y1, pieceWidth, pieceHeight, glassBorderColor, false)
			}
		}
	}
}

func (b *Board) DrawFigure(screen *ebiten.Image, dx, dy int) {
	if b.figure == nil {
		return
	}

	content := b.figure.Content()

	for x := 0; x < len(content); x++ {
		for y := 0; y < len(content[x]); y++ {
			if content[x][y] == 1 {
				x1 := float32(x)*pieceWidth + float32(dx) + float32(b.figure.GetX())*pieceWidth
				y1 := float32(y)*pieceWidth + float32(dy) + float32(b.figure.GetY())*pieceWidth

				vector.FillRect(screen, x1, y1, pieceWidth, pieceHeight, glassBorderColor, false)
			}
		}
	}
}

func (b *Board) DrawInterface(screen *ebiten.Image, x, y int) {
	floatX := float32(x)
	floatY := float32(y)

	screen.Fill(mainMenuBgColor)
	glassWidth := BoardWidth * pieceWidth
	glassHeight := BoardHeight * pieceHeight

	dx := floatX + glassWidth
	dy := floatY + glassHeight
	vector.FillRect(screen, floatX, floatY, glassWidth, glassHeight, glassBgColor, false)

	vector.StrokeLine(screen, floatX, floatY, floatX, dy, 1, glassBorderColor, false)
	vector.StrokeLine(screen, floatX, dy, dx, dy, 1, glassBorderColor, false)
	vector.StrokeLine(screen, dx, dy, dx, floatY, 1, glassBorderColor, false)
}

func (b *Board) Draw(screen *ebiten.Image) {
	deltaX, deltaY := 10, 10

	b.DrawInterface(screen, deltaX, deltaY)
	b.DrawCup(screen, deltaX, deltaY)
	b.DrawFigure(screen, deltaX, deltaY)
	b.DrawScore(screen, deltaX, deltaY)
}

func (b *Board) Merge() {
	if b.figure == nil {
		return
	}

	content := b.figure.Content()

	// TODO: это можно оптимизировать
	for x := 0; x < len(content); x++ {
		for y := 0; y < len(content[x]); y++ {
			if content[x][y] == 1 {
				bfx := b.figure.GetX() + x
				bfy := b.figure.GetY() + y
				b.field[bfx][bfy] = 1
			}
		}
	}
	b.figure = nil
}

func (b *Board) RemoveLine(lineNum int) {
	for x := 0; x < len(b.field); x++ {
		for y := lineNum; y > 0; y-- {
			if x > 1 {
				b.field[x][y] = b.field[x][y-1]
			} else {
				b.field[x][y] = 0
			}
		}
	}
}

func (b *Board) RemoveLines() {
	lines := 0

	for y := 0; y < BoardHeight; y++ {
		cnt := 0
		for x := 0; x < BoardWidth; x++ {
			if b.field[x][y] == 1 {
				cnt++
			}
		}
		if cnt == BoardWidth {
			b.RemoveLine(y)
			lines++
			y--
		}
	}
	percent := 0.1*float32(b.level-1) + 1
	b.score = int(float32(b.score) + float32(lines*40)*percent)
	b.score = b.score

}

func (b *Board) Update() error {
	keys := []ebiten.Key{ebiten.KeyRight, ebiten.KeyLeft, ebiten.KeyDown, ebiten.KeySpace, ebiten.KeyUp}
	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			b.lastKey = key
		}
	}

	if b.figure == nil {
		if b.next == nil {
			b.next = NewRandomFigure()
		}

		b.figure = b.next
		b.next = NewRandomFigure()

		center := (BoardWidth - b.figure.Width()) / 2
		b.figure.SetX(center)
	}

	b.ticks++
	if b.ticks < ticksPerMove {
		return nil
	}
	b.ticks = 0
	b.score += b.level

	if b.figureCanMove(moveDown) {
		b.figure.Move(moveDown)
	} else {
		b.Merge()
		b.RemoveLines()
	}

	if b.lastKey == ebiten.KeyLeft {
		b.lastKey = ebiten.KeyA
		if b.figureCanMove(moveLeft) {
			b.figure.Move(moveLeft)
		}
	}

	if b.lastKey == ebiten.KeyRight {
		b.lastKey = ebiten.KeyA
		if b.figureCanMove(moveRight) {
			b.figure.Move(moveRight)
		}
	}

	if b.lastKey == ebiten.KeyUp {
		b.lastKey = ebiten.KeyA
		if b.figure != nil {
			unrotate := b.figure.content

			b.figure.Rotate()
			if !b.figureCanMove(moveNone) {
				b.figure.content = unrotate
			}
		}
	}

	return nil
}

func (b *Board) figureCanMove(mo MoveOffset) bool {
	fx := b.figure.GetX() + mo.x
	fy := b.figure.GetY() + mo.y

	content := b.figure.Content()

	// TODO: это можно оптимизировать
	for x := 0; x < len(content); x++ {
		for y := 0; y < len(content[x]); y++ {
			if content[x][y] == 1 {
				if fy >= BoardHeight-1 {
					return false
				}

				if fx >= BoardWidth-1 {
					return false
				}

				if b.figure.content[x][y] == 1 {
					if fy+y >= BoardHeight || fy+y < 0 {
						return false
					}

					if fx+x >= BoardWidth || fx+x < 0 {
						return false
					}

					if b.field[fx+x][fy+y] == 1 {
						return false
					}

				}

			}
		}
	}
	return true
}
