package models

// where does all this stuff live though

type Space struct {
	Cells  map[string]Cell  // Use a map for efficient lookup, keyed by cell coordinates (e.g., "0,0,0")
	Chunks map[string]Chunk // A map of chunks, keyed by their coordinates (e.g., "0,0,0")
}

// NewSpace creates a new Space instance with an empty map of chunks.
func NewSpace() Space {
	return Space{
		Chunks: make(map[string]Chunk),
	}
}

// AddChunk adds a chunk to the space at the specified coordinates.
func (s *Space) AddChunk(coord string, chunk Chunk) {
	s.Chunks[coord] = chunk
}

// GetChunk retrieves a chunk from the space by its coordinates.
func (s *Space) GetChunk(coord string) (Chunk, bool) {
	chunk, exists := s.Chunks[coord]
	return chunk, exists
}
