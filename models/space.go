package models

import (
	"gacman/core"
	"gacman/types"
	"math/rand"
	"time"
)

// Global random number generator instance
var rng *rand.Rand

// Initialize the random number generator
func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Space represents the game world containing cells, chunks, and dungeons.
type Space struct {
	Cells    map[string]types.Cell
	Chunks   map[string]Chunk
	Dungeons []types.Dungeon
}

// NewSpace creates a new Space instance with empty maps for cells, chunks, and dungeons.
func NewSpace() Space {
	return Space{
		Cells:    make(map[string]types.Cell),
		Chunks:   make(map[string]Chunk),
		Dungeons: []types.Dungeon{},
	}
}

// AddCell adds a cell to the space at the specified coordinates.
func (s *Space) AddCell(coord string, cell types.Cell) {
	s.Cells[coord] = cell
}

// GetCell retrieves a cell from the space by its coordinates.
func (s *Space) GetCell(coord string) (types.Cell, bool) {
	cell, exists := s.Cells[coord]
	return cell, exists
}

// AddDungeon adds a dungeon to the space.
func (s *Space) AddDungeon(dungeon types.Dungeon) {
	s.Dungeons = append(s.Dungeons, dungeon)
}

// InitSpace initializes a new Space with a starting dungeon and returns the starting coordinates.
func InitSpace(size int) (Space, string) {
	// Create a new space
	space := NewSpace()

	// Create a new dungeon with an initial level
	dungeon := types.NewDungeon(size)

	// Add the dungeon to the space
	space.AddDungeon(dungeon)

	// Add the dungeon's first level's cells to the space's cells for easy access
	for coord, cell := range dungeon.GetCurrentLevel().Cells {
		space.AddCell(coord, cell)
	}

	// Randomly select a starting position within one of the rooms
	rooms := dungeon.GetCurrentLevel().Rooms
	var startingCoord string
	if len(rooms) > 0 {
		// Select a random room
		randomRoom := rooms[rand.Intn(len(rooms))]

		// Choose a random position within the room
		startX := randomRoom.X + rand.Intn(randomRoom.Width)
		startY := randomRoom.Y + rand.Intn(randomRoom.Height)

		startingCoord = core.GetCoordString(startX, startY)

		// Check if the cell exists and is a floor cell
		if cell, exists := space.GetCell(startingCoord); exists && cell.TerrainType == "floor" {
			return space, startingCoord
		}
	}

	// Fallback: If no valid starting point is found, set a default coordinate
	startingCoord = core.GetCoordString(0, 0)
	return space, startingCoord
}
