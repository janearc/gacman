package models

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
)

// Vector3 wraps around mathgl's Vec3 to provide additional functionality.
type Vector3 struct {
	mgl64.Vec3
}

// NewVector3 creates a new Vector3 instance.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{mgl64.Vec3{x, y, z}}
}

// Add adds two vectors and returns the result as a new Vector3.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.Vec3.Add(other.Vec3)}
}

// Subtract subtracts another vector from the current vector and returns the result.
func (v Vector3) Subtract(other Vector3) Vector3 {
	return Vector3{v.Vec3.Sub(other.Vec3)}
}

// Scale scales the vector by a scalar value.
func (v Vector3) Scale(scalar float64) Vector3 {
	return Vector3{v.Vec3.Mul(scalar)}
}

// String provides a string representation of the vector for easy debugging.
func (v Vector3) String() string {
	return fmt.Sprintf("Vector3(%f, %f, %f)", v.X(), v.Y(), v.Z())
}

// Example usage:
// func main() {
//     v1 := NewVector3(1.0, 2.0, 3.0)
//     v2 := NewVector3(4.0, 5.0, 6.0)
//     result := v1.Add(v2)
//     fmt.Println(result) // Outputs: Vector3(5.000000, 7.000000, 9.000000)
// }
