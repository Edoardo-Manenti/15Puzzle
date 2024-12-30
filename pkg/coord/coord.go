package coord

import (
	"errors"
	"fmt"
)

type Direction struct {
	di int
	dj int
}

const (
	UP int = iota
	RIGHT 
	DOWN 
	LEFT 
)

func GetDirection(x int) (Direction, error) {
	var d = Direction{0,0}
	switch x {
	case UP : 
		d = Direction{-1, 0}
	case RIGHT : 
		d = Direction{0, 1}
	case DOWN : 
		d = Direction{1, 0}
	case LEFT : 
		d = Direction{0, -1}
	}
	if d.di == 0 && d.dj == 0 {
		return d, errors.New(fmt.Sprintf("No Direction found for value %d", x))
	}
	return d, nil
}

func (d *Direction) Coord() (int, int) {
	return d.di, d.dj
}
