package main

import (
	"image"
	"image/color"
	"math"

	"code.google.com/p/draw2d/draw2d"
)

type aabb struct {
	minx, miny, maxx, maxy float64
}

func (a *aabb) addPoint(x, y float64) {
	if a.minx > x {
		a.minx = x
	}
	if a.maxx < x {
		a.maxx = x
	}
	if a.miny > y {
		a.miny = y
	}
	if a.maxy < y {
		a.maxy = y
	}
}

func (a *aabb) size() (width, height float64) {
	return a.maxx - a.minx, a.maxy - a.miny
}

type stack struct {
	x, y, a float64
	color   color.Color
	width   float64
	prev    *stack
}

type line struct {
	x1, y1, x2, y2 float64
	color          color.Color
	width          float64
}

type Turtle struct {
	rules    map[string]func(*Turtle)
	cur      *stack
	lines    []line
	boundary *aabb
}

func NewTurtle(rules map[string]func(*Turtle)) *Turtle {
	return &Turtle{
		rules: rules,
		cur: &stack{
			color: color.RGBA{0x00, 0x00, 0x00, 0xFF},
			width: 1,
		},
		lines:    []line{},
		boundary: &aabb{},
	}
}

// save position and angle
func (t *Turtle) Save() {
	t.cur = &stack{
		x:     t.cur.x,
		y:     t.cur.y,
		a:     t.cur.a,
		color: t.cur.color,
		width: t.cur.width,
		prev:  t.cur, // stackception
	}
}

// restore position and angle
func (t *Turtle) Restore() {
	if t.cur == nil {
		return
	}

	if p := t.cur.prev; p != nil {
		t.cur.prev = nil // stackalypse
		t.cur = p
	}
}

// rotate vector
func vrot(forward, right, angle float64) (x, y float64) {
	cs, sn := math.Cos(angle), math.Sin(angle)
	x, y = forward*cs-right*sn, forward*sn+right*cs
	return
}

// move forward without drawing
func (t *Turtle) Move(f, r float64) {
	x, y := vrot(f, r, t.cur.a)
	t.cur.x += x
	t.cur.y += y

	t.boundary.addPoint(t.cur.x, t.cur.y)
}

// draw line forward
func (t *Turtle) Draw(f, r float64) {
	sx, sy := t.cur.x, t.cur.y

	x, y := vrot(f, r, t.cur.a)
	t.cur.x += x
	t.cur.y += y

	t.boundary.addPoint(t.cur.x, t.cur.y)

	t.lines = append(t.lines, line{
		x1:    sx,
		y1:    sy,
		x2:    t.cur.x,
		y2:    t.cur.y,
		color: t.cur.color,
		width: t.cur.width,
	})
}

// get position
func (t *Turtle) Position() (x, y float64) {
	return t.cur.x, t.cur.y
}

// add rotation
func (t *Turtle) Turn(degree float64) {
	t.cur.a += (degree * (math.Pi / 180.0))
}

// get angle
func (t *Turtle) Angle() (degree float64) {
	return t.cur.a * (180.0 / math.Pi)
}

// set color
func (t *Turtle) SetColor(c color.Color) {
	t.cur.color = c
}

// get color
func (t *Turtle) Color() color.Color {
	return t.cur.color
}

// set width
func (t *Turtle) SetWidth(w float64) {
	t.cur.width = w
}

// get width
func (t *Turtle) Width() float64 {
	return t.cur.width
}

// UNLEASH THE TURTLE!
func (t *Turtle) Go(path []string) image.Image {
	// draw lines
	for _, c := range path {
		if f, ok := t.rules[c]; ok {
			f(t)
		}
	}

	border := 5.0
	w, h := t.boundary.size()
	offx, offy := t.boundary.minx-border, t.boundary.miny-border

	img := image.NewRGBA(image.Rect(0, 0, int(w+border*2), int(h+border*2)))
	gc := draw2d.NewGraphicContext(img)

	for _, line := range t.lines {
		gc.SetLineWidth(line.width)
		gc.SetStrokeColor(line.color)

		gc.MoveTo(line.x1-offx, line.y1-offy)
		gc.LineTo(line.x2-offx, line.y2-offy)
		gc.Stroke()
	}

	// cleanup
	t.cur = &stack{
		color: color.RGBA{0x00, 0x00, 0x00, 0xFF},
		width: 1,
	}
	t.lines = []line{}
	t.boundary = &aabb{}

	return img
}
