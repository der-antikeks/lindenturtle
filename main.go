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
	segmentlength := 4.0

	green := color.RGBA{0x60, 0xFF, 0x00, 0xFF}

	path := lindenmayer([]string{"X"}, map[string][]string{
		"X": {"F", "-", "[", "[", "X", "]", "+", "X", "]", "+", "F", "[", "+", "F", "X", "]", "-", "X"},
		"F": {"F", "F"},
	}, 7)

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

	img := turtle.Go(path)

	saveToPngFile("images/plant.png", img)
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

	fmt.Printf("Wrote %s OK.\n", filePath)
}
