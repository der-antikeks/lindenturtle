package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func main() {
	saveToPngFile("images/bintree.png", bintree(8))
	saveToPngFile("images/plant.png", plant(7))
	saveToPngFile("images/tree.png", tree(9))
}

func tree(it int) image.Image {
	segmentlength := 10.0

	green := color.RGBA{0x60, 0xFF, 0x00, 0xFF}
	black := color.RGBA{0x33, 0x33, 0x33, 0xFF}

	path := lindenmayer([]string{"leaf"}, map[string][]string{
		"leaf":   {"branch", "[", "<", "leaf", "]", "<>", "branch", "[", ">", "leaf", "]"},
		"branch": {"trunk", "[", "<", "leaf", "]", "[", "<>", "branch", "]", "[", ">", "leaf", "]"},
		"trunk":  {"+", "trunk", "<>", "trunk", "-"},
		"+":      {"+", "+"},
		"-":      {"-", "-"},
	}, it)

	rmm := func(min, max float64) float64 {
		return rand.Float64()*(max-min) + min
	}

	turtle := NewTurtle(map[string]func(*Turtle){
		"leaf": func(t *Turtle) { // // first gen, green leave
			t.SetColor(green)
			w := t.Width()
			t.SetWidth(10 * rand.Float64())
			t.Draw(segmentlength/2*rand.Float64(), 0)
			t.SetWidth(w)
		},
		"branch": func(t *Turtle) { // // second gen, black branch
			t.SetColor(black)
			//t.SetWidth(1)
			t.Draw(segmentlength*rand.Float64(), 0)
		},
		"trunk": func(t *Turtle) { // // second gen, black branch
			t.SetColor(black)
			//t.SetWidth(2)
			t.Draw(segmentlength*rand.Float64(), 0)
		},
		"+": func(t *Turtle) { // thicken
			t.SetWidth(t.Width() + 0.25)
		},
		"-": func(t *Turtle) { // thinning
			// w := t.Width() - 1
			// if w < 1 { w = 1 }
			t.SetWidth(t.Width() - 0.25)
		},
		"<": func(t *Turtle) { // turn left
			t.Turn(rmm(-45, 0))
		},
		">": func(t *Turtle) { // turn right
			t.Turn(rmm(0, 45))
		},
		"<>": func(t *Turtle) { // wiggle
			t.Turn(rmm(-2, 2))
		},
		"[": func(t *Turtle) { // push position and angle
			t.Save()
		},
		"]": func(t *Turtle) { // pop position and angle
			t.Restore()
		},
	})

	turtle.SetColor(black)
	turtle.SetWidth(1)
	turtle.Turn(-90)

	return turtle.Go(path)
}

func bintree(it int) image.Image {
	segmentlenght := 1.0

	green := color.RGBA{0x33, 0xFF, 0x33, 0xFF}
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF}

	path := lindenmayer([]string{"0"}, map[string][]string{
		"1": {"1", "1"},
		"0": {"1", "[", "0", "]", "0"},
	}, it)

	turtle := NewTurtle(map[string]func(*Turtle){
		"0": func(t *Turtle) {
			// draw a line segment ending in a leaf
			t.SetColor(green)
			t.SetWidth(5)
			t.Draw(segmentlenght, 0)
		},
		"1": func(t *Turtle) {
			// draw a line segment
			t.SetColor(black)
			t.SetWidth(1)
			t.Draw(segmentlenght, 0)
		},
		"[": func(t *Turtle) {
			// push position and angle, turn left 45 degrees
			t.Save()
			t.Turn(-45)
		},
		"]": func(t *Turtle) {
			// pop position and angle, turn right 45 degrees
			t.Restore()
			t.Turn(45)
		},
	})

	turtle.Turn(-90)
	return turtle.Go(path)
}

func plant(it int) image.Image {
	segmentlength := 4.0

	green := color.RGBA{0x60, 0xFF, 0x00, 0xFF}

	path := lindenmayer([]string{"X"}, map[string][]string{
		"X": {"F", "-", "[", "[", "X", "]", "+", "X", "]", "+", "F", "[", "+", "F", "X", "]", "-", "X"},
		"F": {"F", "F"},
	}, it)

	turtle := NewTurtle(map[string]func(*Turtle){
		"F": func(t *Turtle) {
			// draw forward
			t.Draw(segmentlength*rand.Float64(), 0)
		},
		"-": func(t *Turtle) {
			// turn left 25°
			t.Turn(-20 + rand.Float64()*10)
		},
		"+": func(t *Turtle) {
			// turn right 25°
			t.Turn(20 + rand.Float64()*10)
		},
		"[": func(t *Turtle) {
			// push position and angle
			t.Save()
		},
		"]": func(t *Turtle) {
			// pop position and angle
			t.Restore()
		},
	})

	turtle.SetColor(green)
	turtle.SetWidth(1)
	turtle.Turn(-90)

	return turtle.Go(path)
}

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %s\n", filePath)
}
