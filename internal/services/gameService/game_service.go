package gameService

import "math/rand"


type CellState int 
type GameStatus int 

const (
	CellClosed CellState = iota
	CellOpen
	CellFlagged
)
const (
	StatusInProgress GameStatus = iota
	StatusWon
	StatusLost
)

type Cell struct {
	Mine bool
	MinesNear int
	State CellState
}

type Game struct {
	Width int
	Height int
	Mines int 
	Cells []Cell
	Status GameStatus
	num_generator *rand.Rand
}

func NewGame(width, height, mines int, num_generator *rand.Rand) (*Game, error){
	if err:= validate(width, height, mines); err!= nil{
		return nil, err
	}
	g:= &Game{
		Width: width,
		Height: height,
		Mines: mines,
		Cells: make([]Cell, width*height),
		Status: StatusInProgress,
		num_generator: num_generator,
	}
	g.placeMines()

	return g, nil
}

