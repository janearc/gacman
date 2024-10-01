package models

import (
	"encoding/json"
	"log"
)

// Object represents an entity within a cell.
type Object struct {
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	} `json:"position"`
	TerrainType string  `json:"terrainType"`
	Height      float64 `json:"height"`
	ObjectType  string  `json:"objectType"`
}

// NewObject creates a new Object with the specified position, terrain type, height, and object type.
func NewObject(x, y, z int, terrainType string, height float64, objectType string) Object {
	return Object{
		Position: struct {
			X int `json:"x"`
			Y int `json:"y"`
			Z int `json:"z"`
		}{
			X: x,
			Y: y,
			Z: z,
		},
		TerrainType: terrainType,
		Height:      height,
		ObjectType:  objectType,
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
