package models

import (
	"strings"
)

// RenderASCIIMap generates an ASCII representation of the current level.
func RenderASCIIMap(space Space, size int) string {
	// Create a 2D slice to represent the map
	mapGrid := make([][]rune, size)
	for i := range mapGrid {
		mapGrid[i] = make([]rune, size)
		for j := range mapGrid[i] {
			mapGrid[i][j] = '#' // Default to walls
		}
	}

	// Iterate over the cells and update the map representation
	for _, cell := range space.Cells {
		x, y := int(cell.Position.X()), int(cell.Position.Y())
		if x >= 0 && x < size && y >= 0 && y < size {
			switch cell.TerrainType {
			case "floor":
				mapGrid[y][x] = '.'
			case "stairs_up":
				mapGrid[y][x] = '^'
			case "stairs_down":
				mapGrid[y][x] = 'v'
			}
		}
	}

	// Convert the 2D grid into a single string for display
	var sb strings.Builder
	for _, row := range mapGrid {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteRune('\n') // Newline after each row
	}

	return sb.String()
}
