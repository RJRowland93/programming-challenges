package main

import (
	"fmt"
	"math"
	"strings"
)

// type Error struct {
// 	msg string
// }

type Vec4 struct {
	x float64
	y float64
	z float64
	w float64
}

/*
What if Vec4 was implemented as array?
Easier to iterate through and add multiple vectors at once?
*/

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

func (t Vec4) Dot(o Vec4) float64 {

	return t.x*o.x +
		t.y*o.y +
		t.z*o.z +
		t.w*o.w
}

func (t Vec4) Cross(o Vec4) Vec4 {

	return Vec4{
		x: t.y*o.z - t.z*o.y,
		y: t.z*o.x - t.x*o.z,
		z: t.x*o.y - t.y*o.x,
		w: 0,
	}
}

type projectile struct {
	pos Vec4
	vel Vec4
}

type environment struct {
	grav Vec4
	wind Vec4
}

func tick(env environment, proj projectile) projectile {
	return projectile{
		pos: proj.pos.Add(proj.vel),
		vel: proj.vel.Add(env.grav.Add(env.wind)),
	}
}

/*
Enums seem lacking in go.
Actually wanted Tagged Sum types from Rust
Can use bit flags for enums via iota and bit shiting
*/

/*
Generics can be done via interfaces
*/

type ColorF64 struct {
	red   float64
	green float64
	blue  float64
}

func (c ColorF64) Add(o ColorF64) ColorF64 {
	return ColorF64{
		red:   c.red + o.red,
		green: c.green + o.green,
		blue:  c.blue + o.blue,
	}
}

func (c ColorF64) Sub(o ColorF64) ColorF64 {
	return ColorF64{
		red:   c.red - o.red,
		green: c.green - o.green,
		blue:  c.blue - o.blue,
	}
}

func (c ColorF64) Hadamard(o ColorF64) ColorF64 {
	return ColorF64{
		red:   c.red * o.red,
		green: c.green * o.green,
		blue:  c.blue * o.blue,
	}
}

func (c ColorF64) Scale(x float64) ColorF64 {
	return ColorF64{
		red:   c.red * x,
		green: c.green * x,
		blue:  c.blue * x,
	}
}

/*
Floats and Ints are literal values; Not sure if they share an interface
Was trying to not reimplement Color with f64 and uint8...
However rethough original design and made it simpler, which was probably the point?
Was doing a combination of floor + max
Funtions need to be bigger since Go doesn't give you a lot
*/

func clamp(x float64) uint8 {
	if x < 0 {
		return 0
	}
	if x > 255 {
		return 255
	}
	return uint8(x)
}

type Canvas struct {
	height uint
	width  uint
	grid   [][]ColorF64
}

func NewCanvas(height, width uint) Canvas {
	grid := make([][]ColorF64, height)

	for h := range grid {
		grid[h] = make([]ColorF64, width)
	}

	return Canvas{
		height: height,
		width:  width,
		grid:   grid,
	}
}

func (c Canvas) writePixel(x, y uint64, p ColorF64) {
	c.grid[x][y] = p
}

// func (c Canvas) invertY() {
// 	for _, row := range
// }

func (c Canvas) ToPPM() string {

	var b strings.Builder
	// header
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.height, c.width))
	b.WriteString("255\n")

	// rbg values
	for _, row := range c.grid {
		for _, col := range row {
			// clamp to 70 chars
			b.WriteString(fmt.Sprintf("%d %d %d ", clamp(col.red), clamp(col.green), clamp(col.blue)))
		}
		b.WriteString("\n")
	}
	// end
	b.WriteString("\n")

	return b.String()
}

func (c Canvas) SaveToFile(path string) {

}

type Matrix [4][4]float64

func (m Matrix) MatMult(o Matrix) Matrix {
	result := Matrix{}

	for r, row := range m {
		for c := range row {
			result[r][c] = m[r][c] * o[c][r]
		}
	}

	return result
}

type Test struct {
	arr [4]int
	x int = arr[0]
}

func (m Matrix) VecMult(o Vec4) Vec4 {
	result := Vec4{}

	for r, row := range m {
		n := 0.0
		for c := range row {
			n += m[r][c] * o.x
		}
	}

	return result
}

func cannon() {
	p := projectile{
		pos: NewPoint(0, 1, 0),
		vel: NewVec(1, 1, 0).Norm(),
	}
	env := environment{
		grav: NewVec(0, -0.1, 0),
		wind: NewVec(-0.01, 0, 0),
	}

	for p.pos.y >= 0 {
		fmt.Printf("%v\n", p)
		p = tick(env, p)
	}

}

func ppm() {
	c := NewCanvas(3, 3)
	p1 := ColorF64{1, 2, 3}
	p2 := ColorF64{3, 2, 1}

	c.writePixel(1, 1, p1)
	c.writePixel(0, 0, p2)

	fmt.Printf(c.ToPPM())
	// fmt.Println(70/4 - 1)
}

func main() {
	m := Matrix{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4}}

	o := Matrix{
		{1, 2, 3, 4},
		// {1, 2, 3, 4},
		// {1, 2, 3, 4},
		// {1, 2, 3, 4}}
	}

	fmt.Println(m.Mult(o))
}
