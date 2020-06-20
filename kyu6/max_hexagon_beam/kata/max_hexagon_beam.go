package kata

import (
	"math"
)

type Direction uint8

var offsets = [...]Coords{
	{1, -1},
	{2, 0},
	{1, 1},
	{-1, 1},
	{-2, 0},
	{-1, -1},
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

const NE Direction = 0

func (d Direction) next() Direction {
	return (d + 1) % 6
}

type Coords struct {
	x int
	y int
}

func (c Coords) nextCoords(dir Direction) (co Coords) {
	offset := offsets[dir]
	return Coords{c.x + offset.x, c.y + offset.y}
}

type Cell struct {
	Coords
	value     int64
	hexMatrix *HexMatrix
}

func (c *Cell) nextCoords(dir Direction) (co Coords) {
	return c.Coords.nextCoords(dir)
}

func (c *Cell) next(direction Direction) *Cell {
	coords := c.nextCoords(direction)
	return c.hexMatrix.getCell(coords)
}

func (c *Cell) getBeam(to Direction) (beam int64) {
	next := c.next(to)
	beam = c.value
	if next != nil {
		beam += next.getBeam(to)
	}
	return
}

type HexMatrix struct {
	matrix  map[Coords]*Cell
	n       int
	maxBeam int64
}

func NewCell(x, y int, value int64, hm *HexMatrix) *Cell {
	return &Cell{Coords: Coords{x, y}, value: value, hexMatrix: hm}
}

func NewHexMatrix(n int) *HexMatrix {
	matrix := make(map[Coords]*Cell)
	return &HexMatrix{matrix: matrix, n: n, maxBeam: math.MinInt64}
}

func (hm *HexMatrix) fill(seq []int) {
	idxSeq := 0
	for y := 0; y < 2*hm.n-1; y++ {
		startX := abs(hm.n - y - 1)
		for x := startX; x <= (4*hm.n-4)-startX; x += 2 {
			cell := NewCell(x, y, int64(seq[idxSeq]), hm)
			hm.setCell(cell)
			idxSeq++
			idxSeq %= len(seq)
		}
	}
}

func (hm *HexMatrix) checkMaxBeam(beam int64) {
	if beam > hm.maxBeam {
		hm.maxBeam = beam
	}
}
func (hm *HexMatrix) calculateMaxBeam() int64 {
	initCell := hm.getCell(Coords{1, hm.n - 2})
	cell := initCell
	for nextDir := NE; cell != initCell.next(3); nextDir = nextDir.next() {
		cell = hm.calcBeamBatch(cell, nextDir)
	}
	return hm.maxBeam
}
func (hm *HexMatrix) calcBeamBatch(initCell *Cell, nextDir Direction) (lastCell *Cell) {
	for ; initCell != nil; initCell = initCell.next(nextDir) {
		hm.checkMaxBeam(initCell.getBeam(nextDir.next()))
		lastCell = initCell
	}
	return
}

func (hm *HexMatrix) setCell(cell *Cell) {
	hm.matrix[cell.Coords] = cell
}
func (hm *HexMatrix) getCell(coords Coords) (c *Cell) {
	c, _ = hm.matrix[coords]
	return
}

func MaxHexagonBeam(n int, seq []int) int {
	hm := NewHexMatrix(n)
	hm.fill(seq)
	return int(hm.calculateMaxBeam())
}
