package models

// discrete units of space in unity are called cells

type Cell struct {
	Position    Vector3   // Center position of the cell in Unity space
	TerrainType string    // e.g., "grass," "water," "mountain"
	Height      float64   // Elevation of the cell
	IsOccupied  bool      // Whether something is occupying this cell
	ObjectType  string    // Type of object, if any (e.g., "tree," "rock")
	Neighbors   []Vector3 // Adjacent cell positions for logical connectivity
}

// NewCell creates a new Cell with the specified position, terrain type, and height.
func NewCell(position Vector3, terrainType string, height float64) Cell {
	return Cell{
		Position:    position,
		TerrainType: terrainType,
		Height:      height,
		IsOccupied:  false, // Default to not occupied
		ObjectType:  "",    // No object by default
	}
}
