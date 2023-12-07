package main

import (
	"bufio"
	"flag"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	logger slog.Logger = *slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type hand struct {
	cards string
	bet   int
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
	var allhands []hand
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		tmp := strings.Fields(line)
		amt, err := strconv.Atoi(tmp[1])
		check(err)
		h := hand{cards: tmp[0], bet: amt}
		allhands = append(allhands, h)
	}
	for _, h := range allhands {
		logger.Debug("Unsorted", "cards", h.cards, "rank", handRank(h.cards))
	}
	slices.SortFunc(allhands, compareHands)
	for i, h := range allhands {
		logger.Debug("Sorted", "cards", h.cards, "rank", handRank(h.cards))
		result += h.bet * (i + 1)
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
	var allhands []hand
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug(line)
		tmp := strings.Fields(line)
		amt, err := strconv.Atoi(tmp[1])
		check(err)
		h := hand{cards: tmp[0], bet: amt}
		allhands = append(allhands, h)
	}
	for _, h := range allhands {
		logger.Debug("Unsorted", "cards", h.cards, "rank", handRank(h.cards))
	}
	slices.SortFunc(allhands, compareHandsWithJoker)
	for i, h := range allhands {
		logger.Debug("Sorted", "cards", h.cards, "rank", handRank(h.cards))
		result += h.bet * (i + 1)
	}
	check(scanner.Err())
	logger.Info("Finished", "result", result)
}

func compareHands(a, b hand) int {
	ac := a.cards
	bc := b.cards
	if handRank(ac) > handRank(bc) {
		return 1
	} else if handRank(ac) < handRank(bc) {
		return -1
	} else {
		for i := range ac {
			ar, _ := utf8.DecodeRuneInString(ac[i:])
			br, _ := utf8.DecodeRuneInString(bc[i:])
			if cardRank(ar) > cardRank(br) {
				return 1
			} else if cardRank(ar) < cardRank(br) {
				return -1
			}
		}
	}
	return 0
}

func compareHandsWithJoker(a, b hand) int {
	ac := a.cards
	bc := b.cards
	ah := maxHandRank(a.cards)
	bh := maxHandRank(b.cards)
	if ah > bh {
		return 1
	} else if ah < bh {
		return -1
	} else {
		for i := range ac {
			ar, _ := utf8.DecodeRuneInString(ac[i:])
			br, _ := utf8.DecodeRuneInString(bc[i:])
			if cardRankWithJoker(ar) > cardRankWithJoker(br) {
				return 1
			} else if cardRankWithJoker(ar) < cardRankWithJoker(br) {
				return -1
			}
		}
	}
	return 0
}

func maxHandRank(cards string) int {
	vals := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	maxRank := handRank(cards)
	for _, r := range vals {
		s := strings.ReplaceAll(cards, "J", r)
		sr := handRank(s)
		if sr > maxRank {
			maxRank = sr
		}
	}
	return maxRank
}

func handRank(cards string) int {
	if isHighCard(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "HighCard", "rank", 1)
		return 1
	} else if isOnePair(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "OnePair", "rank", 2)
		return 2
	} else if isTwoPair(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "TwoPair", "rank", 3)
		return 3
	} else if isThreeOfAKind(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "ThreeOfAKind", "rank", 4)
		return 4
	} else if isFullHouse(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "isFullHouse", "rank", 5)
		return 5
	} else if isFourOfAKind(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "FourOfAKind", "rank", 6)
		return 6
	} else if isFiveOfAKind(cards) {
		//logger.Debug("Hand Rank", "cards", cards, "hand", "FiveOfAKind", "rank", 7)
		return 7
	}
	logger.Error("Can't assign hand rank", "cards", cards)
	return 0
}

func cardRank(c rune) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		logger.Error("Can't assign card rank", "card", c)
		return 0
	}
}

func cardRankWithJoker(c rune) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 1 // Joker is low now
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		logger.Error("Can't assign card rank", "card", c)
		return 0
	}

}

func isHighCard(cards string) bool {
	return len(mapCards(cards)) == len(cards)
}

func isOnePair(cards string) bool {
	return len(mapCards(cards)) == len(cards)-1
}

func isTwoPair(cards string) bool {
	n := len(cards)
	m := mapCards(cards)
	return len(m) == n-2 && count(m, 2) == 2 && count(m, 1) == n-4
}

func isThreeOfAKind(cards string) bool {
	n := len(cards)
	m := mapCards(cards)
	return len(m) == n-2 && count(m, 3) == 1 && count(m, 2) == 0 && count(m, 1) == n-3
}

func isFullHouse(cards string) bool {
	n := len(cards)
	m := mapCards(cards)
	return len(m) == n-3 && count(m, 3) == 1 && count(m, 2) == 1 && count(m, 1) == n-5
}

func isFourOfAKind(cards string) bool {
	n := len(cards)
	m := mapCards(cards)
	return len(m) == n-3 && count(m, 4) == 1 && count(m, 3) == 0 && count(m, 2) == 0 && count(m, 1) == n-4
}

func isFiveOfAKind(cards string) bool {
	n := len(cards)
	m := mapCards(cards)
	return len(m) == n-4 && count(m, 5) == 1 && count(m, 4) == 0 && count(m, 3) == 0 && count(m, 2) == 0 && count(m, 1) == n-5
}

func mapCards(cards string) map[rune]int {
	r := make(map[rune]int)
	for _, c := range cards {
		_, found := r[c]
		if found {
			r[c] = r[c] + 1
		} else {
			r[c] = 1
		}
	}
	return r
}

func count(m map[rune]int, n int) int {
	c := 0
	for _, v := range m {
		if v == n {
			c++
		}
	}
	return c
}
