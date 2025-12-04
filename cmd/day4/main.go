package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputFile() ([][]bool, error) {
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

	scanner := bufio.NewScanner(inputFile)
	banks := [][]bool{}
	for scanner.Scan() {
		line := scanner.Text()

		bank := []bool{}
		for _, char := range line {
			bank = append(bank, char == '@')

		}
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, err
}

func hasFewerThanFourAdjacentRolls(i int, j int, rolls [][]bool) bool {
	x := i - 1
	if x < 0 {
		x = 0
	}

	xmax := i + 1
	if xmax >= len(rolls) {
		xmax = i
	}

	count := 0

	for ; x <= xmax; x++ {
		y := j - 1
		if y < 0 {
			y = 0
		}

		ymax := j + 1
		if ymax >= len(rolls[x]) {
			ymax = j
		}

		for ; y <= ymax; y++ {
			if x == i && y == j {
				continue
			}
			if rolls[x][y] {
				count++
				if count >= 4 {
					return false
				}
			}
		}
	}

	return true
}

func part1(rollsMatrix [][]bool) int {
	sum := 0
	for i, rollsLine := range rollsMatrix {
		for j := range rollsLine {
			if rollsMatrix[i][j] {
				if hasFewerThanFourAdjacentRolls(i, j, rollsMatrix) {
					sum++
				}
			}
		}
	}
	return sum
}

func part2(rollsMatrix [][]bool) int {
	sum := 0
	hasRemovedSomeRolls := true
	for hasRemovedSomeRolls {
		hasRemovedSomeRolls = false
		newRollsMatrix := [][]bool{}
		for i, rollsLine := range rollsMatrix {
			newRollsMatrix = append(newRollsMatrix, []bool{})
			for j := range rollsLine {
				if rollsMatrix[i][j] {
					if hasFewerThanFourAdjacentRolls(i, j, rollsMatrix) {
						sum++
						hasRemovedSomeRolls = true
						newRollsMatrix[i] = append(newRollsMatrix[i], false)
					} else {
						newRollsMatrix[i] = append(newRollsMatrix[i], true)
					}
				} else {
					newRollsMatrix[i] = append(newRollsMatrix[i], false)
				}
			}
		}
		rollsMatrix = newRollsMatrix
	}
	return sum
}

func main() {
	rollsMatrix, err := readInputFile()
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(rollsMatrix))
	fmt.Println(part2(rollsMatrix))
}
