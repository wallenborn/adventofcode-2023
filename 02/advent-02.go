package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var (
	logger slog.Logger    = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	cubes  map[string]int = map[string]int{"red": 12, "green": 13, "blue": 14}
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
	fmt.Println("Cubes", cubes)
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		valid := true
		fields := strings.Split(line, ":")
		gameNo, err := strconv.Atoi(strings.Split(fields[0], " ")[1])
		check(err)
		logger.Debug(line, "game", gameNo, "result", result)
		draws := strings.Split(fields[1], ";")
		for _, d := range draws {
			for _, c := range strings.Split(d, ",") {
				tmp := strings.Split(c, " ")
				num, err := strconv.Atoi(tmp[1])
				check(err)
				color := tmp[2]
				if num > cubes[color] {
					valid = false
				}
			}
		}
		if valid {
			result += gameNo
		}
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
		cubes := make(map[string]int)
		fields := strings.Split(line, ":")
		draws := strings.Split(fields[1], ";")
		for _, d := range draws {
			for _, c := range strings.Split(d, ",") {
				tmp := strings.Split(c, " ")
				num, err := strconv.Atoi(tmp[1])
				check(err)
				color := tmp[2]
				if num > cubes[color] {
					cubes[color] = num
				}
			}
		}
		power := 1
		for _, v := range cubes {
			power *= v
		}
		logger.Debug(line, "result", result, "power", power)
		result += power
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)

}
