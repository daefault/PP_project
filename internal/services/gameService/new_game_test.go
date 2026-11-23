package gameService

import (
	"errors"
	"math/rand"
	"testing"
)

func newTestGame(t *testing.T, width, height, mines int) *Game {
	t.Helper()
	g, err := NewGame(width, height, mines, rand.New(rand.NewSource(42)))
	if err != nil {
		t.Fatalf("ошибка = %v", err)
	}
	return g
}

func TestInvalidParameters(t *testing.T) {
	cases := []struct {
		name string
		width int
		height int
		mines int
	}{
		{"нулевая ширина", 0, 10, 5},
		{"нулевая высота", 10, 0, 5},
		{"0 мин", 10, 10, 0},
		{"мин слишком много", 3, 3, 9},
	}
	for _, test_case := range cases {
		t.Run(test_case.name, func(t * testing.T) {
			_, err := NewGame(test_case.width, test_case.height, test_case.mines, rand.New(rand.NewSource(42)))
			if !errors.Is(err, InvalidParams) {
				t.Errorf("ожидалась ошибка InvalidParams, была получена %v", err)
			}
		})
	}
}

func TestPlacesMines(t *testing.T) {
	g := newTestGame(t, 9, 9, 10)
	count := 0
	for _, c := range g.Cells{
		if c.Mine{
			count++
		}
	}
	if count != 10{
		t.Errorf("ожидалось 10 мин, создано %d", count)
	}
}