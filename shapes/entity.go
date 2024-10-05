package shapes

import (
	"gacman/core"
)

// Entity represents an object in the game world that Unity will render.
type Entity struct {
	ID       string            `json:"id"`       // Unique identifier
	Name     string            `json:"name"`     // Human-readable name
	Position core.Vector3      `json:"position"` // Position in the game world
	Rotation core.Quaternion   `json:"rotation"` // Orientation (could use Euler angles if simpler)
	Scale    core.Vector3      `json:"scale"`    // Size scaling of the entity
	Type     string            `json:"type"`     // Type (e.g., "monster", "item", "player")
	Health   int               `json:"health"`   // Health for damageable entities
	MeshName string            `json:"meshName"` // Reference to the model/mesh in Unity
	IsActive bool              `json:"isActive"` // Whether the entity is active in the game world
	Metadata map[string]string `json:"metadata"` // Additional custom properties
}

// NewEntity creates a new game entity with default values.
func NewEntity(id, name, entityType, meshName string, position, scale core.Vector3, rotation core.Quaternion) Entity {
	return Entity{
		ID:       id,
		Name:     name,
		Position: position,
		Rotation: rotation,
		Scale:    scale,
		Type:     entityType,
		MeshName: meshName,
		Health:   100,  // Default health
		IsActive: true, // Default to active
		Metadata: make(map[string]string),
	}
}
