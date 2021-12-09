package line

// Point represents a 2D point.
type Point struct {
	X int
	Y int
}

// Line represents a 2D line that is defined by two points.
type Line struct {
	A Point
	B Point
}

// NewLine creates a new line with the provided point coordinates.
func NewLine(x1, y1, x2, y2 int) *Line {
	line := Line{
		Point{X: x1, Y: y1},
		Point{X: x2, Y: y2},
	}

	return &line
}

// MaxX returns maximal x position of the line.
func (l *Line) MaxX() int {
	if l.A.X > l.B.X {
		return l.A.X
	}

	return l.B.X
}

// MaxY finds maximal y position of the line.
func (l *Line) MaxY() int {
	if l.A.Y > l.B.Y {
		return l.A.Y
	}

	return l.B.Y
}

// IsHorizontal checks if the line is horizontal.
func (l *Line) IsHorizontal() bool {
	return l.A.X == l.B.X
}

// IsVertical checks if the line is vertical.
func (l *Line) IsVertical() bool {
	return l.A.Y == l.B.Y
}

func FindMaxXY(lines []*Line) (int, int) {
	maxX := 0
	maxY := 0

	for _, l := range lines {
		lineMaxX := l.MaxX()
		if lineMaxX > maxX {
			maxX = lineMaxX
		}

		lineMaxY := l.MaxY()
		if lineMaxY > maxY {
			maxY = lineMaxY
		}
	}

	return maxX, maxY
}
