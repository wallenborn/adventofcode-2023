package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"strings"
)

var (
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	maze   []string
	path   map[point]bool = make(map[point]bool)
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
	start := point{row: -1, col: -1}
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		maze = append(maze, line)
		col := strings.Index(line, "S")
		if col > -1 {
			start = point{row: row, col: col}
		}
		row++
	}
	check(scanner.Err())
	if start.row == -1 && start.col == -1 {
		logger.Error("No starting point found")
	}
	logger.Debug("Starting path search", "start", start)
	current := findNextPoint(start, start)
	previous := start
	dist := 1
	for current != start {
		tmp := current
		current = findNextPoint(previous, current)
		previous = tmp
		dist++
	}
	result = dist / 2
	logger.Info("Finished", "result", result)
}

func findNextPoint(previous, current point) point {
	c := maze[current.row][current.col]
	logger.Debug("Path search", "previous", previous, "current", current, "char", string(c))
	switch c {
	case 'S':
		if current.row > 0 { // LOOK UP
			cc := maze[current.row-1][current.col]
			if cc == '|' || cc == '7' || cc == 'F' {
				return point{row: current.row - 1, col: current.col}
			}
		}
		if current.col < len(maze[0]) { // LOOK RIGHT
			cc := maze[current.row][current.col+1]
			if cc == '-' || cc == '7' || cc == 'J' {
				return point{row: current.row, col: current.col + 1}
			}
		}
		if current.row < len(maze) { // LOOK DOWN
			cc := maze[current.row+1][current.col]
			if cc == '|' || cc == 'J' || cc == 'L' {
				return point{row: current.row + 1, col: current.col}
			}
		}
		if current.col > 0 { // LOOK LEFT
			cc := maze[current.row][current.col-1]
			if cc == '-' || cc == 'F' || cc == 'L' {
				return point{row: current.row, col: current.col - 1}
			}
		}
	case '|':
		if previous.row == current.row+1 { // GO UP
			return point{row: current.row - 1, col: current.col}
		} else if previous.row == current.row-1 { // GO DOWN
			return point{row: current.row + 1, col: current.col}
		}
	case 'J':
		if previous.col == current.col-1 { // GO UP
			return point{row: current.row - 1, col: current.col}
		} else if previous.row == current.row-1 { // GO LEFT
			return point{row: current.row, col: current.col - 1}
		}
	case 'L':
		if previous.col == current.col+1 { // GO UP
			return point{row: current.row - 1, col: current.col}
		} else if previous.row == current.row-1 { // GO RIGHT
			return point{row: current.row, col: current.col + 1}
		}
	case 'F':
		if previous.row == current.row+1 { // GO RIGHT
			return point{row: current.row, col: current.col + 1}
		} else if previous.col == current.col+1 { // GO DOWN
			return point{row: current.row + 1, col: current.col}
		}
	case '7':
		if previous.col == current.col-1 { // GO DOWN
			return point{row: current.row + 1, col: current.col}
		} else if previous.row == current.row+1 { // GO LEFT
			return point{row: current.row, col: current.col - 1}
		}
	case '-':
		if previous.col == current.col-1 { // GO RIGHT
			return point{row: current.row, col: current.col + 1}
		} else if previous.col == current.col+1 { // GO LEFT
			return point{row: current.row, col: current.col - 1}
		}
	}
	logger.Error("Should not be here")
	return current
}

func partTwo(filename string) {
	logger.Info("Part Two, reading from " + filename)
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	row := 0
	start := point{row: -1, col: -1}
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		maze = append(maze, line)
		col := strings.Index(line, "S")
		if col > -1 {
			start = point{row: row, col: col}
		}
		row++
	}
	check(scanner.Err())
	if start.row == -1 && start.col == -1 {
		logger.Error("No starting point found")
	}
	logger.Debug("Starting path search", "start", start)
	path[start] = true
	current := findNextPoint(start, start)
	path[current] = true
	outbound := current
	previous := start
	for current != start {
		tmp := current
		current = findNextPoint(previous, current)
		path[current] = true
		previous = tmp
	}
	maze[start.row] = strings.Replace(maze[start.row], "S", fixStart(start, outbound, previous), 1)
	for i := range maze {
		for j := range maze[i] {
			current := point{row: i, col: j}
			if path[current] {
				continue
			}
			crossed := 0
			cc := maze[i][j]
			for k := j - 1; k >= 0; k-- {
				p := point{row: i, col: k}
				c := maze[i][k]
				if !path[p] {
					cc = c
					continue
				}
				if c == '-' {
					continue
				}
				if c == 'F' && cc == 'J' {
					cc = c
					continue
				}
				if c == 'L' && cc == '7' {
					cc = c
					continue
				}
				crossed++
				cc = c
			}
			if crossed%2 != 0 {
				logger.Debug("Point is inside the loop", "point", current)
				result++
			}
		}
	}
	logger.Info("Finished", "result", result)
}

func fixStart(start, outbound, previous point) string {
	var up, ri, do, le bool
	if outbound.row == start.row-1 || previous.row == start.row-1 {
		up = true
	}
	if outbound.col == start.col+1 || previous.col == start.col+1 {
		ri = true
	}
	if outbound.row == start.row+1 || previous.row == start.row+1 {
		do = true
	}
	if outbound.col == start.col-1 || previous.col == start.col-1 {
		le = true
	}
	if up && ri {
		return "L"
	}
	if up && do {
		return "|"
	}
	if up && le {
		return "J"
	}
	if ri && do {
		return "F"
	}
	if ri && le {
		return "-"
	}
	if do && le {
		return "7"
	}
	logger.Error("Start not connected", "start", start, "out", outbound, "in", previous)
	return "."
}
