package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
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
	}
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
	}
	logger.Info("Finished", "result", result)
}

func Count(str string, pattern []int) int {
	logger.Debug("Count called", "str", str, "pattern", pattern)
	if len(str) == 0 {
		logger.Error("Count called with zero-length string", "pattern", pattern)
	} else if len(str) == 1 {
		if str == "#" {
			if len(pattern) == 1 && pattern[0] == 1 {
				logger.Info("Found match", "str", str, "pattern", pattern)
				return 1
			} else {
				return 0
			}
		} else if str == "." {
			if len(pattern) == 1 && pattern[0] == 0 {
				return Count(str, []int{})
			} else if len(pattern) == 0 {
				logger.Info("Found match", "str", str, "pattern", pattern)
				return 1
			} else {
				return 0
			}
		} else if str == "?" {
			return Count("#", pattern) + Count(".", pattern)
		}
	}

	if str[0] == '#' {
		var tmp []int
		if len(pattern) == 0 || pattern[0] == 0 {
			return 0
		} else {
			tmp = make([]int, len(pattern))
			copy(tmp, pattern)
			tmp[0] = pattern[0] - 1
		}
		return Count(str[1:], tmp)
	} else if str[0] == '.' {
		var tmp []int
		if len(pattern) > 0 && pattern[0] == 0 {
			tmp = make([]int, len(pattern)-1)
			copy(tmp, pattern[1:])
			return Count(str, tmp)
		} else {
			return Count(str[1:], pattern)
		}
	} else if str[0] == '?' {
		return Count("#"+str[1:], pattern) + Count("."+str[1:], pattern)
	}
	logger.Error("Unreachable code", "str", str, "pattern", pattern)
	return 0
}
