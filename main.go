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

			case t.KeyArrowDown:
				field = shiftRowsDown(field)
			case t.KeyArrowLeft:

			case t.KeyArrowRight:

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
			for x := 0; x < 4; x++ {
				if field[x][y] == 0 {
					if field[x][y-1] != 0 {
						if field[x][y-1] == field[x][y] {
							field[x][y] = -(field[x][y] + 1)
							field[x][y-1] = 0
						} else {
							field[x][y] = field[x][y-1]
							field[x][y-1] = 0
						}
					}
				}
			}
		}
	}
	return mod(field)
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

func shiftCols(yd int, field [4][4]int) [4][4]int {
	// for row := 0; row < 4; row++ {

	// }
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
