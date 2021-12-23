package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var closeForOpen = map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
var openForClose = map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
var openers = map[string]struct{}{"(": {}, "[": {}, "{": {}, "<": {}}
var points = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137, "": 0}
var completionPoints = map[string]int{")": 1, "]": 2, "}": 3, ">": 4, "": 0}

func main() {
	//	input := `[({(<(())[]>[[{[]{<()<>>
	//[(()[<>])]({[<{<<[]>>(
	//{([(<{}[<>[]}>{[]{[(<()>
	//(((({<>}<{<{<>}{[]{[]{}
	//[[<[([]))<([[{}[[()]]]
	//[{[{({}]{}}([{[{{{}}([]
	//{<[[]]>}<{[{[{[]{()[[[]
	//[<(<(<(<{}))><([]([]()
	//<{([([[(<>()){}]>(<<{{
	//<{([{{}}[<[[[<>{}]]]>[]]`

	fb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fb)
	lines := strings.Split(input, "\n")
	completionScores := make([]int, 0)
	for _, line := range lines {
		incomplete, completion := syntaxFix(line)

		if incomplete {
			score := calculateCompletion(completion)
			completionScores = append(completionScores, score)
		}

	}

	sort.Ints(completionScores)

	var middle int = len(completionScores) / 2
	fmt.Println(completionScores)
	fmt.Println(completionScores[middle])
}

func syntaxFix(input string) (bool, string) {
	if len(input) == 0 {
		return false, ""
	}

	startsWith, input := string(input[0]), input[1:]

	openChunks := []string{startsWith}

	corruptFinisher := ""

	for _, ch := range input {
		next := string(ch)

		// check if we're a new opener
		if _, ok := openers[next]; ok {
			openChunks = append(openChunks, next)
			continue
		}

		// ensure open chunks len > 0??

		lastOpener := openChunks[len(openChunks)-1]

		// check if we're closing the last opener
		if closeForOpen[lastOpener] == next {
			openChunks = openChunks[:len(openChunks)-1]
			continue
		}

		corruptFinisher = next
		break
	}

	if len(openChunks) > 0 && corruptFinisher != "" {
		return false, ""
	}

	if len(openChunks) > 0 && corruptFinisher == "" {
		fixer := strings.Builder{}

		for i := len(openChunks) - 1; i >= 0; i-- {
			fixer.WriteString(closeForOpen[openChunks[i]])
		}

		return true, fixer.String()
	}

	return false, ""
}

func calculateCompletion(completion string) int {
	score := 0

	for _, ch := range completion {
		st := string(ch)

		score *= 5
		score += completionPoints[st]
	}

	return score
}

func main2() {
	//	input := `[({(<(())[]>[[{[]{<()<>>
	//[(()[<>])]({[<{<<[]>>(
	//{([(<{}[<>[]}>{[]{[(<()>
	//(((({<>}<{<{<>}{[]{[]{}
	//[[<[([]))<([[{}[[()]]]
	//[{[{({}]{}}([{[{{{}}([]
	//{<[[]]>}<{[{[{[]{()[[[]
	//[<(<(<(<{}))><([]([]()
	//<{([([[(<>()){}]>(<<{{
	//<{([{{}}[<[[[<>{}]]]>[]]`

	fb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fb)

	illegalPoints := 0
	for _, line := range strings.Split(input, "\n") {
		incomplete, corrupt, illegalChar := syntaxCheck(line)

		if incomplete {
			continue
		}

		if corrupt {
			illegalPoints += points[illegalChar]
		}
	}

	fmt.Println(illegalPoints)

}

func syntaxCheck(input string) (bool, bool, string) {
	if len(input) == 0 {
		return false, false, ""
	}

	startsWith, input := string(input[0]), input[1:]

	openChunks := []string{startsWith}

	corruptFinisher := ""

	for _, ch := range input {
		next := string(ch)

		// check if we're a new opener
		if _, ok := openers[next]; ok {
			openChunks = append(openChunks, next)
			continue
		}

		// ensure open chunks len > 0??

		lastOpener := openChunks[len(openChunks)-1]

		// check if we're closing the last opener
		if closeForOpen[lastOpener] == next {
			openChunks = openChunks[:len(openChunks)-1]
			continue
		}

		corruptFinisher = next
		break
	}

	if len(openChunks) > 0 && corruptFinisher != "" {
		return false, true, corruptFinisher
	}

	if len(openChunks) > 0 && corruptFinisher == "" {
		return true, false, ""
	}

	return false, false, ""
}
