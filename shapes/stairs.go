package shapes

import "gacman/core"

// Stairs represents a special type of shape for moving between levels.
type Stairs struct {
	Shape
	Direction string // Direction of the stairs: "up" or "down"
}

// NewStairs creates a new Stairs shape with the specified direction ("up" or "down").
func NewStairs(direction string) Stairs {
	// Set specific properties for stairs
	objectType := "stairs_" + direction
	position := core.NewVector3(0, 0, 0) // Default position, to be set when added to a cell

	return Stairs{
		Shape:     NewShape(position, objectType, false, true, objectType),
		Direction: direction,
	}
}
