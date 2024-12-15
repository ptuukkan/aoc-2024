package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Convert []byte to string
	text := string(fileContent)
	return text, nil
}

func SplitNewLines(input string) []string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

type Point struct {
	Y int
	X int
}

func (p Point) Add(o Point) Point {
	return Point{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func (p *Point) Move(v Point) {
	p.Y += v.Y
	p.X += v.X
}

func NewPoint(y, x int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

var Directions = []Point{
	NewPoint(-1, 0),
	NewPoint(0, 1),
	NewPoint(1, 0),
	NewPoint(0, -1),
}

func (p *Point) OutOfBounds(length int) bool {
	if p.X < 0 || p.Y < 0 || p.X >= length || p.Y >= length {
		return true
	}
	return false
}

func (p Point) Up() Point {
	return Point{
		Y: p.Y + 1,
		X: p.X,
	}
}

func (p Point) Right() Point {
	return Point{
		Y: p.Y,
		X: p.X + 1,
	}
}

func (p Point) Down() Point {
	return Point{
		Y: p.Y - 1,
		X: p.X,
	}
}

func (p Point) Left() Point {
	return Point{
		Y: p.Y,
		X: p.X - 1,
	}
}
