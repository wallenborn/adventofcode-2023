package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
)

var (
	logger   slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	stars    []point
	emptyrow []bool
	emptycol []bool
)

type point struct {
	row int
	col int
}

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
	row := 0
	findStars, err := regexp.Compile(`#`)
	check(err)
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		emptyrow = append(emptyrow, false)
		if len(emptycol) == 0 {
			for range line {
				emptycol = append(emptycol, true)
			}
		}
		allStars := findStars.FindAllStringSubmatchIndex(line, -1)
		if len(allStars) == 0 {
			emptyrow[row] = true
		}
		for _, val := range allStars {
			emptycol[val[0]] = false
			p := point{row: row, col: val[0]}
			stars = append(stars, p)
		}
		row++
	}
	check(scanner.Err())
	logger.Debug("Before expansion", "stars", stars)
	for i := len(emptycol) - 1; i >= 0; i-- {
		if emptycol[i] {
			for j, s := range stars {
				if s.col > i {
					stars[j] = point{row: s.row, col: s.col + 1}
				}
			}
		}
	}
	for i := len(emptyrow) - 1; i >= 0; i-- {
		if emptyrow[i] {
			for j, s := range stars {
				if s.row > i {
					stars[j] = point{row: s.row + 1, col: s.col}
				}
			}
		}
	}
	logger.Debug("After expansion", "stars", stars)
	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			result += manhattanDistance(stars[i], stars[j])
		}
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
	row := 0
	findStars, err := regexp.Compile(`#`)
	check(err)
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		emptyrow = append(emptyrow, false)
		if len(emptycol) == 0 {
			for range line {
				emptycol = append(emptycol, true)
			}
		}
		allStars := findStars.FindAllStringSubmatchIndex(line, -1)
		if len(allStars) == 0 {
			emptyrow[row] = true
		}
		for _, val := range allStars {
			emptycol[val[0]] = false
			p := point{row: row, col: val[0]}
			stars = append(stars, p)
		}
		row++
	}
	check(scanner.Err())
	logger.Debug("Before expansion", "stars", stars)
	expansion := 1000000 - 1

	for i := len(emptycol) - 1; i >= 0; i-- {
		if emptycol[i] {
			for j, s := range stars {
				if s.col > i {
					stars[j] = point{row: s.row, col: s.col + expansion}
				}
			}
		}
	}
	for i := len(emptyrow) - 1; i >= 0; i-- {
		if emptyrow[i] {
			for j, s := range stars {
				if s.row > i {
					stars[j] = point{row: s.row + expansion, col: s.col}
				}
			}
		}
	}
	logger.Debug("After expansion", "stars", stars)
	k := 0
	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			k++
			dist := manhattanDistance(stars[i], stars[j])
			result += manhattanDistance(stars[i], stars[j])
			logger.Debug("Next pair", "num", k, "i", i+1, "j", j+1, "dist", dist)
		}
	}
	logger.Info("Finished", "result", result)
}

func manhattanDistance(a, b point) int {
	d := absDiff(a.row, b.row) + absDiff(a.col, b.col)
	return d
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
