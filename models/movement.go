package models

import (
	"gacman/core"
	"gacman/types"
)

// GenerateNewCell attempts to generate a new cell in the specified direction.
// It returns the new cell and a status message indicating the result of the movement.
func GenerateNewCell(currentCell types.Cell, direction string, space *Space) (types.Cell, string) {
	// Determine the new coordinates based on the direction
	newX, newY := int(currentCell.Position.X()), int(currentCell.Position.Y())

	switch direction {
	case "n":
		newY--
	case "s":
		newY++
	case "e":
		newX++
	case "w":
		newX--
	default:
		return currentCell, "Invalid direction"
	}

	// Get the new cell's coordinates
	newCoord := core.GetCoordString(newX, newY)

	// Check if the cell exists in the space
	if newCell, exists := space.GetCell(newCoord); exists {
		// Check if the new cell is a wall
		if newCell.TerrainType == "wall" {
			return currentCell, "Movement blocked: You can't walk through walls!"
		}
		// Movement successful
		return newCell, "Movement successful"
	}

	// If the new cell doesn't exist in the space, return the current cell
	return currentCell, "Movement blocked: You can't move outside the map!"
}
