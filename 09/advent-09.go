package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var (
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		elms := parseLine(line)
		tmp := findNextElement(elms)
		result += tmp
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)
}

func partTwo(filename string) {
	logger.Info("Part Two, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		elms := parseLine(line)
		tmp := findPreviousElement(elms)
		result += tmp
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)
}

func findNextElement(elms []int) int {
	if allZeroes(elms) {
		return 0
	}
	var tmp []int
	for i := 0; i < len(elms)-1; i++ {
		tmp = append(tmp, elms[i+1]-elms[i])
	}
	n := findNextElement(tmp)
	return elms[len(elms)-1] + n
}

func findPreviousElement(elms []int) int {
	if allZeroes(elms) {
		return 0
	}
	var tmp []int
	for i := 0; i < len(elms)-1; i++ {
		tmp = append(tmp, elms[i+1]-elms[i])
	}
	n := findPreviousElement(tmp)
	return elms[0] - n
}

func allZeroes(elms []int) bool {
	for _, v := range elms {
		if v != 0 {
			return false
		}
	}
	return true
}

func parseLine(line string) []int {
	var res []int
	tmp := strings.Fields(line)
	for _, v := range tmp {
		n, err := strconv.Atoi(v)
		check(err)
		res = append(res, n)
	}
	return res
}
