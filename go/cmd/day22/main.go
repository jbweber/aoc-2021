package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var input = `on x=-20..26,y=-36..17,z=-47..7
on x=-20..33,y=-21..23,z=-26..28
on x=-22..28,y=-29..23,z=-38..16
on x=-46..7,y=-6..46,z=-50..-1
on x=-49..1,y=-3..46,z=-24..28
on x=2..47,y=-22..22,z=-23..27
on x=-27..23,y=-28..26,z=-21..29
on x=-39..5,y=-6..47,z=-3..44
on x=-30..21,y=-8..43,z=-13..34
on x=-22..26,y=-27..20,z=-29..19
off x=-48..-32,y=26..41,z=-47..-37
on x=-12..35,y=6..50,z=-50..-2
off x=-48..-32,y=-32..-16,z=-15..-5
on x=-18..26,y=-33..15,z=-7..46
off x=-40..-22,y=-38..-28,z=23..41
on x=-16..35,y=-41..10,z=-47..6
off x=-32..-23,y=11..30,z=-14..3
on x=-49..-5,y=-3..45,z=-29..18
off x=18..30,y=-20..-8,z=-3..13
on x=-41..9,y=-7..43,z=-33..15
on x=-54112..-39298,y=-85059..-49293,z=-27449..7877
on x=967..23432,y=45373..81175,z=27513..53682`

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() { f.Close() }()

	rss, err := parseInput(f)
	if err != nil {
		panic(err)
	}

	reactor := map[string]bool{}
	processRebootSteps(rss, reactor)

	fmt.Println(countOn(reactor))
}

func countOn(reactor map[string]bool) int {
	count := 0
	for _, v := range reactor {
		if v {
			count++
		}
	}

	return count
}

func processRebootSteps(steps []RebootStep, reactor map[string]bool) {
	for _, step := range steps {
		for x := step.XStart; x <= step.XStop; x++ {

			if x < -50 || x > 50 {
				continue
			}

			for y := step.YStart; y <= step.YStop; y++ {

				if y < -50 || y > 50 {
					continue
				}

				for z := step.ZStart; z <= step.ZStop; z++ {

					if z < -50 || z > 50 {
						continue
					}

					k := fmt.Sprintf("%d,%d,%d", x, y, z)
					reactor[k] = step.Command
				}
			}
		}
	}
}

func parseInput(input io.Reader) ([]RebootStep, error) {
	scanner := bufio.NewScanner(input)

	rebootSteps := make([]RebootStep, 0)

	for scanner.Scan() {
		line := scanner.Text()

		rs := RebootStep{}

		rs.Command = strings.HasPrefix(line, "on ")

		line = strings.TrimPrefix(line, "on ")
		line = strings.TrimPrefix(line, "off ")

		points := strings.Split(line, ",")

		points[0] = strings.TrimPrefix(points[0], "x=")
		points[1] = strings.TrimPrefix(points[1], "y=")
		points[2] = strings.TrimPrefix(points[2], "z=")

		xss := strings.Split(points[0], "..")
		rs.XStart, _ = strconv.Atoi(xss[0])
		rs.XStop, _ = strconv.Atoi(xss[1])

		yss := strings.Split(points[1], "..")
		rs.YStart, _ = strconv.Atoi(yss[0])
		rs.YStop, _ = strconv.Atoi(yss[1])

		zss := strings.Split(points[2], "..")
		rs.ZStart, _ = strconv.Atoi(zss[0])
		rs.ZStop, _ = strconv.Atoi(zss[1])

		rebootSteps = append(rebootSteps, rs)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rebootSteps, nil
}

type RebootStep struct {
	Command bool
	XStart  int
	XStop   int
	YStart  int
	YStop   int
	ZStart  int
	ZStop   int
}
