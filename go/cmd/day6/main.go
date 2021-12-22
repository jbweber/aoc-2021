package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "5,3,2,2,1,1,4,1,5,5,1,3,1,5,1,2,1,4,1,2,1,2,1,4,2,4,1,5,1,3,5,4,3,3,1,4,1,3,4,4,1,5,4,3,3,2,5,1,1,3,1,4,3,2,2,3,1,3,1,3,1,5,3,5,1,3,1,4,2,1,4,1,5,5,5,2,4,2,1,4,1,3,5,5,1,4,1,1,4,2,2,1,3,1,1,1,1,3,4,1,4,1,1,1,4,4,4,1,3,1,3,4,1,4,1,2,2,2,5,4,1,3,1,2,1,4,1,4,5,2,4,5,4,1,2,1,4,2,2,2,1,3,5,2,5,1,1,4,5,4,3,2,4,1,5,2,2,5,1,4,1,5,1,3,5,1,2,1,1,1,5,4,4,5,1,1,1,4,1,3,3,5,5,1,5,2,1,1,3,1,1,3,2,3,4,4,1,5,5,3,2,1,1,1,4,3,1,3,3,1,1,2,2,1,2,2,2,1,1,5,1,2,2,5,2,4,1,1,2,4,1,2,3,4,1,2,1,2,4,2,1,1,5,3,1,4,4,4,1,5,2,3,4,4,1,5,1,2,2,4,1,1,2,1,1,1,1,5,1,3,3,1,1,1,1,4,1,2,2,5,1,2,1,3,4,1,3,4,3,3,1,1,5,5,5,2,4,3,1,4"

func main() {
	part2(256)
}

func part2(days int) {
	//fishes := map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 0, 6: 0, 7: 0, 8: 0}
	fishes := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	inputs := strings.Split(input, ",")
	for _, i := range inputs {
		j, _ := strconv.Atoi(i)
		fishes[j]++
	}
	for days > 0 {
		curr0 := fishes[0]
		curr1 := fishes[1]
		curr2 := fishes[2]
		curr3 := fishes[3]
		curr4 := fishes[4]
		curr5 := fishes[5]
		curr6 := fishes[6]
		curr7 := fishes[7]
		curr8 := fishes[8]

		fishes[8] = curr0
		fishes[7] = curr8
		fishes[6] = curr0 + curr7
		fishes[5] = curr6
		fishes[4] = curr5
		fishes[3] = curr4
		fishes[2] = curr3
		fishes[1] = curr2
		fishes[0] = curr1

		days--
	}

	total := 0
	for _, v := range fishes {
		total += v
	}

	fmt.Println(total)
}

func part1() {
	lanternFish := []Fish{Fish{3}, Fish{4}, Fish{3}, Fish{1}, Fish{2}}

	//lanternFish := []Fish{}
	//
	//inputs := strings.Split(input, ",")
	//for _, i := range inputs {
	//	j, _ := strconv.Atoi(i)
	//	lanternFish = append(lanternFish, Fish{j})
	//}

	simulate(lanternFish, 256)
}

func simulate(fishes []Fish, days int) {
	for days > 0 {
		l := len(fishes)
		for i := 0; i < l; i++ {
			if fishes[i].State == 0 {
				fishes[i].State = 6
				fishes = append(fishes, Fish{8})
			} else {
				fishes[i].State--
			}
		}
		days--
	}
	fmt.Println(len(fishes))
}

type Fish struct {
	State int
}
