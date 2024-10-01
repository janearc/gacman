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

// TODO: NewCell
