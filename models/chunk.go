package models

import (
	"gacman/core"
	"gacman/types"
)

// Chunk represents a section of the world containing a grid of cells.
type Chunk struct {
	Position core.Vector3          // The world position of this chunk (e.g., its center or corner)
	Size     int                   // The size of the chunk (e.g., number of cells along one axis)
	Cells    map[string]types.Cell // A map of cells within this chunk, keyed by their local coordinates (e.g., "0,0,0")
}

// NewChunk creates a new Chunk with the specified position and size.
func NewChunk(position core.Vector3, size int) Chunk {
	return Chunk{
		Position: position,
		Size:     size,
		Cells:    make(map[string]types.Cell),
	}
}

// AddCell adds a Cell to the chunk at the given local coordinates (e.g., "0,0,0").
func (c *Chunk) AddCell(localCoord string, cell types.Cell) {
	c.Cells[localCoord] = cell
}

// GetCell retrieves a Cell from the chunk by its local coordinates.
func (c *Chunk) GetCell(localCoord string) (types.Cell, bool) {
	cell, exists := c.Cells[localCoord]
	return cell, exists
}
