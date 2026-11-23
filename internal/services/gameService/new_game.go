package gameService

func (g *Game) index(x, y int) int {
	return y*g.Width + x
}

func (g *Game) placeMines() {
	total := g.Width * g.Height
	placement := g.num_generator.Perm(total)
	for _, pos:= range placement[:g.Mines]{
		g.Cells[pos].Mine = true
	}
	for y :=0; y< g.Height; y++ {
		for x:= 0; x < g.Width; x++{
			if g.Cells[g.index(x,y)].Mine {
				continue
			}
			g.Cells[g.index(x,y)].MinesNear = g.countNeighborMines(x, y)
		}
	}
}

func (g *Game) countNeighborMines(x, y int) int {
	count := 0
	for dy := -1; dy <=1; dy++{
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0{
				continue
			}
			neighbor_x, neighbor_y := x + dx, y + dy
			if neighbor_x < 0 || neighbor_y < 0 || neighbor_x >= g.Width || neighbor_y >= g.Height {
				continue
			} 
			if g.Cells[g.index(neighbor_x, neighbor_y)].Mine {
				count++
			}
		}
	}
	return count
}
