package main

import "testing"

func TestLine_Orientation(t *testing.T) {
	line := Line{
		Start: Vector2{1, 1},
		End:   Vector2{3, 1},
	}

	if line.Orientation() != Horizontal {
		t.Error("Line should be horizontal")
	}

	line = Line{
		Start: Vector2{1, 3},
		End:   Vector2{1, 1},
	}

	if line.Orientation() != Vertical {
		t.Error("Line should be vertical")
	}
}

func TestLine_Intersects(t *testing.T) {
	l1 := Line{
		Start: Vector2{1, 1},
		End:   Vector2{3, 1},
	}

	l2 := Line{
		Start: Vector2{5, 5},
		End:   Vector2{6, 5},
	}

	if l1.Intersects(l2) {
		t.Error("Lines should not intersect")
	}

	l1 = Line{
		Start: Vector2{1, 1},
		End:   Vector2{3, 1},
	}

	l2 = Line{
		Start: Vector2{2, 2},
		End:   Vector2{2, 0},
	}

	if !l1.Intersects(l2) {
		t.Error("Lines should  intersect")
	}

	l1 = Line{
		Start: Vector2{1, 1},
		End:   Vector2{3, 1},
	}

	l2 = Line{
		Start: Vector2{2, 0},
		End:   Vector2{2, 2},
	}

	if !l1.Intersects(l2) {
		t.Error("Lines should  intersect")
	}
}
