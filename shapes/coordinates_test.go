package shapes

import (
	"reflect"
	"testing"
)

func TestCoordinatesAreEqual(t *testing.T) {
	a := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}})
	b := a
	if !CoordinatesAreEqual(a, b) {
		t.Errorf("Coordinates %v and %v should be equal", a, b)
	}
	c := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 2, Y: 2}})
	if CoordinatesAreEqual(a, c) {
		t.Errorf("Coordinates %v and %v should not be equal", a, c)
	}
}

func TestContainsCoordinate(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}})
	shouldContain := Coordinate{X: 0, Y: 0}
	if !coordinates.ContainsCoordinate(shouldContain) {
		t.Errorf("Coordinates %v should contain coordinate %v", coordinates, shouldContain)
	}
	shouldNotContain := Coordinate{X: 2, Y: 2}
	if coordinates.ContainsCoordinate(shouldNotContain) {
		t.Errorf("Coordinates %v should not contain coordinate %v", coordinates, shouldNotContain)
	}
}

func TestContainsCoordinates(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}})
	shouldContain := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}})
	if !coordinates.ContainsCoordinates(shouldContain) {
		t.Errorf("Coordinates %v should contain coordinates %v", coordinates, shouldContain)
	}
	shouldNotContain := NewCoordinates([]Coordinate{{X: 0, Y: 1}, {X: 1, Y: 1}})
	if coordinates.ContainsCoordinates(shouldNotContain) {
		t.Errorf("Coordinates %v should not contain coordinate %v", coordinates, shouldNotContain)
	}
}

func TestAsSlice(t *testing.T) {
	coordinateSlice := []Coordinate{{X: 1, Y: 2}, {X: 2, Y: 2}}
	coordinates := NewCoordinates(coordinateSlice)
	asSlice := coordinates.AsSlice()
	if !reflect.DeepEqual(asSlice, coordinateSlice) {
		t.Errorf("AsSlice slice %v not equal to expected: %v", asSlice, coordinateSlice)
	}
}

func TestAnchor(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 1, Y: 2}, {X: 2, Y: 2}})
	wantCoordinates := NewCoordinates([]Coordinate{{0, 0}, {1, 0}})
	wantOffset := Coordinate{1, 2}
	anchored, offset := coordinates.Anchor()
	if !CoordinatesAreEqual(anchored, wantCoordinates) {
		t.Errorf("Ancored shape not equal to expected: %v", anchored)
	}
	if offset != wantOffset {
		t.Errorf("Offset not equal to expected: %v", offset)
	}
}

func TestSubtract(t *testing.T) {
	// check that we subtract correctly
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}})
	toSubtract := NewCoordinates([]Coordinate{{X: 0, Y: 0}})
	rest, ok := coordinates.Subtract(toSubtract)
	wantRest := NewCoordinates([]Coordinate{{X: 1, Y: 1}})
	if !ok || !CoordinatesAreEqual(rest, wantRest) {
		t.Errorf("Coordinates %v and %v should be equal", rest, wantRest)
	}
	// Also check that something that should not be subtractable is correctly returned
	toNotSubtract := NewCoordinates([]Coordinate{{X: 2, Y: 2}})
	actualRest, ok := coordinates.Subtract(toNotSubtract)
	wantNotRest := Coordinates{}
	if ok || !CoordinatesAreEqual(actualRest, wantNotRest) {
		t.Errorf("Coordinates %v and %v should be equal", actualRest, wantNotRest)
	}
}

func TestTranspose(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}})
	transposed := coordinates.Transpose()
	wantCoordinates := NewCoordinates([]Coordinate{{1, 0}, {0, 1}, {1, 1}})
	if !CoordinatesAreEqual(transposed, wantCoordinates) {
		t.Errorf("Coordinates %v and %v should be equal", transposed, wantCoordinates)
	}
}

func TestReflectX(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}})
	reflected := coordinates.ReflectX()
	wantCoordinates := NewCoordinates([]Coordinate{{X: 0, Y: 1}, {X: -1, Y: 0}, {X: -1, Y: 1}})
	if !CoordinatesAreEqual(reflected, wantCoordinates) {
		t.Errorf("Coordinates %v and %v should be equal", reflected, wantCoordinates)
	}
}

func TestReflectY(t *testing.T) {
	coordinates := NewCoordinates([]Coordinate{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}})
	reflected := coordinates.ReflectY()
	wantCoordinates := NewCoordinates([]Coordinate{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: -1}})
	if !CoordinatesAreEqual(reflected, wantCoordinates) {
		t.Errorf("Coordinates %v and %v should be equal", reflected, wantCoordinates)
	}
}

func TestIdentifier(t *testing.T) {
	t.Error()
}

func TestLeftHeavy(t *testing.T) {
	var tests = []struct {
		input Coordinates
		want  bool
	}{
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}}), true},
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 0}}), false},
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}}), true},
	}
	for i, tt := range tests {
		outcome := tt.input.leftHeavy()
		if tt.want != outcome {
			t.Errorf("test %v, leftHeavy not correct for %v, got %v", i+1, tt, outcome)
		}
	}
}

func TestTopHeavy(t *testing.T) {
	var tests = []struct {
		input Coordinates
		want  bool
	}{
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}}), true},
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1}}), false},
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}}), true},
	}
	for i, tt := range tests {
		outcome := tt.input.topHeavy()
		if tt.want != outcome {
			t.Errorf("test %v, topHeavy not correct for %v, got %v", i+1, tt, outcome)
		}
	}
}

func TestRangeXY(t *testing.T) {
	var tests = []struct {
		input Coordinates
		wantX int
		wantY int
	}{
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 2}}), 1, 2},
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}}), 0, 2},
		{NewCoordinates([]Coordinate{{X: 4, Y: -2}, {X: 4, Y: 0}, {X: 7, Y: 2}}), 3, 4},
	}
	for i, tt := range tests {
		rangeX, rangeY := tt.input.rangeXY()
		if tt.wantX != rangeX || tt.wantY != rangeY {
			t.Errorf("test %v, rangeXY not correct for %v, got [%v, %v]", i+1, tt, rangeX, rangeY)
		}
	}
}

func TestIntervalXY(t *testing.T) {
	var tests = []struct {
		input Coordinates
		wantX [2]int
		wantY [2]int
	}{
		{NewCoordinates([]Coordinate{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 2}}), [2]int{0, 1}, [2]int{0, 2}},
		{NewCoordinates([]Coordinate{{X: 4, Y: -2}, {X: 4, Y: 0}, {X: 7, Y: 2}}), [2]int{4, 7}, [2]int{-2, 2}},
	}
	for i, tt := range tests {
		intervalX, intervalY := tt.input.intervalXY()
		if tt.wantX != intervalX || tt.wantY != intervalY {
			t.Errorf("test %v, rangeXY not correct for %v, got [%v, %v]", i+1, tt, intervalX, intervalY)
		}
	}
}

func TestLongSideIsX(t *testing.T) {
	t.Error()
}
