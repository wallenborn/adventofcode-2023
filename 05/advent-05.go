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

	scanner := bufio.NewScanner(file)
	allMappers := make(map[string]Mapper)
	var thisMapper *CompositeMapper
	var allSeeds []int
	findNumbers, err := regexp.Compile(`\d+`)
	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		if strings.HasPrefix(line, "seeds:") {
			logger.Debug("Found seeds")
			tmp := findNumbers.FindAllString(line, -1)
			for _, s := range tmp {
				i, err := strconv.Atoi(s)
				check(err)
				allSeeds = append(allSeeds, i)
			}
		} else if strings.HasSuffix(line, "map:") {
			name := strings.Split(line, " ")[0]
			thisMapper = &CompositeMapper{name: name}
		} else if line != "" {
			tmp := findNumbers.FindAllString(line, -1)
			dst, err := strconv.Atoi(tmp[0])
			check(err)
			src, err := strconv.Atoi(tmp[1])
			check(err)
			len, err := strconv.Atoi(tmp[2])
			check(err)
			thisMapper.add(&RangeMapper{dst: dst, src: src, len: len})
		} else if thisMapper != nil {
			logger.Debug("This Mapper", "name", thisMapper.name)
			allMappers[thisMapper.name] = thisMapper
		}
	}
	logger.Debug("This Mapper", "name", thisMapper.name)
	allMappers[thisMapper.name] = thisMapper

	for i, seed := range allSeeds {
		soil := allMappers["seed-to-soil"].lookup(seed)
		fertilizer := allMappers["soil-to-fertilizer"].lookup(soil)
		water := allMappers["fertilizer-to-water"].lookup(fertilizer)
		light := allMappers["water-to-light"].lookup(water)
		temperature := allMappers["light-to-temperature"].lookup(light)
		humidity := allMappers["temperature-to-humidity"].lookup(temperature)
		location := allMappers["humidity-to-location"].lookup(humidity)
		logger.Debug("Seed", "seed", seed, "soil", soil, "location", location)
		if i == 0 || location < result {
			result = location
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
	allMappers := make(map[string]Mapper)
	var thisMapper *CompositeMapper
	var allSeeds []int
	findNumbers, err := regexp.Compile(`\d+`)
	check(err)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		if strings.HasPrefix(line, "seeds:") {
			logger.Debug("Found seeds")
			tmp := findNumbers.FindAllString(line, -1)
			for _, s := range tmp {
				i, err := strconv.Atoi(s)
				check(err)
				allSeeds = append(allSeeds, i)
			}
		} else if strings.HasSuffix(line, "map:") {
			name := strings.Split(line, " ")[0]
			thisMapper = &CompositeMapper{name: name}
		} else if line != "" {
			tmp := findNumbers.FindAllString(line, -1)
			dst, err := strconv.Atoi(tmp[0])
			check(err)
			src, err := strconv.Atoi(tmp[1])
			check(err)
			len, err := strconv.Atoi(tmp[2])
			check(err)
			thisMapper.add(&RangeMapper{dst: dst, src: src, len: len})
		} else if thisMapper != nil {
			logger.Debug("This Mapper", "name", thisMapper.name)
			allMappers[thisMapper.name] = thisMapper
		}
	}
	logger.Debug("This Mapper", "name", thisMapper.name)
	allMappers[thisMapper.name] = thisMapper

	for i := 0; i < len(allSeeds); i += 2 {
		start := allSeeds[i]
		length := allSeeds[i+1]
		logger.Info("Looping over seeds", "start", start, "length", length)
		for seed := start; seed <= start+length; seed++ {
			soil := allMappers["seed-to-soil"].lookup(seed)
			fertilizer := allMappers["soil-to-fertilizer"].lookup(soil)
			water := allMappers["fertilizer-to-water"].lookup(fertilizer)
			light := allMappers["water-to-light"].lookup(water)
			temperature := allMappers["light-to-temperature"].lookup(light)
			humidity := allMappers["temperature-to-humidity"].lookup(temperature)
			location := allMappers["humidity-to-location"].lookup(humidity)
			logger.Debug("Seed", "seed", seed, "soil", soil, "location", location)
			if result == 0 || location < result {
				result = location
			}
		}
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)
}
