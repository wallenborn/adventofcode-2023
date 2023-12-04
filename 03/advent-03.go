package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var (
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	field  []string    = make([]string, 0)
	gears  [][][]int   = make([][][]int, 0)
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
	findNumbers, err := regexp.Compile(`\d+`)
	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
		logger.Debug(line, "numlines", len(field))
	}
	check(scanner.Err())

	for n, line := range field {
		allNumbers := findNumbers.FindAllStringSubmatchIndex(line, -1)
		for _, match := range allNumbers {
			i := match[0]
			j := match[1]
			number, err := strconv.Atoi(line[i:j])
			check(err)
			logger.Debug("Found number", "num", number)
			if isPartNumber(number, n, i, j) {
				logger.Debug("Found part number", "line", line, "number", number)
				result += number
			}
		}
	}
	check(err)
	logger.Info("Finished", "result", result)
}

func partTwo(filename string) {
	logger.Info("Part Two, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	findNumbers, err := regexp.Compile(`\d+`)
	check(err)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
		gearline := make([][]int, len(line))
		gears = append(gears, gearline)
		logger.Debug(line, "numlines", len(field))
	}
	check(scanner.Err())

	for n, line := range field {
		allNumbers := findNumbers.FindAllStringSubmatchIndex(line, -1)
		for _, match := range allNumbers {
			i := match[0]
			j := match[1]
			number, err := strconv.Atoi(line[i:j])
			check(err)
			logger.Debug("Found number", "num", number)
			if isPartNumber(number, n, i, j) {
				logger.Debug("Found part number", "line", line, "number", number)
			}
		}
	}
	result := 0
	for _, row := range gears {
		for _, col := range row {
			if col != nil {
				if len(col) == 2 {
					power := col[0] * col[1]
					logger.Debug("Found gear", "num1", col[0], "num2", col[1], "power", power)
					result += power
				}
			}
		}
	}
	check(err)
	logger.Info("Finished", "result", result)
}

func isPartNumber(number, n, i, j int) bool {
	logger.Debug("scanning line", "n", n, "i", i, "j", j)
	result := false
	if n > 0 {
		line := field[n-1]
		for k := max(0, i-1); k < min(j+1, len(line)); k++ {
			ch := line[k]
			logger.Debug(line, "column", k, "char", ch)
			if !isBlank(ch) && !isDigit(ch) {
				result = true
				if isAsterisk(ch) {
					if gears[n-1][k] == nil {
						gears[n-1][k] = make([]int, 0)
					}
					gears[n-1][k] = append(gears[n-1][k], number)
				}
			}
		}
	}
	line := field[n]
	if i > 0 {
		ch := line[i-1]
		logger.Debug(line, "column", i-1, "char", ch)
		if !isBlank(ch) && !isDigit(ch) {
			result = true
			if isAsterisk(ch) {
				if gears[n][i-1] == nil {
					gears[n][i-1] = make([]int, 0)
				}
				gears[n][i-1] = append(gears[n][i-1], number)
			}
		}
	}
	if j < len(line)-1 {
		ch := line[j]
		logger.Debug(line, "column", j, "char", ch)
		if !isBlank(ch) && !isDigit(ch) {
			result = true
			if isAsterisk(ch) {
				if gears[n][j] == nil {
					gears[n][j] = make([]int, 0)
				}
				gears[n][j] = append(gears[n][j], number)
			}
		}
	}
	if n < len(field)-1 {
		line := field[n+1]
		for k := max(0, i-1); k < min(j+1, len(line)); k++ {
			ch := line[k]
			logger.Debug(line, "column", k, "char", ch)
			if !isBlank(ch) && !isDigit(ch) {
				result = true
				if isAsterisk(ch) {
					if gears[n+1][k] == nil {
						gears[n+1][k] = make([]int, 0)
					}
					gears[n+1][k] = append(gears[n+1][k], number)
				}
			}
		}
	}
	return result
}

func isDigit(ch byte) bool {
	return ch >= 48 && ch <= 57 // ASCII for '0'-'9'
}

func isBlank(ch byte) bool {
	return ch == 46 // ASCII for '.'
}

func isAsterisk(ch byte) bool {
	return ch == 42 // ASCII for '*'
}
