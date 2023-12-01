package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
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
	fmt.Println("Part One, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstDigit, err := regexp.Compile(`.*?(\d).*`)
	check(err)
	lastDigit, err := regexp.Compile(`.*(\d).*?`)
	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := firstDigit.FindStringSubmatch(line)
		last := lastDigit.FindStringSubmatch(line)
		summand := fmt.Sprintf("%s%s", first[1], last[1])
		i, err := strconv.Atoi(summand)
		check(err)
		result += i
		logger.Debug(line, "first", first[1], "last", last[1], "summand", summand, "sum", result)
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)
}

func partTwo(filename string) {
	fmt.Println("Part Two")
}
