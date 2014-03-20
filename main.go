package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	segmentlenght := 1.0

	green := color.RGBA{0x33, 0xFF, 0x33, 0xFF}
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF}

	path := lindenmayer([]string{"0"}, map[string][]string{
		"1": {"1", "1"},
		"0": {"1", "[", "0", "]", "0"},
	}, 8)

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

	img := turtle.Go(path)

	saveToPngFile("images/tree.png", img)
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
