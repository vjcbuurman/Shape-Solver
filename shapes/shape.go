package shapes

type Shape struct {
	definition       Coordinates
	definitionOffset Coordinate
	variations       []Coordinates
	normalized       *Coordinates
}

func NewShape(sc Coordinates) (shape Shape) {
	// This is the intended way to construct an object of type Shape
	shape.definition, shape.definitionOffset = sc.Anchor()
	// perform all rotations/reflections etc
	shape.variations = make([]Coordinates, 8)
	shape.variations[0], _ = shape.definition.Anchor()
	shape.variations[1], _ = shape.definition.ReflectX().Anchor()
	shape.variations[2], _ = shape.definition.ReflectY().Anchor()
	shape.variations[3], _ = shape.definition.ReflectY().ReflectX().Anchor()
	shape.variations[4], _ = shape.definition.Transpose().Anchor()
	shape.variations[5], _ = shape.definition.Transpose().ReflectX().Anchor()
	shape.variations[6], _ = shape.definition.Transpose().ReflectY().Anchor()
	shape.variations[7], _ = shape.definition.Transpose().ReflectX().ReflectY().Anchor()
	// TODO select only unique variations
	// TODO select the left / top heavy shape
	shape.normalized = &(shape.variations[0])

	return shape
}

func (s Shape) Definition() []Coordinate {
	// Returns the definition (the coordinates) that define the Shape
	return s.definition.AsSlice()
}

func (s Shape) Normalized() []Coordinate {
	// Returns the definition (the coordinates) that define the Shape
	return s.normalized.AsSlice()
}

//
//func (s Shape) Normalized() []Coordinate {
//	// TODO should return Coordinates or []Coordinate?
//	// Returns the definition (the coordinates) that define the Shape
//	return extractKeys(s.normalized)
//}
//
//
//func (s Shape) Subtract(o Shape) (Shape, bool) {
//	// Returns a new shape that is the result of taking s and
//	// 'removing' the elements in o
//	// returns uninitialized shape and false if o cannot be subtracted from s
//	if s.ContainsShape(o) {
//		subtracted := make([]Coordinate, 0)
//		for coordinate := range s.definition {
//			if !o.ContainsCoordinate(coordinate) {
//				subtracted = append(subtracted, s_coord)
//			}
//		}
//		return NewShape(subtracted), true
//	} else {
//		return Shape{}, false
//	}
//}
//
//

func ShapesAreEqual(a, b Shape) bool {
	// Returns true if the shapes have the exact same definition (up to a possible offset)
	return CoordinatesAreEqual(a.definition, b.definition)
}

func ShapesAreEquivalent(a, b Shape) bool {
	// Returns true if the shapes are equivalent, meaning they are equal up to rotations and reflections
	return CoordinatesAreEqual(*(a.normalized), *(b.normalized))
}
