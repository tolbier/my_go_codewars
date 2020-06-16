package kata

import (
	"math"
)

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
)

var Directions = []Direction{Up, Right, Down}

type Coords struct {
	x int
	y int
}
type Cell struct {
	Coords
	value     int
	beam      *[3]int
	hexMatrix *HexMatrix
}
type HexMatrix struct {
	matrix  map[Coords]*Cell
	n       int
	maxBeam int
}

func NewHexMatrix(n int) *HexMatrix {
	matrix := make(map[Coords]*Cell)
	return &HexMatrix{matrix: matrix, n: n}
}
func (m *HexMatrix) fill(seq []int) {
	idxSeq := 0
	for y := 0; y < 2*m.n-1; y++ {
		startX := abs(m.n - y - 1)
		for x := startX; x <= (4*m.n-4)-startX; x += 2 {
			cell := Cell{Coords{x, y}, seq[idxSeq], nil, m}
			m.matrix[cell.Coords] = &cell
			idxSeq++
			idxSeq %= len(seq)
		}
	}
}

func (c *Cell) nextCoords(direction Direction) (co *Coords) {
	var x, y int
	switch direction {
	case Up:
		x, y = c.x+1, c.y-1
	case Right:
		x, y = c.x+2, c.y
	case Down:
		x, y = c.x+1, c.y+1
	}
	return &Coords{x, y}
}
func (c *Cell) next(direction Direction) *Cell {
	coords := *c.nextCoords(direction)
	if result, ok := c.hexMatrix.matrix[coords]; ok {
		return result
	}
	return nil
}
func (c *Cell) Spread() {
	c.beam = &[3]int{}
	for _, dir := range Directions {
		if nextCell := c.next(dir); nextCell != nil {
			if nextCell.beam == nil {
				nextCell.Spread()
			}
		}
		c.beam[dir] = c.value
		if nextCell := c.next(dir); nextCell != nil {
			if nextCell.beam != nil {
				c.setNextBeam(dir, nextCell)
			}
		}
	}
}

func (c *Cell) setNextBeam(dir Direction, nextCell *Cell) {
	c.beam[dir] += (*nextCell.beam)[dir]
	if c.beam[dir] > c.hexMatrix.maxBeam {
		c.hexMatrix.maxBeam = c.beam[dir]
	}
}
func abs(i int) int {
	return int(math.Abs(float64(i)))
}
func MaxHexagonBeam(n int, seq []int) int {
	hm := NewHexMatrix(n)
	hm.fill(seq)
	initCell := hm.matrix[Coords{0, n - 1}]
	initCell.Spread()

	return hm.maxBeam
}
