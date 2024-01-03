package life

import (
	"math/rand"
	"strings"
	"time"
)

const (
	brownSquare = "\xF0\x9F\x9F\xAB"
	greenSquare = "\xF0\x9F\x9F\xA9"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) String() string {
	var lines []string

	for i := 0; i < w.Height; i++ {
		current := ""
		for j := 0; j < w.Width; j++ {
			s := brownSquare
			if w.Cells[i][j] {
				s = greenSquare
			}
			current += s
		}
		lines = append(lines, current)
	}

	return strings.Join(lines, "\n")
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) *World {
	matrix := createMatrix(height, width)
	return &World{
		Height: height,
		Width:  width,
		Cells:  matrix,
	}
}

func createMatrix(height int, width int) [][]bool {
	matrix := make([][]bool, height)
	rows := make([]bool, height*width)
	for i := 0; i < height; i++ {
		matrix[i] = rows[i*width : (i+1)*width]
	}
	return matrix
}
func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)
	return n == 3 || (w.Cells[y][x] && n >= 2 && n <= 3)
}
func (w *World) neighbors(x, y int) (s int) {
	reminder := func(a, b int) int {
		return (a%b + b) % b
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx := reminder(x+j, w.Width)
			ny := reminder(y+i, w.Height)
			if nx >= 0 && ny >= 0 && nx < w.Width && ny < w.Height && w.Cells[ny][nx] {
				s++
			}
		}
	}
	return
}
func NextState(oldWorld, newWorld *World) {
	for i := 0; i < newWorld.Height; i++ {
		for j := 0; j < newWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

func Assign(receiver, source *World) {
	for i := 0; i < source.Height; i++ {
		for j := 0; j < source.Width; j++ {
			receiver.Cells[i][j] = source.Cells[i][j]
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {

			// randRowLeft := r.Intn(w.Height)
			// randColLeft := r.Intn(w.Width)
			randRowRight := r.Intn(w.Height)
			randColRight := r.Intn(w.Width)

			// t := w.Cells[randRowLeft][randColLeft]
			// w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]

			t := w.Cells[i][j]
			w.Cells[i][j] = w.Cells[randRowRight][randColRight]
			w.Cells[randRowRight][randColRight] = t
		}
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
