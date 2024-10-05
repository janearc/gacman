package core

import (
	"encoding/json"
	"log"
)

// TODO: this should be replaced by the entity type in shapes and
//       i suspect that shapes should probably just go away

// Object represents an entity within a cell.
type Object struct {
	Position    Vector3 `json:"position"`    // Uses Vector3 for position representation
	TerrainType string  `json:"terrainType"` // Type of terrain the object is on
	Height      float64 `json:"height"`      // Height of the object
	ObjectType  string  `json:"objectType"`  // Type of the object (e.g., "stairs")
	IsMovable   bool    `json:"isMovable"`   // Whether the object can be moved
	IsPassable  bool    `json:"isPassable"`  // Whether the object can be passed through
}

// NewObject creates a new Object with the specified properties.
func NewObject(position Vector3, terrainType string, height float64, objectType string, isMovable, isPassable bool) Object {
	return Object{
		Position:    position,
		TerrainType: terrainType,
		Height:      height,
		ObjectType:  objectType,
		IsMovable:   isMovable,
		IsPassable:  isPassable,
	}
}

// ToJSON serializes the Object into a JSON string.
func (o *Object) ToJSON() string {
	jsonData, err := json.Marshal(o)
	if err != nil {
		log.Printf("Error serializing Object to JSON: %v", err)
		return ""
	}
	return string(jsonData)
}
