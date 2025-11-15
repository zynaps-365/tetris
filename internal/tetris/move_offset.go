package tetris

var (
	moveLeft  = MoveOffset{-1, 0}
	moveRight = MoveOffset{1, 0}
	moveDown  = MoveOffset{0, 1}
	moveNone  = MoveOffset{0, 0}
)

type MoveOffset struct {
	x int
	y int
}
