package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var test = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var test2 = `11111
19991
19191
19991
11111`

var myInput = `1326253315
3427728113
5751612542
6543868322
4422526221
2234325647
1773174887
7281321674
6562513118
4824541522`

func main() {
	dumbos := parseInput(myInput)

	count := 0
	for {
		applyStep(dumbos)
		count++
		if allFlashed(dumbos) {
			break
		}
	}

	fmt.Println(count)
	format(dumbos)

}

func applyStep(dumbos [][]dumbo) {
	increase(dumbos)
	flashed := flash(dumbos)
	fmt.Println(flashed)
}

func increase(dumbos [][]dumbo) {
	height := len(dumbos)
	width := len(dumbos[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dumbos[y][x].energy++
			dumbos[y][x].flashed = false
		}
	}
}

func flash(dumbos [][]dumbo) bool {
	height := len(dumbos)
	width := len(dumbos[0])

	flashed := false

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			if dumbos[y][x].energy > 9 && !dumbos[y][x].flashed {
				dumbos[y][x].energy = 0
				dumbos[y][x].flashes++
				dumbos[y][x].flashed = true
				prop(x, y, dumbos)
				flashed = true
			}
		}
	}

	if flashed {
		flashed = flash(dumbos)
	}

	return flashed
}

func prop(x, y int, dumbos [][]dumbo) {
	height := len(dumbos)
	width := len(dumbos[0])

	// above y+1, x
	if y+1 >= 0 && y+1 < height {
		if !dumbos[y+1][x].flashed {
			dumbos[y+1][x].energy++
		}
	}

	// below y-1, x
	if y-1 >= 0 && y-1 < height {
		if !dumbos[y-1][x].flashed {
			dumbos[y-1][x].energy++
		}
	}

	// left  y, x-1
	if x-1 >= 0 && x-1 < width {
		if !dumbos[y][x-1].flashed {
			dumbos[y][x-1].energy++
		}
	}

	// right y, x+1
	if x+1 >= 0 && x+1 < width {
		if !dumbos[y][x+1].flashed {
			dumbos[y][x+1].energy++
		}
	}

	// top left		y+1, x-1
	if y+1 >= 0 && y+1 < height && x-1 >= 0 && x-1 < width {
		if !dumbos[y+1][x-1].flashed {
			dumbos[y+1][x-1].energy++
		}
	}

	// top right	y+1, x+1
	if y+1 >= 0 && y+1 < height && x+1 >= 0 && x+1 < width {
		if !dumbos[y+1][x+1].flashed {
			dumbos[y+1][x+1].energy++
		}
	}

	// bottom left	y-1, x-1
	if y-1 >= 0 && y-1 < height && x-1 >= 0 && x-1 < width {
		if !dumbos[y-1][x-1].flashed {
			dumbos[y-1][x-1].energy++
		}
	}

	// bottom right y-1, x+1
	if y-1 >= 0 && y-1 < height && x+1 >= 0 && x+1 < width {
		if !dumbos[y-1][x+1].flashed {
			dumbos[y-1][x+1].energy++
		}
	}
}

func parseInput(input string) [][]dumbo {
	dumbos := make([][]dumbo, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := make([]dumbo, 0)
		for _, ch := range line {
			d, _ := strconv.Atoi(string(ch))

			row = append(row, dumbo{d, false, 0})
		}

		dumbos = append(dumbos, row)
	}

	return dumbos
}

type dumbo struct {
	energy  int
	flashed bool
	flashes int
}

func format(dumbos [][]dumbo) {
	for _, row := range dumbos {
		for _, dumbo := range row {
			if dumbo.energy == 0 {
				red := color.New(color.FgRed)
				boldRed := red.Add(color.Bold)
				boldRed.Printf("%d", dumbo.energy)
			} else {
				fmt.Printf("%d", dumbo.energy)
			}
		}
		fmt.Println()
	}
}

func countFlashes(dumbos [][]dumbo) int {
	flashes := 0

	height := len(dumbos)
	width := len(dumbos[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			flashes += dumbos[y][x].flashes
		}
	}

	return flashes
}

func allFlashed(dumbos [][]dumbo) bool {
	height := len(dumbos)
	width := len(dumbos[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dumbos[y][x].energy != 0 {
				return false
			}
		}
	}

	return true
}
