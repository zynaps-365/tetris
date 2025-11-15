package tetris

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	active   Scene
	mainMenu *MainMenu
	board    *Board
}

func (sc *SceneManager) Draw(screen *ebiten.Image) {
	sc.active.Draw(screen)
}

func (sc *SceneManager) Update() error {
	return sc.active.Update()
}

func (sc *SceneManager) SwitchToMainMenu() {
	sc.active = sc.mainMenu
}

func (sc *SceneManager) SwitchGameBoard() {
	sc.active = sc.board
}

func NewSceneManager() *SceneManager {
	sc := &SceneManager{}
	mm := newMainMenu(sc)
	b := newBoard(sc)
	sc.mainMenu = mm
	sc.board = b

	return sc
}
