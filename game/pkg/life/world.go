package life

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height < 0 || width < 0 {
		return nil, errors.New("invalid data")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) Next(x, y int) bool {
	n := w.neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}

	if n == 3 && !alive {
		return true
	}

	return false
}

func (w *World) neighbors(x, y int) int {
	count := 0
	neighboringOffsets := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, offset := range neighboringOffsets {
		neighborX := (x + offset[1] + w.Width) % w.Width
		neighborY := (y + offset[0] + w.Height) % w.Height

		if w.Cells[neighborY][neighborX] {
			count++
		}
	}

	return count
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
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
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
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

func (w World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	var sb strings.Builder
	for _, row := range w.Cells {
		for _, cell := range row {
			if cell {
				sb.WriteString(greenSquare)
			} else {
				sb.WriteString(brownSquare)
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.Cells[y][x] {
				file.WriteString("1")
			} else {
				file.WriteString("0")
			}
		}
		if y != w.Height-1 {
			file.WriteString("\n")
		}
	}

	return nil
}

func (w *World) LoadState(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	height := len(lines)
	if height == 0 {
		return errors.New("Empty file")
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			return errors.New("Width")
		}
	}

	w.Height = height
	w.Width = width

	w.Cells = make([][]bool, height)
	for i := 0; i < height; i++ {
		w.Cells[i] = make([]bool, width)
	}

	for i, line := range lines {
		for j, char := range line {
			if char == '1' {
				w.Cells[i][j] = true
			} else if char != '0' {
				return errors.New("Invalid character")
			}
		}
	}

	return nil
}
