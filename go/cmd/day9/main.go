package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"strconv"
)

var input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() { f.Close() }()

	floor, err := parseInput(f)
	if err != nil {
		panic(err)
	}

	setLow(floor)

	//for _, row := range floor {
	//	fmt.Println(row)
	//}

	//riskLevel := int32(0)
	//
	//for _, row := range floor {
	//	for _, column := range row {
	//		if column.Low {
	//			riskLevel = riskLevel + column.Height + 1
	//		}
	//	}
	//}

	format(floor)

}

func format(floor [][]FloorSpace) {
	for _, row := range floor {
		for _, column := range row {
			if column.Low {
				red := color.New(color.FgRed)
				boldRed := red.Add(color.Bold)
				boldRed.Printf("%d", column.Height)
			} else {
				fmt.Printf("%d", column.Height)
			}
		}
		fmt.Println()
	}
}

func setLow(floor [][]FloorSpace) {
	height := len(floor)
	width := len(floor[0])

	riskSum := int32(0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			current := floor[y][x]

			// above y+1, x
			if y+1 >= 0 && y+1 < height {
				above := floor[y+1][x]

				if above.Height < current.Height {
					continue
				}
			}

			// below y-1, x
			if y-1 >= 0 && y-1 < height {
				below := floor[y-1][x]

				if below.Height <= current.Height {
					continue
				}
			}

			// left  y, x-1
			if x-1 >= 0 && x-1 < width {
				left := floor[y][x-1]

				if left.Height <= current.Height {
					continue
				}
			}

			// right y, x+1
			if x+1 >= 0 && x+1 < width {
				right := floor[y][x+1]

				if right.Height <= current.Height {
					continue
				}
			}

			current.Low = true

			floor[y][x] = current
			riskSum = riskSum + 1 + current.Height
		}
	}

	fmt.Println(riskSum)
}

func parseInput(input io.Reader) ([][]FloorSpace, error) {
	scanner := bufio.NewScanner(input)

	floor := make([][]FloorSpace, 0)

	for scanner.Scan() {
		line := scanner.Text()

		row := make([]FloorSpace, 0)

		for _, c := range line {

			d, _ := strconv.Atoi(string(c))

			row = append(row, FloorSpace{int32(d), false})
		}

		floor = append(floor, row)
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return floor, nil
}

type FloorSpace struct {
	Height int32
	Low    bool
}
