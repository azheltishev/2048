package main

import (
	"log"
	"os"
	"strconv"

	"github.com/azheltishev/_2048"
	t "github.com/nsf/termbox-go"
)

var wWidth, wHeight int

func main() {
	if err := t.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var field _2048.Field

	cellSize := 4
	fieldSize := 4

	wWidth, wHeight = t.Size()

	field.Init(fieldSize, fieldSize)

	for {
		if err := field.SpawnTile(); err != nil {
			log.Println(err)
			break
		}

		drawField(field.Tiles, fieldSize, cellSize)
		t.Flush()

		e := t.PollEvent()
		if e.Type == t.EventKey {

			if e.Ch == 'q' {
				break
			}

			switch e.Key {
			case t.KeyArrowUp:
				field.ShiftUp()
			case t.KeyArrowDown:
				field.ShiftDown()
			case t.KeyArrowLeft:
				field.ShiftLeft()
			case t.KeyArrowRight:
				field.ShiftRight()
			}
		}
	}

	log.Println("game over")
}

func drawField(tiles [][]uint64, fieldSize int, cellSize int) {
	for i := 0; i < fieldSize; i++ {
		for j := 0; j < fieldSize; j++ {
			drawCell(i, j, determineColor(tiles[i][j]), (wWidth/2)-(fieldSize*cellSize/2), (wHeight/2)-(fieldSize*cellSize/2), cellSize, strconv.Itoa(int(tiles[i][j])))
		}
	}
}

func determineColor(x uint64) t.Attribute {
	switch x {
	case 0:
		return t.ColorWhite
	case 2:
		return t.ColorYellow
	case 4:
		return t.ColorCyan
	case 8:
		return t.ColorBlue
	case 16:
		return t.ColorGreen
	case 32:
		return t.ColorRed
	default:
		return t.ColorBlack
	}
}

func drawCell(x, y int, color t.Attribute, offsetX, offsetY int, cellSize int, s string) {
	for i := 0; i < cellSize; i++ {
		for j := 0; j < cellSize; j++ {
			t.SetCell(offsetX+((x*cellSize)+i), offsetY+(y*cellSize)+j, ' ', t.ColorDefault, color)
		}
	}
	printText(x, y, cellSize, offsetX, offsetY, s)
}

func printText(x, y int, cellSize int, offsetX, offsetY int, s string) {

	for i, c := range s {
		t.SetCell(offsetX+(x*cellSize)+i, offsetY+y*cellSize, c, t.ColorBlack, t.ColorWhite)
	}
}
