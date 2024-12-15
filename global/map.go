package global

import (
	"fmt"
	"strconv"
	"strings"
)

////////////////////
// n-dimensional point

// Represents a n-dimensional point
type Point struct {
	Coordinates []int
}

// Helper function to convert Point to a string key for use in `Map`
func (p Point) String() string {
	return fmt.Sprint(p.Coordinates) // Example: "[1 2 3]"
}

func PointFromString(str string) (p Point) {
	p.Coordinates = []int{}
	str = str[1 : len(str)-1]
	values := strings.Split(str, " ")
	for _, v := range values {
		if x, err := strconv.Atoi(v); err != nil {
			panic(err)
		} else {
			p.Coordinates = append(p.Coordinates, x)
		}
	}
	return p
}

func IsSamePoint(p1 Point, p2 Point) bool {
	if len(p1.Coordinates) != len(p2.Coordinates) {
		panic("Cannot compare points with different dimensions")
	}
	for i := 0; i < len(p1.Coordinates); i++ {
		if p1.Coordinates[i] != p2.Coordinates[i] {
			return false
		}
	}

	return true
}

func Distance(first Point, second Point) (v Vector) {
	if len(first.Coordinates) != len(second.Coordinates) {
		panic("Cannot compare points with different dimensions")
	}
	for i := 0; i < len(first.Coordinates); i++ {
		v.Deltas = append(v.Deltas, second.Coordinates[i]-first.Coordinates[i])
	}
	return
}

////////////////////
// n-dimensional vector

// Represents a n-dimensional vector
type Vector struct {
	Deltas []int
}

// Helper function to convert Vector to a string
func (v Vector) String() string {
	return fmt.Sprint(v.Deltas)
}

// Add another vector to the current vector
func (a Vector) Add(b Vector) (Vector, error) {
	if len(a.Deltas) != len(b.Deltas) {
		return Vector{}, fmt.Errorf("Dimension mismatch: a=%d vs. b=%d", len(a.Deltas), len(b.Deltas))
	}

	newVector := Vector{Deltas: make([]int, len(a.Deltas), len(a.Deltas))}
	for i := 0; i < len(a.Deltas); i++ {
		newVector.Deltas[i] = a.Deltas[i] + b.Deltas[i]
	}
	return newVector, nil
}

// Move a point in the direction of the vector, by the value specified in each delta of the vector
func Move(p Point, v Vector) (Point, error) {
	if len(p.Coordinates) != len(v.Deltas) {
		return Point{}, fmt.Errorf("Dimension mismatch: p=%d vs. v=%d", len(p.Coordinates), len(v.Deltas))
	}

	newPoint := Point{Coordinates: make([]int, len(p.Coordinates), len(p.Coordinates))}
	for i := 0; i < len(p.Coordinates); i++ {
		newPoint.Coordinates[i] = p.Coordinates[i] + v.Deltas[i]
	}
	return newPoint, nil
}

////////////////////
// n-dimensional Map

// Represents a n-dimensional map
type Map[T any] struct {
	Positions  map[string]T
	Dimensions []int
}

// Build a map with specified dimensions with values for each point in the map.
func (m *Map[T]) Build(dimensions []int, values []T) {
	// Calculate the total number of positions
	totalEntries := 1
	for _, dim := range dimensions {
		totalEntries *= dim
	}
	// Ensure the values length matches the required number of positions
	if len(values) != totalEntries {
		panic(fmt.Sprintf("Number of values (%d) does not match the number of points (%d) on the map", len(values), totalEntries))
	}

	m.Dimensions = dimensions
	m.Positions = make(map[string]T)

	// Helper function to recursively generate points for the map.
	var generatePoints func([]int, int, int)
	generatePoints = func(currentPoint []int, dimensionIndex int, valueIndex int) {

		// When values on all axes are set for a point, store the value at that point
		if dimensionIndex < 0 {
			// Create the point by copying the current point values
			point := Point{Coordinates: append([]int{}, currentPoint...)}
			m.Positions[point.String()] = values[valueIndex]
			return
		}

		// Iterate over the length of current dimension
		for i := 0; i < dimensions[dimensionIndex]; i++ {
			// Lock the position in the current dimension and fill up the points for all the remaining lower dimensions.
			// Example: For a 3D map, start with locking z=0, and fill up the XY plane. In XY plane, lock y=0 and fill up X line.
			// valueIndex is calculated by using any calculated offset (the passed valueIndex) from next higher dimension, multiplying it by the length of current dimension, to arrive at the start of the current dimension line and finally adding the index in the current dimension. All this is done so that we can exactly get to the location in the array where value for the current point will be assigned.
			generatePoints(append([]int{i}, currentPoint...), dimensionIndex-1, valueIndex*dimensions[dimensionIndex]+i)
		}
	}

	// Start the recursive point generation from the highest dimension
	generatePoints([]int{}, len(dimensions)-1, 0)
}

// Move a point by a vector while making sure it stays within the bounds of the map
func (m *Map[T]) Move(p Point, v Vector) Point {
	newPoint, err := Move(p, v)
	if err != nil {
		fmt.Println("error. point=", Point{})
		return Point{}
	}

	return newPoint
}

// CHeck if a point is inside the map.
func (m *Map[T]) IsInsideMap(point Point) bool {
	// Check if the new point is within the bounds of the map
	for i, coord := range point.Coordinates {
		if coord < 0 || coord >= m.Dimensions[i] {
			return false
		}
	}
	return true
}

// Find the value / sprite at a given point on the map
func (m *Map[T]) ValueAt(location Point) (T, error) {
	if _, ok := m.Positions[location.String()]; !ok {
		var zero T
		return zero, fmt.Errorf("Location %v is not in the map", location)
	}
	return m.Positions[location.String()], nil
}
