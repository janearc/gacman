package types

import (
	"encoding/json"
	"fmt"
	"gacman/core"
)

// Cell represents a discrete unit of space in Unity.
type Cell struct {
	Position    core.Vector3   // Center position of the cell in Unity space
	TerrainType string         // e.g., "grass," "water," "mountain"
	Height      float64        // Elevation of the cell
	IsOccupied  bool           // Whether something is occupying this cell
	ObjectType  string         // Type of object, if any (e.g., "tree," "rock")
	Neighbors   []core.Vector3 // Adjacent cell positions for logical connectivity
}

// NewCell creates a new Cell with the specified position, terrain type, and height.
func NewCell(position core.Vector3, terrainType string, height float64) Cell {
	return Cell{
		Position:    position,
		TerrainType: terrainType,
		Height:      height,
		IsOccupied:  false, // Default to not occupied
		ObjectType:  "",    // No object by default
		Neighbors:   []core.Vector3{},
	}
}

// GenerateNewCell creates a new cell in the given direction relative to the current cell.
func GenerateNewCell(current Cell, direction string) Cell {
	var newX, newY int

	// Convert the current position coordinates to integers
	currentX := int(current.Position.X())
	currentY := int(current.Position.Y())

	switch direction {
	case "n":
		newX, newY = currentX, currentY+1
	case "s":
		newX, newY = currentX, currentY-1
	case "e":
		newX, newY = currentX+1, currentY
	case "w":
		newX, newY = currentX-1, currentY
	}

	// Create a new Vector3 for the new cell's position
	position := core.NewVector3(float64(newX), float64(newY), 0)

	// Create the new cell and set its neighbors
	newCell := NewCell(position, "empty", 0)
	newCell.Neighbors = core.GetNeighborPositions(newX, newY)

	return newCell
}

// ToJSON serializes the Cell into a JSON string.
func (c *Cell) ToJSON() string {
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error serializing Cell to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}
