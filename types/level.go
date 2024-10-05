package types

import (
	"gacman/core"
	"gacman/shapes"
	"math/rand"
)

// Room represents a rectangular room on the map, which can contain entities.
type Room struct {
	X, Y          int // Top-left corner of the room
	Width, Height int
	Entities      []shapes.Entity // Entities present in this room
}

// Level represents a single level in the dungeon.
type Level struct {
	Cells      map[string]Cell // Cells within this level
	StairsUp   shapes.Stairs   // The "up" stairs object
	StairsDown shapes.Stairs   // The "down" stairs object
	Rooms      []Room          // List of rooms in the level
}

// NewLevel creates a new level with stairs.
func NewLevel(size int) Level {
	// Create an empty map to hold cells for this level
	cells := createEmptyCells(size)

	// Place "up" and "down" stairs in the level
	stairsUp, stairsDown := placeStairs(cells, size)

	return Level{
		Cells:      cells,
		StairsUp:   stairsUp,
		StairsDown: stairsDown,
	}
}

// roomOverlaps checks if a room overlaps with existing rooms.
func roomOverlaps(rooms []Room, newRoom Room) bool {
	for _, room := range rooms {
		if newRoom.X < room.X+room.Width && newRoom.X+newRoom.Width > room.X &&
			newRoom.Y < room.Y+room.Height && newRoom.Y+newRoom.Height > room.Y {
			return true
		}
	}
	return false
}

// fillRoom changes cells to be part of the room.
func fillRoom(cells map[string]Cell, room Room) {
	for x := room.X; x < room.X+room.Width; x++ {
		for y := room.Y; y < room.Y+room.Height; y++ {
			position := core.NewVector3(float64(x), float64(y), 0)
			coord := core.GetCoordString(x, y)
			cells[coord] = NewCell(position, "floor", 0)
		}
	}
}

// connectRooms creates corridors between rooms.
func connectRooms(cells map[string]Cell, rooms []Room) {
	for i := 0; i < len(rooms)-1; i++ {
		roomA := rooms[i]
		roomB := rooms[i+1]

		// Get the center of each room
		ax, ay := roomA.X+roomA.Width/2, roomA.Y+roomA.Height/2
		bx, by := roomB.X+roomB.Width/2, roomB.Y+roomB.Height/2

		// Create corridors by digging horizontal and vertical paths
		if rand.Intn(2) == 0 {
			createHorizontalCorridor(cells, ax, bx, ay)
			createVerticalCorridor(cells, ay, by, bx)
		} else {
			createVerticalCorridor(cells, ay, by, ax)
			createHorizontalCorridor(cells, ax, bx, by)
		}
	}
}

// createHorizontalCorridor creates a horizontal path between two x-coordinates.
func createHorizontalCorridor(cells map[string]Cell, x1, x2, y int) {
	for x := core.Min(x1, x2); x <= core.Max(x1, x2); x++ {
		position := core.NewVector3(float64(x), float64(y), 0)
		coord := core.GetCoordString(x, y)
		cells[coord] = NewCell(position, "floor", 0)
	}
}

// createVerticalCorridor creates a vertical path between two y-coordinates.
func createVerticalCorridor(cells map[string]Cell, y1, y2, x int) {
	for y := core.Min(y1, y2); y <= core.Max(y1, y2); y++ {
		position := core.NewVector3(float64(x), float64(y), 0)
		coord := core.GetCoordString(x, y)
		cells[coord] = NewCell(position, "floor", 0)
	}
}

// createEmptyCells initializes an empty map of cells for the level.
func createEmptyCells(size int) map[string]Cell {
	cells := make(map[string]Cell)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			position := core.NewVector3(float64(x), float64(y), 0)
			cell := NewCell(position, "empty", 0)
			coord := core.GetCoordString(x, y)
			cells[coord] = cell
		}
	}
	return cells
}

// placeStairs randomly places "up" and "down" stairs in the level's cells.
func placeStairs(cells map[string]Cell, size int) (shapes.Stairs, shapes.Stairs) {
	// Generate random coordinates for the "up" and "down" stairs
	upX, upY := rand.Intn(size), rand.Intn(size)
	downX, downY := rand.Intn(size), rand.Intn(size)

	// Create the "up" stairs
	stairsUp := shapes.NewStairs("up")
	upPosition := core.NewVector3(float64(upX), float64(upY), 0)
	upCell := NewCell(upPosition, stairsUp.ObjectType, 0)
	upCoord := core.GetCoordString(upX, upY)
	cells[upCoord] = upCell

	// Create the "down" stairs
	stairsDown := shapes.NewStairs("down")
	downPosition := core.NewVector3(float64(downX), float64(downY), 0)
	downCell := NewCell(downPosition, stairsDown.ObjectType, 0)
	downCoord := core.GetCoordString(downX, downY)
	downCell.Neighbors = core.GetNeighborPositions(downX, downY)
	cells[downCoord] = downCell

	// Return the stairs objects
	return stairsUp, stairsDown
}

// placeStairsInRooms randomly places "up" and "down" stairs in the rooms.
func placeStairsInRooms(cells map[string]Cell, rooms []Room) (shapes.Stairs, shapes.Stairs) {
	upRoom := rooms[rand.Intn(len(rooms))]
	downRoom := rooms[rand.Intn(len(rooms))]

	stairsUp := NewStairsInRoom(upRoom, "up", cells)
	stairsDown := NewStairsInRoom(downRoom, "down", cells)

	return stairsUp, stairsDown
}

// NewStairsInRoom places stairs in a specified room.
func NewStairsInRoom(room Room, direction string, cells map[string]Cell) shapes.Stairs {
	stairs := shapes.NewStairs(direction)
	stairsX := room.X + rand.Intn(room.Width)
	stairsY := room.Y + rand.Intn(room.Height)
	position := core.NewVector3(float64(stairsX), float64(stairsY), 0)
	coord := core.GetCoordString(stairsX, stairsY)
	cells[coord] = NewCell(position, stairs.ObjectType, 0)
	return stairs
}

// createRooms randomly places rooms on the map and initializes their entities.
func createRooms(cells map[string]Cell, numRooms, size int) []Room {
	rooms := []Room{}
	for i := 0; i < numRooms; i++ {
		// Random room size
		width := rand.Intn(10) + 5 // Room width between 5 and 15
		height := rand.Intn(6) + 4 // Room height between 4 and 10

		// Ensure the room size does not exceed the level's size
		if width >= size {
			width = size - 1
		}
		if height >= size {
			height = size - 1
		}

		// Ensure that x and y are within valid ranges
		maxX := size - width - 1
		maxY := size - height - 1

		if maxX <= 0 || maxY <= 0 {
			// Skip creating this room if it doesn't fit within the level's bounds
			continue
		}

		// Random room position
		x := rand.Intn(maxX)
		y := rand.Intn(maxY)

		newRoom := Room{
			X:        x,
			Y:        y,
			Width:    width,
			Height:   height,
			Entities: []shapes.Entity{}, // Initialize the entities slice
		}

		// Add room to the map if it doesn't overlap
		if !roomOverlaps(rooms, newRoom) {
			rooms = append(rooms, newRoom)
			fillRoom(cells, newRoom)
		}
	}
	return rooms
}

// GenerateLevel creates a Nethack-like random level with rooms and corridors.
func GenerateLevel(size int) Level {
	// Create an empty grid of walls
	cells := make(map[string]Cell)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			position := core.NewVector3(float64(x), float64(y), 0)
			cells[core.GetCoordString(x, y)] = NewCell(position, "wall", 1) // Default cells are walls
		}
	}

	// Randomly create rooms
	numRooms := rand.Intn(10) + 5 // Random number of rooms between 5 and 15
	rooms := createRooms(cells, numRooms, size)

	// Connect rooms with corridors
	connectRooms(cells, rooms)

	// Place "up" and "down" stairs in random rooms
	stairsUp, stairsDown := placeStairsInRooms(cells, rooms)

	// Return the level with cells, stairs, and rooms
	return Level{
		Cells:      cells,
		StairsUp:   stairsUp,
		StairsDown: stairsDown,
		Rooms:      rooms,
	}
}
