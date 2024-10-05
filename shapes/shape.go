package shapes

import "gacman/core"

// Shape represents any object that can take on a shape within the game world, wrapping a core.Object.
type Shape struct {
	core.Object              // Embed core.Object
	SpecificShapeType string // Additional property to define more specific shape details (e.g., "stairs_up", "stairs_down")
}

// NewShape creates a new Shape based on the core.Object with additional shape-specific properties.
func NewShape(position core.Vector3, objectType string, isMovable, isPassable bool, specificShapeType string) Shape {
	return Shape{
		Object: core.Object{
			Position:   position,
			ObjectType: objectType,
			IsMovable:  isMovable,
			IsPassable: isPassable,
		},
		SpecificShapeType: specificShapeType,
	}
}
