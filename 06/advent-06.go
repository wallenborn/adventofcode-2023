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
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
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
	check(err)
	scanner := bufio.NewScanner(file)
	result := 0
	var time []int
	var distance []int
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		if strings.HasPrefix(line, "Time:") {
			time = parseRaces(line)
		}
		if strings.HasPrefix(line, "Distance:") {
			distance = parseRaces(line)
		}
	}
	check(scanner.Err())
	ways := numberOfWays(time, distance)
	if len(ways) != 0 {
		result = 1
		for _, x := range ways {
			result *= x
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
	var time []int
	var distance []int
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		if strings.HasPrefix(line, "Time:") {
			time = parseOneRace(line)
		}
		if strings.HasPrefix(line, "Distance:") {
			distance = parseOneRace(line)
		}
	}
	check(scanner.Err())
	ways := numberOfWays(time, distance)
	if len(ways) != 0 {
		result = 1
		for _, x := range ways {
			result *= x
		}
	}
	logger.Info("Finished", "result", result)
}

func parseRaces(line string) []int {
	var result []int
	tmp := strings.Fields(line)
	for i := range tmp {
		if i == 0 {
			continue
		}
		n, err := strconv.Atoi(tmp[i])
		check(err)
		result = append(result, n)
	}
	return result
}

func parseOneRace(line string) []int {
	var result []int
	s := strings.ReplaceAll(line, " ", "")
	tmp := strings.Split(s, ":")
	for i := range tmp {
		if i == 0 {
			continue
		}
		n, err := strconv.Atoi(tmp[i])
		check(err)
		result = append(result, n)
	}
	return result
}

func numberOfWays(time []int, distance []int) []int {
	var result []int
	for i := range time {
		t := time[i]
		d := distance[i]
		num := 0
		for j := 0; j <= t; j++ {
			dd := j * (t - j)
			logger.Debug("Race result", "race", i, "record", d, "time", j, "distance", dd)
			if dd > d {
				num++
			}
		}
		result = append(result, num)
	}
	return result
}
