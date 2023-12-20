package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type MatchPos struct {
  firstPos int
  lastPos int
  matchValue string

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

func main() {
    lines, err := readLines("d1p1")
    if err != nil {
        log.Fatalf("Failed to read lines %s", err)
    }

    namedNumberMap := map[string]string{
          "zero": "0",
          "one": "1",
          "two": "2",
          "three": "3",
          "four": "4",
          "five": "5",
          "six": "6",
          "seven": "7",
          "eight": "8",
          "nine": "9",
    }

    keys := make([]string, 0, len(namedNumberMap))
    for k := range namedNumberMap {
      keys = append(keys, k)
    }
    checkValues := []string{ "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", }
    // add named numbers to check list
    checkValues = append(checkValues, keys...)


    sum := 0
    for _, line := range lines {
        var matches []MatchPos

        for _, checkValue := range checkValues {
            if strings.Contains(line, checkValue) {
              matchPos := MatchPos{
                firstPos: strings.Index(line, checkValue),
                lastPos: strings.LastIndex(line, checkValue),
                matchValue: checkValue,
              }
              matches = append(matches, matchPos)
              log.Printf("matchValue %s - firstPos %d lastPos %d", matchPos.matchValue, matchPos.firstPos, matchPos.lastPos)
            }
        }
        
        first := matches[0]
        last := matches[0]
        // find the fist
        for _, match := range matches {
           if match.firstPos < first.firstPos {
              first = match
            }
           if match.lastPos > last.lastPos {
              last = match
            }
        }

        firstDigit := first.matchValue
        // map named numbers to ints
        if len(firstDigit) > 1 {
          firstDigit = namedNumberMap[firstDigit]
        }
        lastDigit := last.matchValue
        if len(lastDigit) > 1 {
          lastDigit = namedNumberMap[lastDigit]
        }

        combinedDigit, err := strconv.Atoi(string(firstDigit + lastDigit))
        if err != nil {
          log.Fatalf("Failed to convert string to it %s", err)
        }

        log.Printf(`%s -> first: %s last: %s`, line, firstDigit, lastDigit)

        // add it to the sum
        sum += combinedDigit
    }

    log.Printf("Totoal is %d", sum)
}
