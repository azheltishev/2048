package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	t "github.com/nsf/termbox-go"
)

var field [4][4]int

func main() {
	if err := t.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())

	cellSize := 4

	x, y := rand.Intn(4), rand.Intn(4)
	val := rand.Intn(2) + 1

	field[x][y] = val

	drawField(field, cellSize)
	t.Flush()
}

func drawField(field [4][4]int, cellSize int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			drawCell(i, j, determineColor(field[i][j]), cellSize)
		}
	}
}

func determineColor(x int) t.Attribute {
	switch x {
	case 0:
		return t.ColorWhite
	case 1:
		return t.ColorYellow
	case 2:
		return t.ColorCyan
	case 3:
		return t.ColorBlue
	case 4:
		return t.ColorGreen
	case 5:
		return t.ColorRed
	case 6:
		return t.ColorRed
	default:
		return t.ColorBlack
	}
}

func drawCell(x, y int, color t.Attribute, cellSize int) {
	for i := 0; i < cellSize; i++ {
		for j := 0; j < cellSize; j++ {
			t.SetCell((x*cellSize)+i, (y*cellSize)+j, ' ', t.ColorDefault, color)
		}
	}
}
