package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	cards  []string    = make([]string, 0)
	count  []int       = make([]int, 0)
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

	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		n := countWinners(line)
		if n > 0 {
			p := 1
			for i := 1; i < n; i++ {
				p = p << 1
			}
			result += p
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
	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, line)
		check(err)
		count = append(count, 1)
	}
	logger.Debug("Read cards", "cards", len(cards))
	for i, line := range cards {
		logger.Debug("Scanning for cards", "card", i, "total", sum(count))
		winners := countWinners(line)
		for j := 0; j < winners; j++ {
			count[cardNo(line)+j] += count[i]
		}
	}
	check(scanner.Err())
	result = sum(count)
	logger.Info("Finished", "result", result)
}

func sum(ary []int) int {
	result := 0
	for _, v := range ary {
		result += v
	}
	return result
}

func cardNo(line string) int {
	numbers, err := regexp.Compile(`Card +(\d+):`)
	check(err)
	tmp := numbers.FindStringSubmatch(line)
	n, err := strconv.Atoi(tmp[1])
	check(err)
	return n
}

func countWinners(line string) int {
	numbers, err := regexp.Compile(`(\d+)`)
	check(err)
	allNumbers := strings.Split(line, ":")[1]
	tmp := strings.Split(allNumbers, "|")
	winningNumbers := numbers.FindAllString(tmp[0], -1)
	ourNumbers := numbers.FindAllString(tmp[1], -1)
	//logger.Debug("", "winning", winningNumbers, "ours", ourNumbers)
	return countCommonElements(winningNumbers, ourNumbers)
}

func countCommonElements(a, b []string) int {
	result := 0
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				result += 1
			}
		}
	}
	return result
}
