package main

import (
	"log"
	"math"
	"os"

	t "github.com/nsf/termbox-go"
)

var field [4][4]int

func main() {
	if err := t.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	w, h := t.Size()

	cellSize := int(math.Min(float64(h), float64(w))-2) / 4

	printField(cellSize)
	t.Flush()
}

func printField(cellSize int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			printSquare(cellSize*i, cellSize*j, 5, 5)
		}
	}
}

func printSquare(x, y, w, h int) {
	for i := x; i < x+w; i++ {
		t.SetCell(i, y, '+', t.ColorBlack, t.ColorWhite)
		t.SetCell(i, y+h, '+', t.ColorBlack, t.ColorWhite)
	}
	for j := y; j < y+h; j++ {
		t.SetCell(x, j, '+', t.ColorBlack, t.ColorWhite)
		t.SetCell(x+w, j, '+', t.ColorBlack, t.ColorWhite)
	}
}
