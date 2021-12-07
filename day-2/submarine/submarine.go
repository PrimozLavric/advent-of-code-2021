package submarine

import (
	"errors"
)

type Submarine struct {
	x   int
	y   int
	aim int
}

func (sub *Submarine) PositionX() int {
	return sub.x
}

func (sub *Submarine) PositionY() int {
	return sub.y
}

func (sub *Submarine) Aim() int {
	return sub.aim
}

func (sub *Submarine) MultipliedPositions() int {
	return sub.x * sub.y
}

func (sub *Submarine) MovePartOne(dir Direction, distance uint) error {
	switch dir {
	case Forward:
		sub.x += int(distance)
	case Up:
		sub.y -= int(distance)
	case Down:
		sub.y += int(distance)
	default:
		return errors.New("unknown direction")
	}

	return nil
}

func (sub *Submarine) MovePartTwo(dir Direction, distance uint) error {
	switch dir {
	case Forward:
		sub.x += int(distance)
		sub.y += int(distance) * sub.aim
	case Up:
		sub.aim -= int(distance)
	case Down:
		sub.aim += int(distance)
	default:
		return errors.New("unknown direction")
	}

	return nil
}
