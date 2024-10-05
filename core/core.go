package core

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// You might call this "primitives" or "math", but this is the types and functions
// that are the base classes (again, imprecise word) for lots of other things this
// software uses.

// Vector3 wraps around mathgl's Vec3 to provide additional functionality.
type Vector3 struct {
	mgl64.Vec3
}

// Quaternion represents a rotation in 3D space.
type Quaternion struct {
	X, Y, Z, W float64
}

// NewVector3 creates a new Vector3 instance.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{mgl64.Vec3{x, y, z}}
}

// GetCoordString converts coordinates to a string key (e.g., "x,y").
func GetCoordString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// GetNeighborPositions returns a slice of Vector3 positions representing neighboring cells.
func GetNeighborPositions(x, y int) []Vector3 {
	return []Vector3{
		NewVector3(float64(x+1), float64(y), 0),
		NewVector3(float64(x-1), float64(y), 0),
		NewVector3(float64(x), float64(y+1), 0),
		NewVector3(float64(x), float64(y-1), 0),
	}
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

// NewQuaternion creates a new quaternion given x, y, z, and w components.
func NewQuaternion(x, y, z, w float64) Quaternion {
	return Quaternion{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// IdentityQuaternion returns the identity quaternion (no rotation).
func IdentityQuaternion() Quaternion {
	return Quaternion{
		X: 0,
		Y: 0,
		Z: 0,
		W: 1,
	}
}

// QuaternionFromEuler creates a quaternion from Euler angles (in degrees).
func QuaternionFromEuler(pitch, yaw, roll float64) Quaternion {
	// Convert angles from degrees to radians
	pitch = pitch * (math.Pi / 180)
	yaw = yaw * (math.Pi / 180)
	roll = roll * (math.Pi / 180)

	c1 := math.Cos(yaw / 2)
	c2 := math.Cos(pitch / 2)
	c3 := math.Cos(roll / 2)
	s1 := math.Sin(yaw / 2)
	s2 := math.Sin(pitch / 2)
	s3 := math.Sin(roll / 2)

	return Quaternion{
		X: s1*s2*c3 + c1*c2*s3,
		Y: s1*c2*c3 + c1*s2*s3,
		Z: c1*s2*c3 - s1*c2*s3,
		W: c1*c2*c3 - s1*s2*s3,
	}
}

// Normalize normalizes the quaternion to unit length.
func (q *Quaternion) Normalize() {
	length := math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
	if length > 0 {
		q.X /= length
		q.Y /= length
		q.Z /= length
		q.W /= length
	}
}
