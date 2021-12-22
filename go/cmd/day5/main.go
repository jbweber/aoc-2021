package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sample = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func main() {
	q1()
}

func q1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	pairs, err := parseInput(f)
	if err != nil {
		panic(err)
	}

	mapping := make(map[Point]int)

	for _, pair := range pairs {
		mapping[pair.A]++

		for pair.A.X != pair.B.X || pair.A.Y != pair.B.Y {
			if pair.A.X > pair.B.X {
				pair.A.X--
			} else if pair.A.X < pair.B.X {
				pair.A.X++
			}

			if pair.A.Y > pair.B.Y {
				pair.A.Y--
			} else if pair.A.Y < pair.B.Y {
				pair.A.Y++
			}

			mapping[pair.A]++
		}
	}

	sum := 0
	for _, v := range mapping {
		if v >= 2 {
			sum++
		}
	}

	fmt.Println(sum)
}

func parseInput(input io.Reader) ([]Pair, error) {
	scanner := bufio.NewScanner(input)

	pairs := make([]Pair, 0)

	for scanner.Scan() {
		line := scanner.Text()

		points := strings.Split(line, " -> ")

		a := strings.Split(points[0], ",")
		b := strings.Split(points[1], ",")

		ax, _ := strconv.Atoi(a[0])
		ay, _ := strconv.Atoi(a[1])

		bx, _ := strconv.Atoi(b[0])
		by, _ := strconv.Atoi(b[1])

		pair := Pair{
			Point{ax, ay},
			Point{bx, by},
		}

		//if pair.IsHorizontalOrVertical() {
		pairs = append(pairs, pair)
		//}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pairs, nil
}

type Pair struct {
	A Point
	B Point
}

type Point struct {
	X int
	Y int
}

func (p Pair) IsHorizontalOrVertical() bool {
	if p.A.X == p.B.X || p.A.Y == p.B.Y {
		return true
	}

	return false
}
