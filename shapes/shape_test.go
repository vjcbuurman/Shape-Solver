package shapes

//
//import (
//	"testing"
//)
//
//func TestCompareShapes(t *testing.T) {
//
//	shape1 := NewShape([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}})
//	shape2 := NewShape([]Coordinate{{X: 1, Y: 1}, {X: 0, Y: 0}})
//
//	if !IsEqual(shape1, shape2) {
//		t.Fatal()
//	}
//
//}
//
//func TestContainsShape(t *testing.T) {
//	smaller_coords := []Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}}
//	larger_coords := []Coordinate{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}}
//	smaller_shape := NewShape(smaller_coords)
//	larger_shape := NewShape(larger_coords)
//	if !larger_shape.ContainsShape(smaller_shape) {
//		t.Error()
//	}
//}
//
//func TestReflect(t *testing.T) {
//	shape := NewShape([]Coordinate{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}})
//	reflected_shape := NewShape([]Coordinate{{X: 0, Y: 0}, {X: -1, Y: 0}, {X: -1, Y: 1}})
//	if !IsEqual(shape.Reflect(), reflected_shape) {
//		t.Error()
//	}
//}
//

//func TestSubtract(t *testing.T) {
//	var tests = []struct {
//		input      [2]Shape
//		want_shape Shape
//		want_ok    bool
//	}{
//		{[2]Shape{NewShape([]Coordinate{{X: 1, Y: 2}, {X: 2, Y: 2}}), NewShape([]Coordinate{{X: 1, Y: 2}})}, NewShape([]Coordinate{{X: 2, Y: 2}}), true},
//		{[2]Shape{NewShape([]Coordinate{{X: 0, Y: 0}}), NewShape([]Coordinate{{X: 1, Y: 1}})}, Shape{}, false},
//	}
//
//	for _, tt := range tests {
//		subtracted, ok := tt.input[0].Subtract(tt.input[1])
//		if !IsEqual(subtracted, tt.want_shape) {
//			t.Errorf("Subtracted shape not equal to expected: %v and %v", tt.want_shape, subtracted)
//		}
//		if tt.want_ok != ok {
//			t.Errorf("ok not equal to expected: %v and %v", tt.want_ok, ok)
//		}
//	}
//
//}
