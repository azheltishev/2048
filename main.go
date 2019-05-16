package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
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

	// x, y := rand.Intn(4), rand.Intn(4)
	// val := rand.Intn(2) + 1

	// field[x][y] = val

	for {
		field = randomlySpawn(field)
		drawField(field, cellSize)
		t.Flush()
		e := t.PollEvent()
		if e.Type == t.EventKey {

			if e.Ch == 'q' {
				break
			}

			switch e.Key {
			case t.KeyArrowUp:
				field = shiftRowsUp(field)
			case t.KeyArrowDown:
				field = shiftRowsDown(field)
			case t.KeyArrowLeft:
				field = shiftColsLeft(field)
			case t.KeyArrowRight:
				field = shiftColsRight(field)
			}
		}
	}

	log.Println("good bye!")
}

func randomlySpawn(field [4][4]int) [4][4]int {
	freePlaceAvailable := false

out:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if field[i][j] == 0 {
				freePlaceAvailable = true
				break out
			}
		}
	}

	if !freePlaceAvailable {
		return field
	}
	x, y := rand.Intn(4), rand.Intn(4)

	for field[x][y] != 0 {
		x, y = rand.Intn(4), rand.Intn(4)
	}

	field[x][y] = rand.Intn(2) + 1
	return field
}

func shiftRowsDown(field [4][4]int) [4][4]int {
	for times := 0; times < 4; times++ {
		for y := 3; y > 0; y-- {
			field = shiftRow(y, -1, field)
		}
	}
	return mod(field)
}

func shiftRowsUp(field [4][4]int) [4][4]int {
	for times := 0; times < 4; times++ {
		for y := 0; y < 3; y++ {
			field = shiftRow(y, 1, field)
		}
	}
	return mod(field)
}

func shiftRow(y int, delta int, field [4][4]int) [4][4]int {
	for x := 0; x < 4; x++ {
		if field[x][y+delta] != 0 {
			if field[x][y+delta] == field[x][y] {
				field[x][y] = -(field[x][y] + 1)
				field[x][y+delta] = 0
			} else if field[x][y] == 0 {
				field[x][y] = field[x][y+delta]
				field[x][y+delta] = 0
			}
		}
	}
	return field
}

func mod(field [4][4]int) [4][4]int {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if field[i][j] < 0 {
				field[i][j] = -field[i][j]
			}
		}
	}
	return field
}

func shiftColsRight(field [4][4]int) [4][4]int {
	for times := 0; times < 4; times++ {
		for x := 3; x > 0; x-- {
			field = shiftCols(x, -1, field)
		}
	}
	return mod(field)
}

func shiftColsLeft(field [4][4]int) [4][4]int {
	for times := 0; times < 4; times++ {
		for x := 0; x < 3; x++ {
			field = shiftCols(x, 1, field)
		}
	}
	return mod(field)
}

func shiftCols(x int, delta int, field [4][4]int) [4][4]int {
	for y := 0; y < 4; y++ {
		if field[x+delta][y] != 0 {
			if field[x+delta][y] == field[x][y] {
				field[x][y] = -(field[x][y] + 1)
				field[x+delta][y] = 0
			} else if field[x][y] == 0 {
				field[x][y] = field[x+delta][y]
				field[x+delta][y] = 0
			}
		}
	}
	return field
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
	printText(x, y, cellSize)
}

func printText(x, y int, cellSize int) {
	s := strconv.Itoa((2 << (uint(field[x][y] - 1))))

	for i, c := range s {
		t.SetCell((x*cellSize)+i, y*cellSize, c, t.ColorBlack, t.ColorWhite)
	}
}
