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

// ToJSON serializes the Cell into a JSON string.
func (c *Cell) ToJSON() string {
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error serializing Cell to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}
