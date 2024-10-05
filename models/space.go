package models

import (
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

// InitSpace initializes a new Space with a starting dungeon.
func InitSpace(size int) Space {
	// Create a new space
	space := NewSpace()

	// Create and add the initial dungeon
	initialDungeon := types.NewDungeon(size)
	space.AddDungeon(initialDungeon)

	// Return the initialized space
	return space
}
