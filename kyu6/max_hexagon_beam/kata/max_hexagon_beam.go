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
	for y := 0; y < m.diameter(); y++ {
		for x := m.startX(y); x <= m.endX(y); x += 2 {
			cell := Cell{Coords{x, y}, seq[idxSeq], nil, m}
			m.matrix[cell.Coords] = &cell
			idxSeq++
			idxSeq %= len(seq)
		}
	}
}
func (m *HexMatrix) diameter() int {
	return 2*m.n - 1
}

func (m *HexMatrix) startX(y int) int {
	return abs(m.n - y - 1)
}

func (m *HexMatrix) endX(y int) int {
	return 2*(m.diameter()-1) - m.startX(y)
}
func (c *Cell) cellRight() *Cell {
	if result, ok := c.hexMatrix.matrix[Coords{c.x + 2, c.y}]; ok {
		return result
	}
	return nil
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
	for _, dir := range Directions {
		if nextCell := c.next(dir); nextCell != nil {
			if nextCell.beam == nil {
				nextCell.Spread()
			}
		}
	}
	c.beam = &[3]int{c.value, c.value, c.value}
	for _, dir := range Directions {
		if nextCell := c.next(dir); nextCell != nil {
			if nextCell.beam != nil {
				c.beam[dir] += (*nextCell.beam)[dir]
			}
		}
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
