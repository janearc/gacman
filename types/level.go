package types

import (
	"gacman/core"
	"gacman/shapes"
	"math/rand"
)

// Level represents a single level in the dungeon.
type Level struct {
	Cells      map[string]Cell // Cells within this level
	StairsUp   shapes.Stairs   // The "up" stairs object
	StairsDown shapes.Stairs   // The "down" stairs object
}

// NewLevel creates a new level with stairs.
func NewLevel(size int) Level {
	// Create an empty map to hold cells for this level
	cells := createEmptyCells(size)

	// Place "up" and "down" stairs in the level
	stairsUp, stairsDown := placeStairs(cells, size)

	return Level{
		Cells:      cells,
		StairsUp:   stairsUp,
		StairsDown: stairsDown,
	}
}

// createEmptyCells initializes an empty map of cells for the level.
func createEmptyCells(size int) map[string]Cell {
	cells := make(map[string]Cell)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			position := core.NewVector3(float64(x), float64(y), 0)
			cell := NewCell(position, "empty", 0)
			coord := core.GetCoordString(x, y)
			cells[coord] = cell
		}
	}
	return cells
}

// placeStairs randomly places "up" and "down" stairs in the level's cells.
func placeStairs(cells map[string]Cell, size int) (shapes.Stairs, shapes.Stairs) {
	// Generate random coordinates for the "up" and "down" stairs
	upX, upY := rand.Intn(size), rand.Intn(size)
	downX, downY := rand.Intn(size), rand.Intn(size)

	// Create the "up" stairs
	stairsUp := shapes.NewStairs("up")
	upPosition := core.NewVector3(float64(upX), float64(upY), 0)
	upCell := NewCell(upPosition, stairsUp.ObjectType, 0)
	upCoord := core.GetCoordString(upX, upY)
	cells[upCoord] = upCell

	// Create the "down" stairs
	stairsDown := shapes.NewStairs("down")
	downPosition := core.NewVector3(float64(downX), float64(downY), 0)
	downCell := NewCell(downPosition, stairsDown.ObjectType, 0)
	downCoord := core.GetCoordString(downX, downY)
	downCell.Neighbors = core.GetNeighborPositions(downX, downY)
	cells[downCoord] = downCell

	// Return the stairs objects
	return stairsUp, stairsDown
}
