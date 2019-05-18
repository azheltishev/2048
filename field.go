package _2048

import (
	"errors"
	"math/rand"
	"time"
)

type Field struct {
	Width, Height int
	Tiles         [][]uint64
	blocked       [][]bool
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (f *Field) Init(x, y int) {
	f.Width = x
	f.Height = y
	f.Tiles = make([][]uint64, x, x)
	f.blocked = make([][]bool, x, x)
	for i := 0; i < x; i++ {
		f.Tiles[i] = make([]uint64, y, y)
		f.blocked[i] = make([]bool, y, y)
	}
}

func (f *Field) isFull() bool {
	for x := 0; x < f.Width; x++ {
		for y := 0; y < f.Height; y++ {
			if f.Tiles[x][y] == 0 {
				return false
			}
		}
	}

	return true
}

func (f *Field) SpawnTile() error {

	if f.isFull() {
		return errors.New("field is full - cannot spawn")
	}

	x, y := rand.Intn(f.Width), rand.Intn(f.Height)

	for f.Tiles[x][y] != 0 {
		x, y = rand.Intn(4), rand.Intn(4)
	}

	f.Tiles[x][y] = 2 << uint64(rand.Intn(2))

	return nil
}

func (f *Field) ShiftDown() {
	for times := 0; times < f.Height; times++ {
		for y := f.Height - 1; y > 0; y-- {
			f.shiftRows(y, -1)
		}
	}
	f.unblock()
}

func (f *Field) ShiftUp() {
	for times := 0; times < f.Height; times++ {
		for y := 0; y < f.Height-1; y++ {
			f.shiftRows(y, 1)
		}
	}
	f.unblock()
}

func (f *Field) shiftRows(y int, delta int) {
	for x := 0; x < f.Width; x++ {
		if f.Tiles[x][y+delta] != 0 {
			if f.Tiles[x][y+delta] == f.Tiles[x][y] && !f.blocked[x][y+delta] && !f.blocked[x][y] {
				f.Tiles[x][y] = f.Tiles[x][y] * 2
				f.blocked[x][y] = true
				f.Tiles[x][y+delta] = 0
			} else if f.Tiles[x][y] == 0 {
				f.Tiles[x][y] = f.Tiles[x][y+delta]
				f.Tiles[x][y+delta] = 0
			}
		}
	}
}

func (f *Field) ShiftRight() {
	for times := 0; times < f.Width; times++ {
		for x := f.Width - 1; x > 0; x-- {
			f.shiftColumns(x, -1)
		}
	}
	f.unblock()
}

func (f *Field) ShiftLeft() {
	for times := 0; times < f.Width; times++ {
		for x := 0; x < f.Width-1; x++ {
			f.shiftColumns(x, 1)
		}
	}
	f.unblock()
}

func (f *Field) shiftColumns(x int, delta int) {
	for y := 0; y < f.Height; y++ {
		if f.Tiles[x+delta][y] != 0 {
			if f.Tiles[x+delta][y] == f.Tiles[x][y] && !f.blocked[x+delta][y] && !f.blocked[x][y] {
				f.Tiles[x][y] = f.Tiles[x][y] * 2
				f.blocked[x][y] = true
				f.Tiles[x+delta][y] = 0
			} else if f.Tiles[x][y] == 0 {
				f.Tiles[x][y] = f.Tiles[x+delta][y]
				f.Tiles[x+delta][y] = 0
			}
		}
	}
}

func (f *Field) unblock() {
	for x := 0; x < f.Width; x++ {
		for y := 0; y < f.Height; y++ {
			f.blocked[x][y] = false
		}
	}

}
