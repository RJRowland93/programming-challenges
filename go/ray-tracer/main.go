package main

import (
	"fmt"
	"math"
)

// type Error struct {
// 	msg string
// }

type Vec4 struct {
	x, y, z, w float64
}

func NewPoint(x, y, z float64) Vec4 {
	return Vec4{x: x, y: y, z: z, w: 1}
}

func NewVec(x, y, z float64) Vec4 {
	return Vec4{x: x, y: y, z: z, w: 0}
}

func (t Vec4) Add(o Vec4) Vec4 {
	return Vec4{
		x: t.x + o.x,
		y: t.y + o.y,
		z: t.z + o.z,
		w: t.w + o.w,
	}
}

func (t Vec4) Sub(o Vec4) Vec4 {
	return Vec4{
		x: t.x - o.x,
		y: t.y - o.y,
		z: t.z - o.z,
		w: t.w - o.w,
	}
}

func (t Vec4) Mult(o Vec4) Vec4 {
	return Vec4{
		x: t.x * o.x,
		y: t.y * o.y,
		z: t.z * o.z,
		w: t.w * o.w,
	}
}

func (t Vec4) Div(o Vec4) Vec4 {
	return Vec4{
		x: t.x / o.x,
		y: t.y / o.y,
		z: t.z / o.z,
		w: t.w / o.w,
	}
}

// func (t Vec4) Mag() (float64, error) {
// 	if t.w != 0 {
// 		nil, errors.New("tried to get magnitude of point")
// 	}

// 	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2)), nil
// }

/*
Is throwing a panic the better design choice? Tried to initially return an error, but ran into issues with the result return value
Could return nil value for result struct, but only if it was a pointer
Returning a pointer to a float doesn't make sense since it's a scalar (?)
*/

func (t Vec4) Mag() float64 {
	if t.w != 0 {
		panic("tried to get magnitude of point")
	}

	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2))
}

func (t Vec4) Norm() Vec4 {
	mag := t.Mag()

	return Vec4{
		x: t.x / mag,
		y: t.y / mag,
		z: t.z / mag,
		w: t.w / mag,
	}
}

func main() {
	// x := NewPoint(1, 2, 3)
	// y := NewPoint(3, 2, 1)
	z := NewVec(4, 0, 4).Mag()

	fmt.Printf("%v\n", z)
}
