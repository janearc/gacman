package types

// Dungeon represents a collection of levels within the space.
type Dungeon struct {
	Levels       []Level
	CurrentLevel int // Index of the current level in the dungeon
}

// NewDungeon creates a new dungeon and generates an initial level.
func NewDungeon(size int) Dungeon {
	// Generate the first level of the dungeon
	initialLevel := GenerateLevel(size)

	return Dungeon{
		Levels:       []Level{initialLevel}, // Add the generated level to the dungeon
		CurrentLevel: 0,
	}
}

// AddLevel adds a new level to the dungeon.
func (d *Dungeon) AddLevel(size int) {
	newLevel := NewLevel(size)
	d.Levels = append(d.Levels, newLevel)
}

// GetCurrentLevel returns the current active level in the dungeon.
func (d *Dungeon) GetCurrentLevel() Level {
	return d.Levels[d.CurrentLevel]
}

// TraverseStairs moves between levels based on the direction ("up" or "down").
func (d *Dungeon) TraverseStairs(direction string) {
	if direction == "up" && d.CurrentLevel > 0 {
		// Move up to the previous level
		d.CurrentLevel--
	} else if direction == "down" {
		if d.CurrentLevel < len(d.Levels)-1 {
			// Move down to the next existing level
			d.CurrentLevel++
		} else {
			// No more levels, so create a new one and move down to it
			d.AddLevel(10) // Create a new level with a default size of 10
			d.CurrentLevel++
		}
	}
}
