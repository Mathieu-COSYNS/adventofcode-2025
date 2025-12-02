package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func readInputFile() ([]Range, error) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	// We will capture close errors and combine them
	defer func() {
		cerr := inputFile.Close()
		if err == nil { // only overwrite if we had no previous error
			err = cerr
		}
	}()

	var ranges []Range
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		for _, rangeString := range split {
			r := strings.Split(rangeString, "-")

			if len(r) != 2 {
				return nil, errors.New("a wrong range exist")
			}

			start, err := strconv.Atoi(string(r[0]))
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(string(r[1]))
			if err != nil {
				return nil, err
			}

			ranges = append(ranges, Range{
				start,
				end,
			})

		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ranges, err
}

func isInvalidIdInPart1(id int) bool {
	idStr := fmt.Sprint(id)
	if len(idStr)%2 == 1 {
		return false
	}
	strIdHalfLen := len(idStr) / 2
	part1 := idStr[:strIdHalfLen]
	part2 := idStr[strIdHalfLen:]
	return part1 == part2
}

func isInvalidIdInPart2(id int) bool {
	idStr := fmt.Sprint(id)
	idStrLen := len(idStr)
	for i := 1; i < idStrLen; i++ {
		if idStrLen%i != 0 {
			continue
		}
		repeated := strings.Repeat(idStr[:i], idStrLen/i)

		if repeated == idStr {
			return true
		}
	}

	return false
}

func sumInvalidIDs(ranges []Range, isInvalidId func(int) bool) int {
	sum := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if isInvalidId(i) {
				sum += i
			}
		}
	}

	return sum
}

func main() {
	ranges, err := readInputFile()
	if err != nil {
		panic(err)
	}

	fmt.Println(sumInvalidIDs(ranges, isInvalidIdInPart1))
	fmt.Println(sumInvalidIDs(ranges, isInvalidIdInPart2))
}
