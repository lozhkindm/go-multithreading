package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v Vector2D) Add(vector Vector2D) Vector2D {
	return Vector2D{
		x: v.x + vector.x,
		y: v.y + vector.y,
	}
}

func (v Vector2D) Sub(vector Vector2D) Vector2D {
	return Vector2D{
		x: v.x - vector.x,
		y: v.y - vector.y,
	}
}

func (v Vector2D) Multi(vector Vector2D) Vector2D {
	return Vector2D{
		x: v.x * vector.x,
		y: v.y * vector.y,
	}
}

func (v Vector2D) AddVal(val float64) Vector2D {
	return Vector2D{
		x: v.x + val,
		y: v.y + val,
	}
}

func (v Vector2D) MultiVal(val float64) Vector2D {
	return Vector2D{
		x: v.x * val,
		y: v.y * val,
	}
}

func (v Vector2D) DivVal(val float64) Vector2D {
	return Vector2D{
		x: v.x / val,
		y: v.y / val,
	}
}

func (v Vector2D) limit(low, up float64) Vector2D {
	return Vector2D{
		x: math.Min(math.Max(v.x, low), up),
		y: math.Min(math.Max(v.y, low), up),
	}
}

func (v Vector2D) Distance(vector Vector2D) float64 {
	return math.Sqrt(math.Pow(v.x-vector.x, 2) + math.Pow(v.y-vector.y, 2))
}
