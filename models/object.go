package models

type Object struct {
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	} `json:"position"`
	TerrainType string  `json:"terrainType"`
	Height      float64 `json:"height"`
	ObjectType  string  `json:"objectType"`
}

// TODO: NewObject
//       Object.ToJson
