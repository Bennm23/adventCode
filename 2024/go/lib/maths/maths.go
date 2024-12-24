package maths

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

type Number interface {
	int | int8 | int16 | int32 | int64
}

func Gcd[T Number](a, b T) T {
	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
}

func Lcm[T Number](a, b T) T {
	return (a * b) / Gcd(a, b)
}

func LcmRange[T Number](vals ...T) T {
	if len(vals) < 2 {
		panic("LCM RANGE TO SMALL")
	}
	var res T = Lcm(vals[0], vals[1])

	if len(vals) == 2 {
		return res
	}

	for _, val := range vals[2:] {
		res = Lcm(res, val)
	}

	return res
}

func Transpose[T any](matrix [][]T) [][]T {
	transposed := make([][]T, len(matrix[0]))

	for i := range matrix[0] {
		transposed[i] = make([]T, len(matrix))
	}

	for i, row := range matrix {

		for j := range row {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func GenerateHash(values ...interface{}) uint64 {
	hash := fnv.New64a()

	for _, val := range values {
		hash.Write([]byte(fmt.Sprintf("%v", val)))
	}

	return hash.Sum64()
}

func ToInt(s string) int {
	res, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return res
}

type Position struct {
	X int
	Y int
}
var ALL_MOVES = []Position{
	{X: -1, Y: 0}, //N
	{X: 0, Y: 1},  //E
	{X: 1, Y: 0},  //S
	{X: 0, Y: -1}, //W
	{X: -1, Y: 1}, //NE
	{X: 1, Y: 1},  //SE
	{X: 1, Y: -1}, //SW
	{X: -1, Y: -1},//NW
}
const (
    N = iota
    E
    S
    W
    NE
    SE
    SW
    NW
)
func GetNeighbors(position Position) []Position {
	neighbors := make([]Position, 0)

	for _, move := range ALL_MOVES {
		neighbors = append(neighbors, move.Add(position))
	}
	return neighbors
}


func (p Position) InBounds(size int) bool {
	return p.X < size && p.X >= 0 && p.Y < size && p.Y >= 0
}

func (p *Position) PAdd(o Position) {
	p.X += o.X
	p.Y += o.Y
}
func (p *Position) PMinus(o Position) {
	p.X -= o.X
	p.Y -= o.Y
}
func ToInt64(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}
	return res
}

func (p Position) OutOfBounds(size int) bool {
	return p.X < 0 || p.X >= size || p.Y < 0 || p.Y >= size
}

func (p Position) EvaluateFor(grid [][]int) int {
	if p.OutOfBounds(len(grid)) {
		panic("GRID OUT OF BOUNDS")
	}
	return grid[p.X][p.Y]
}

func (p Position) Add(other Position) Position {
	return Position{
		p.X + other.X,
		p.Y + other.Y,
	}
}

func NewPosition(row int, col int) Position {
	return Position{row, col}
}

func CountDigits(num int) int {
	if num == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(math.Abs(float64(num))))) + 1
}
func numEvenDigits(num int) bool {
	return CountDigits(num)%2 == 0
}
func splitNum(num int) (int, int) {
	tens := int(math.Pow10(CountDigits(num) / 2))

	return num / tens, num % tens
}

var HORIZONTAL_MOVES = []Position{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

func (p Position) Distance(other Position) float64 {
	return math.Sqrt((math.Pow(float64(other.X - p.X), 2) + math.Pow(float64(other.Y - p.Y), 2)))
}

func InitTypeGrid[T any](val T, height, width int) [][]T {
	grid := make([][]T, 0)

	for range height {
		row := make([]T, 0)
		for range width {
			row = append(row, val)
		}
		grid = append(grid, row)
	}
	return grid
}

type AnyNum interface {
	int | float32 | float64
}

func Max[T AnyNum](a, b T) T {
	if a >= b {
		return a
	}

	return b
}

func (p *Position) Times(scalar int) {
	p.X *= scalar
	p.Y *= scalar
}

func (p Position) TurnClockwise() Position {
	if p.X == -1 {
		return NewPosition(0, 1)
	} else if p.X == 1 {
		return NewPosition(0, -1)
	} else if p.Y == -1 {
		return NewPosition(-1, 0)
	} else {
		return NewPosition(1, 0)
	}
}
func (p Position) TurnCounterClockwise() Position {
	counterClockwise := p.TurnClockwise()
	counterClockwise.Times(-1)
	return counterClockwise
}