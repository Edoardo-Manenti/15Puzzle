package board

import (
    	"math/rand/v2"
	"15puzzle/pkg/coord"
	"errors"
	"fmt"
)

type Board struct {
	width int
	data [][]int
}

func NewBoard(width int) Board {
	b := Board{}
	b.width = width
	b.data = populateData(width)
	return b
}

func populateData(w int) [][]int {
	data := make([][]int, w)
	var counter = 1
	var upbound = w*w
	for i,_ := range data {
		data[i] = make([]int, w)
		for j,_ := range data[i] {
			data[i][j] = counter
			counter += 1
			if counter == upbound {
				break
			}
		}
	}
	return data 
}

func (b *Board) Move(d coord.Direction) error {
	var di, dj = d.Coord()
	var ex, ey, er = b.findEmptySpace()
	if er != nil {
		return er
	}
	if ex-di < 0 || ex-di >= b.width {
		return errors.New(fmt.Sprintf("Move out of bounds for width %d", b.width))
	}
	if ey-dj < 0 || ey-dj >= b.width {
		return errors.New(fmt.Sprintf("Move out of bounds for width %d", b.width))
	}
	b.data[ex][ey] = b.data[ex-di][ey-dj]
	b.data[ex-di][ey-dj] = 0
	return nil
}

func (b *Board) findEmptySpace() (int, int, error) {
	for i:=0; i<b.width; i++ {
		for j:=0; j<b.width; j++ {
			if b.data[i][j] == 0 {
				return i,j, nil
			}
		}
	}
	return -1, -1, errors.New("Couldn't find empty space")
}


func (b *Board) reset() {
	b.data = populateData(b.width)
}

func (b *Board) Shuffle() {
	var w = b.width
	var numbers = linearizeTable(b.data, b.width)
	shuffleArray(&numbers)
	b.data = arrayToTable(numbers, w)
}

func (b *Board) Width() int {
	return b.width
}

func (b *Board) Data() [][]int {
	return b.data
}

func shuffleArray(data *[]int) {
	for i:=0; i < len(*data)-1; i++ {
		var j = rand.IntN(i+1)
		var temp = (*data)[i]
		(*data)[i] = (*data)[j]
		(*data)[j] = temp
	}
}

func linearizeTable(data [][]int, w int) []int {
	numbers := make([]int, w*w)
	for i,r := range data {
		for j,v := range r {
			numbers[i*w+j] = v
		}
	}
	return numbers
}

func arrayToTable(array []int, w int) [][]int {
	data := make([][]int, w)
	var i = 0
	var j = 0
	for k,e := range array {
		i = k/w
		j = k%w
		if j == 0 {
			data[i] = make([]int, w)
		}
		data[i][j] = e
	}
	return data
}

