package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

var (
	logger   slog.Logger     = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	allnodes map[string]node = make(map[string]node)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node struct {
	name  string
	left  string
	right string
}

func main() {
	inputFile := flag.String("input", "input.txt", "Input for this puzzle")
	part := flag.String("part", "all", "Part of the puzzle to solve [one|two|all]")
	flag.Parse()

	if *part == "all" || *part == "one" {
		partOne(*inputFile)
	}
	if *part == "all" || *part == "two" {
		partTwo(*inputFile)
	}
}

func partOne(filename string) {
	logger.Info("Part One, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	scanner.Scan()
	path := scanner.Text()
	findTurn, err := regexp.Compile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	check(err)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		tmp := findTurn.FindStringSubmatch(line)
		allnodes[tmp[1]] = node{name: tmp[1], left: tmp[2], right: tmp[3]}
	}
	check(scanner.Err())
	result = pathLength(path, "AAA")
	logger.Info("Finished", "result", result)
}

func partTwo(filename string) {
	logger.Info("Part Two, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	scanner.Scan()
	var currentnodes []string
	path := scanner.Text()
	findTurn, err := regexp.Compile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
	check(err)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		tmp := findTurn.FindStringSubmatch(line)
		name := tmp[1]
		left := tmp[2]
		right := tmp[3]
		allnodes[name] = node{name: name, left: left, right: right}
		if strings.HasSuffix(name, "A") {
			currentnodes = append(currentnodes, name)
		}
	}
	check(scanner.Err())
	var pathLen []int
	for _, v := range currentnodes {
		thisLength := pathLength(path, v)
		pathLen = append(pathLen, thisLength)
		logger.Debug("Found path", "start", v, "length", thisLength)
	}
	result = LCM(pathLen)
	logger.Info("Finished", "result", result)
}

func pathLength(path string, start string) int {
	here := start
	result := 0
	i := 0
	for !strings.HasSuffix(here, "Z") {
		t := path[i]
		n := allnodes[here]
		if t == 'L' {
			here = n.left
		} else if t == 'R' {
			here = n.right
		} else {
			logger.Error("Unknown direction", "turn", t)
		}
		result++
		i++
		if i >= len(path) {
			i = 0
		}
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers []int) int {
	result := LCMint(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		result = LCMint(result, integers[i])
	}
	return result
}

func LCMint(a int, b int) int {
	return a * b / GCD(a, b)
}
