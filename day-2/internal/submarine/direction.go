package submarine

import (
	"errors"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
	Unknown
)

func (dir Direction) String() string {
	switch dir {
	case Forward:
		return "forward"
	case Up:
		return "up"
	case Down:
		return "down"
	default:
		return "unknown"
	}
}

func MakeDirection(strDir string) (Direction, error) {
	switch strings.ToLower(strDir) {
	case "forward":
		return Forward, nil
	case "up":
		return Up, nil
	case "down":
		return Down, nil
	}

	return Unknown, errors.New("failed to parse Direction from string")
}
