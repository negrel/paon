package geometry

type privateVec2D = Vec2D

// Vec3D define a point in 3 dimensional Euclidian space.
type Vec3D struct {
	privateVec2D

	z int
}

// NewVec3D returns a new 3D vector with the given values.
func NewVec3D(x, y, z int) Vec3D {
	return Vec3D{
		privateVec2D: privateVec2D{
			x: x,
			y: y,
		},
		z: z,
	}
}

// Z returns the location of the vector on the Z axis.
func (v3 Vec3D) Z() int {
	return v3.z
}

// Add returns a new Vec3D translated by adding the given vector.
func (v3 Vec3D) Add(other Vec3D) Vec3D {
	v3.x += other.x
	v3.y += other.y
	v3.z += other.z

	return v3
}

// Sub returns a new Vec3D translated by subtracting the given vector.
func (v3 Vec3D) Sub(other Vec3D) Vec3D {
	v3.x -= other.x
	v3.y -= other.y
	v3.z -= other.z

	return v3
}

// Equals returns true if the given vector is equal to this Vec3DD.
func (v3 Vec3D) Equals(other Vec3D) bool {
	return v3.privateVec2D.Equals(other.privateVec2D) && v3.z == other.z
}
