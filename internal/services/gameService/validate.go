package gameService

import "errors"

var (
	InvalidParams = errors.New("Некорректные параметры игры")
	CellOutOfBounds = errors.New("Клетка вне диапазона")
	ErrGameOver = errors.New("Игра закончена")
)

func validate(width, height, mines int) error {
	total := width * height
	if width < 1 || height < 1 || mines < 1 || mines >= total{
		return InvalidParams
	}
	return nil 
}