package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var input = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() { f.Close() }()

	err = parseInput(f)
	if err != nil {
		panic(err)
	}
}

func parseInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "|")

		parts[1] = strings.TrimSpace(parts[1])

		digits := strings.Split(parts[1], " ")

		for _, digit := range digits {
			dl := len(digit)

			if dl == 2 || dl == 3 || dl == 4 || dl == 7 {
				fmt.Printf("%s ", digit)
				count++
			}
		}

		fmt.Println()
	}

	err := scanner.Err()
	if err != nil {
		return err
	}

	fmt.Println(count)

	return nil
}
