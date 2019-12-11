package main

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2) Normalize() Vector2{
	length := v.Magnitude()
	return Vector2{
		X: v.X / length,
		Y: v.Y / length,
	}
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector2) EqualEpsilon(v2 Vector2, epsilon float64) bool {
	return math.Abs(v.X - v2.X) <= epsilon && math.Abs(v.Y - v2.Y) <= epsilon
}
