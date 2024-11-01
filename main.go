package main

import (
	"fmt"
	"solver/shapes"
)

func main() {

	// L shape coordinates
	//coordinates := shapes.NewCoordinates([]shapes.Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}})
	//coordinates := shapes.NewCoordinates([]shapes.Coordinate{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 2, Y: 0}})
	coordinates := shapes.NewCoordinates([]shapes.Coordinate{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 2}})
	//shape := shapes.NewShape(coordinates)

	coordinates.AsSlice()

	fmt.Printf(coordinates.Identifier())

	//fmt.Printf("Variations: %v", shape)

}
