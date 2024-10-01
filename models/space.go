package models

// where does all this stuff live though

type Space struct {
	Cells map[string]Cell // Use a map for efficient lookup, keyed by cell coordinates (e.g., "0,0,0")
}
