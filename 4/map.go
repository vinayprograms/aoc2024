package main

import "fmt"

type Point struct {
	coordinates []int
}

// Helper function to convert Point to a string key for use in `Map`
func (p Point) StringKey() string {
	return fmt.Sprint(p.coordinates) // Example: "[1 2 3]"
}

type Vector struct {
	deltas []int
}

func (a Vector) Add(b Vector) (Vector, error) {
	if len(a.deltas) != len(b.deltas) {
		return Vector{}, fmt.Errorf("Dimension mismatch: a=%d vs. b=%d", len(a.deltas), len(b.deltas))
	}

	newVector := Vector{deltas: make([]int, len(a.deltas), len(a.deltas))}
	for i := 0; i < len(a.deltas); i++ {
		newVector.deltas[i] = a.deltas[i] + b.deltas[i]
	}
	return newVector, nil
}

func Move(p Point, v Vector) (Point, error) {
	if len(p.coordinates) != len(v.deltas) {
		return Point{}, fmt.Errorf("Dimension mismatch: p=%d vs. v=%d", len(p.coordinates), len(v.deltas))
	}

	newPoint := Point{coordinates: make([]int, len(p.coordinates), len(p.coordinates))}
	for i := 0; i < len(p.coordinates); i++ {
		newPoint.coordinates[i] = p.coordinates[i] + v.deltas[i]
	}
	return newPoint, nil
}

type Map[T any] struct {
	positions  map[string]T
	dimensions []int
}

func (m *Map[T]) Build(dimensions []int, values []T) {
	// Calculate the total number of positions
	totalEntries := 1
	for _, dim := range dimensions {
		totalEntries *= dim
	}
	// Ensure the values length matches the required number of positions
	if len(values) != totalEntries {
		panic(fmt.Sprintf("Number of values (%d) does not match the total positions (%d)", len(values), totalEntries))
	}

	m.dimensions = dimensions
	m.positions = make(map[string]T)

	// Helper function to recursively generate points
	var generatePoints func([]int, int, int)
	generatePoints = func(currentPoint []int, dimensionIndex int, valueIndex int) {
		if dimensionIndex < 0 {
			// Base case: Add Point to the map
			point := Point{coordinates: append([]int{}, currentPoint...)} // Copy the slice
			m.positions[point.StringKey()] = values[valueIndex]
			return
		}

		// Iterate over the range of the current dimension
		for i := 0; i < dimensions[dimensionIndex]; i++ {
			generatePoints(append([]int{i}, currentPoint...), dimensionIndex-1, valueIndex*dimensions[dimensionIndex]+i)
		}
	}

	// Start the recursive point generation from the rightmost dimension
	generatePoints([]int{}, len(dimensions)-1, 0)
}

// Move a point by a vector while making sure it stays within the bounds of the map
func (m *Map[T]) Move(p Point, v Vector) (Point, error) {
	newPoint, err := Move(p, v)
	if err != nil {
		return Point{}, err
	}

	// Check if the new point is within the bounds of the map
	for i, coord := range newPoint.coordinates {
		if coord < 0 || coord >= m.dimensions[i] {
			return Point{}, fmt.Errorf("Point %v is out of bounds for map dimensions %v", newPoint, m.dimensions)
		}
	}

	return newPoint, nil
}

func (m *Map[T]) ValueAt(location Point) T {
	if _, ok := m.positions[location.StringKey()]; !ok {
		panic(fmt.Sprintf("Location %v is not in the map", location))
	}
	return m.positions[location.StringKey()]
}
