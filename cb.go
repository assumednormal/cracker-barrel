package cb

import "fmt"

// Position represents a position on the playing board
type Position struct {
	X, Y int
}

// Move represents a move on the playing board
type Move struct {
	Start, End, Drop *Position
}

func (m *Move) String() string {
	return fmt.Sprintf("Move:\t\nStart:\t%v\nDrop:\t%v\nEnd:\t%v\n", m.Start, m.Drop, m.End)
}

// Game represents a game
type Game struct {
	Empty  []*Position
	Filled []*Position
}

// NewGame returns a new board with all positions filled except one corner
func NewGame() *Game {
	// all positions except one corner
	empty := []*Position{&Position{5, 5}}

	filled := []*Position{
		&Position{4, 4}, &Position{6, 4},
		&Position{3, 3}, &Position{5, 3}, &Position{7, 3},
		&Position{2, 2}, &Position{4, 2}, &Position{6, 2}, &Position{8, 2},
		&Position{1, 1}, &Position{3, 1}, &Position{5, 1}, &Position{7, 1}, &Position{9, 1},
	}

	return &Game{empty, filled}
}

// AvailableMoves lists available moves
func (g *Game) AvailableMoves() []*Move {
	moves := []*Move{}

	// find neighboring filled positions
	for _, start := range g.Filled {
		for _, drop := range g.Filled {
			if (start.Y == drop.Y && abs(start.X-drop.X) == 2) || (abs(start.X-drop.X) == 1 && abs(start.Y-drop.Y) == 1) {
				xDiff := drop.X - start.X
				yDiff := drop.Y - start.Y
				end := &Position{drop.X + xDiff, drop.Y + yDiff}
				if g.isEmpty(end.X, end.Y) {
					moves = append(moves, &Move{start, end, drop})
				}
			}
		}
	}

	return moves
}

// PlayMove plays a given move
func (g *Game) PlayMove(move *Move) {
	// remove start and drop from filled list
	var idx int
	for i, filled := range g.Filled {
		if filled.X == move.Start.X && filled.Y == move.Start.Y {
			idx = i
		}
	}
	g.Filled = append(g.Filled[:idx], g.Filled[(idx+1):]...)

	for i, filled := range g.Filled {
		if filled.X == move.Drop.X && filled.Y == move.Drop.Y {
			idx = i
		}
	}
	g.Filled = append(g.Filled[:idx], g.Filled[(idx+1):]...)

	// add start and drop to empty list
	g.Empty = append(g.Empty, move.Start, move.Drop)

	// remove end from empty list
	for i, empty := range g.Empty {
		if empty.X == move.End.X && empty.Y == move.End.Y {
			idx = i
		}
	}
	g.Empty = append(g.Empty[:idx], g.Empty[(idx+1):]...)

	// add end to filled list
	g.Filled = append(g.Filled, move.End)
}

func (g *Game) isEmpty(x, y int) bool {
	for _, p := range g.Empty {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
