package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
  id int
  rounds []Round
  possible bool
}

type Round struct {
  red int
  blue int
  green int
  possible bool
}

type MaxRanges struct {
  red int
  blue int
  green int
}


// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func extractColorCount(str string, color string) int {
      reColor := regexp.MustCompile(`([\d]+) ` + color)
      colorCountFind := reColor.FindStringSubmatch(str)
      if len(colorCountFind) > 1 {
	colorCount, err := strconv.Atoi(colorCountFind[1])
	if err != nil {
	  log.Fatalf("Could not parse %s count to int %s", color, err)
	}
	return colorCount
      }
    return 0
}

func isValidRound(round Round) bool {
  if round.blue > maxRanges.blue {
    return false
  }
  if round.red > maxRanges.red {
    return false
  }
  if round.green > maxRanges.green {
    return false
  }

  return true
}

func solvePartOne() {
    lines, err := readLines("d2")
    if err != nil {
        log.Fatalf("Failed to read lines %s", err)
    }

    maxRanges = MaxRanges {
      red: 12,
      blue: 14,
      green: 13,
    }

    sum := 0
    for _, line := range lines {
      log.Print(line)
      var game Game
      // always start out thinking the game is possible
      game.possible = true
      reGameId :=  regexp.MustCompile(`^Game ([\d]+):`)
      gameIdStr := reGameId.FindStringSubmatch(line)[1]
      gameId, err := strconv.Atoi(gameIdStr)

      if err != nil {
	log.Fatalf("Could not parse gameId to int %s", err)
      }

      game.id = gameId

      roundsIdx := strings.Index(line, ":") + 1
      roundsStr := line[roundsIdx:]
      rounds := strings.Split(roundsStr, ";")
      for _, roundStr := range rounds {
	var round Round
	round.blue = extractColorCount(roundStr, "blue")
	round.red = extractColorCount(roundStr, "red")
	round.green = extractColorCount(roundStr, "green")
	round.possible = isValidRound(round)

	if !round.possible {
	  game.possible = false
	}

	game.rounds = append(game.rounds, round)
    }
    log.Printf("game: %+v", game)
    if game.possible {
      sum += game.id
    }
  }

    log.Printf("Total is %d", sum)
}

func solvePartTwo() {
    lines, err := readLines("input")
    if err != nil {
        log.Fatalf("Failed to read lines %s", err)
    }

    maxRanges = MaxRanges {
      red: 12,
      blue: 14,
      green: 13,
    }

    sum := 0
    for _, line := range lines {
      log.Print(line)
      var game Game
      // always start out thinking the game is possible
      game.possible = true
      reGameId :=  regexp.MustCompile(`^Game ([\d]+):`)
      gameIdStr := reGameId.FindStringSubmatch(line)[1]
      gameId, err := strconv.Atoi(gameIdStr)

      if err != nil {
	log.Fatalf("Could not parse gameId to int %s", err)
      }

      game.id = gameId

      roundsIdx := strings.Index(line, ":") + 1
      roundsStr := line[roundsIdx:]
      rounds := strings.Split(roundsStr, ";")
      for _, roundStr := range rounds {
	var round Round
	round.blue = extractColorCount(roundStr, "blue")
	round.red = extractColorCount(roundStr, "red")
	round.green = extractColorCount(roundStr, "green")
	round.possible = isValidRound(round)

	if !round.possible {
	  game.possible = false
	}

	game.rounds = append(game.rounds, round)
    }
    log.Printf("game: %+v", game)
    if game.possible {
      sum += game.id
    }
  }

    log.Printf("Total is %d", sum)
}

var (
  maxRanges MaxRanges
)

func main() {
  solvePartTwo()
}
