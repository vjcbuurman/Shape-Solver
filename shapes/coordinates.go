package shapes

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type Coordinates map[Coordinate]bool

func (c Coordinates) AsSlice() []Coordinate {
	// returns the Coordinates object as a slice of Coordinate objects
	slice := make([]Coordinate, len(c))
	i := 0
	for element := range c {
		slice[i] = element
		i++
	}
	sortCoordinateSlice(slice)
	return slice
}

func sortCoordinateSlice(coordinates []Coordinate) {
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i].X < coordinates[j].X {
			// if X is less than always less
			return true
		} else if coordinates[i].X == coordinates[j].X && coordinates[i].Y < coordinates[j].Y {
			// if X equal, then Y MUST be less
			return true
		} else {
			return false
		}
	})
}

func (c Coordinates) Identifier() string {
	// Returns an identifier based on a sorted, normalized version of the Coordinates
	// get the coordinates as a simple array of Coordinate objects
	slice := c.Normalize().AsSlice()
	sortCoordinateSlice(slice)
	// simply return as a string
	return fmt.Sprint(slice)
}

func NewCoordinates(coordinateSlice []Coordinate) Coordinates {
	// This is the intended way to create a Coordinates object
	// Creates a Coordinates object from a slice of Coordinate objects
	coordinates := make(Coordinates)
	for _, coordinate := range coordinateSlice {
		coordinates[coordinate] = true
	}
	return coordinates
}

func CoordinatesAreEqual(a, b Coordinates) bool {
	return reflect.DeepEqual(a, b)
}

func (c Coordinates) ContainsCoordinate(coordinate Coordinate) bool {
	// Returns whether the shape contains coordinate c
	return c[coordinate]
}

func (c Coordinates) ContainsCoordinates(o Coordinates) bool {
	for coordinates, _ := range o {
		if !c.ContainsCoordinate(coordinates) {
			return false
		}
	}
	return true
}

func (c Coordinates) Anchor() (Coordinates, Coordinate) {
	// Creates a new Coordinates object that is equivalent in figure, but offset so that:
	// X_min = 0 and Y_min = 0
	// first calculate, over all coordinates the minimum values for X and Y
	minX := math.MaxInt
	minY := math.MaxInt
	for coordinates := range c {
		if coordinates.X < minX {
			minX = coordinates.X
		}
		if coordinates.Y < minY {
			minY = coordinates.Y
		}
	}
	// calculate new anchored coordinates
	anchoredCoordinates := make(Coordinates)
	for coordinates := range c {
		anchoredCoordinates[Coordinate{coordinates.X - minX, coordinates.Y - minY}] = true
	}
	// Return the anchored coordinates and the offset that was applied
	return anchoredCoordinates, Coordinate{minX, minY}
}

func (c Coordinates) Subtract(o Coordinates) (Coordinates, bool) {
	// Returns a new Coordinates object that is the result of taking c and
	// 'removing' the elements in o
	// returns uninitialized shape and false if o cannot be subtracted from s
	if c.ContainsCoordinates(o) {
		rest := make(Coordinates)
		for coordinate := range c {
			if !o.ContainsCoordinate(coordinate) {
				rest[coordinate] = true
			}
		}
		return rest, true
	} else {
		return Coordinates{}, false
	}
}

func (c Coordinates) Transpose() Coordinates {
	// Transposes the Coordinates (X <-> Y)
	transposed := make(Coordinates)
	for coordinates := range c {
		transposed[Coordinate{coordinates.Y, coordinates.X}] = true
	}
	return transposed
}

func (c Coordinates) ReflectX() Coordinates {
	reflected := make(Coordinates)
	for coordinates := range c {
		reflected[Coordinate{-coordinates.X, coordinates.Y}] = true
	}
	return reflected
}

func (c Coordinates) ReflectY() Coordinates {
	reflected := make(Coordinates)
	for coordinates := range c {
		reflected[Coordinate{coordinates.X, -coordinates.Y}] = true
	}
	return reflected
}

func (c Coordinates) Normalize() Coordinates {
	coordinates := c
	// first make sure X is always the 'longest' side
	if !coordinates.longSideIsX() {
		coordinates = coordinates.Transpose()
	}
	// make sure shape is left heavy
	if !coordinates.leftHeavy() {
		coordinates = coordinates.ReflectX()
	}
	// make sure shape is top heavy
	if !coordinates.topHeavy() {
		coordinates = coordinates.ReflectY()
	}
	coordinates, _ = coordinates.Anchor()
	return coordinates
}

func (c Coordinates) intervalXY() ([2]int, [2]int) {
	// returns 2 [int, int] arrays with minimum and maximum values for X and Y
	minX := math.MaxInt
	minY := math.MaxInt
	maxX := -math.MaxInt
	maxY := -math.MaxInt
	// calculate the min and max value for each coordinate in the set
	for coordinate := range c {
		if coordinate.X < minX {
			minX = coordinate.X
		}
		if coordinate.X > maxX {
			maxX = coordinate.X
		}
		if coordinate.Y < minY {
			minY = coordinate.Y
		}
		if coordinate.Y > maxY {
			maxY = coordinate.Y
		}
	}
	return [2]int{minX, maxX}, [2]int{minY, maxY}
}

func (c Coordinates) rangeXY() (int, int) {
	rangeX, rangeY := c.intervalXY()
	return rangeX[1] - rangeX[0], rangeY[1] - rangeY[0]
}

func (c Coordinates) longSideIsX() bool {
	// Also returns true if both are of equal length
	rangeX, rangeY := c.rangeXY()
	return rangeX >= rangeY
}

func (c Coordinates) leftHeavy() bool {
	// get the 'middle' of the Coordinates
	intervalX, _ := c.intervalXY()
	middleX := float64(intervalX[1]-intervalX[0]) / 2.0
	sumX := 0
	for coordinate := range c {
		sumX += coordinate.X
	}
	// average over X
	meanX := float64(sumX) / float64(len(c))
	// check if average is smaller or larger than middle value
	return meanX <= middleX
}

func (c Coordinates) topHeavy() bool {
	// get the 'middle' of the Coordinates
	_, intervalY := c.intervalXY()
	middleY := float64(intervalY[1]-intervalY[0]) / 2.0
	// average over Y
	sumY := 0
	for coordinate := range c {
		sumY += coordinate.Y
	}
	meanY := float64(sumY) / float64(len(c))
	// check if average is smaller or larger than middle value
	return meanY <= middleY
}
