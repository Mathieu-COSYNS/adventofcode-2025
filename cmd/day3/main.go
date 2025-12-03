package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInputFile() ([][]int, error) {
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
	banks := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		bank := []int{}
		for _, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}

			bank = append(bank, value)

		}
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, err
}

func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func getBankJoltage(bank []int, numberOfDigitsToTake int) int {
	highestJoltages := make([]int, numberOfDigitsToTake)
	highestJoltagesIndexes := make([]int, numberOfDigitsToTake)

	for j := 0; j < numberOfDigitsToTake; j++ {
		minIndex := -1
		if j > 0 {
			minIndex = highestJoltagesIndexes[j-1]
		}
		for i := len(bank) - numberOfDigitsToTake + j; i > minIndex; i-- {
			if bank[i] >= highestJoltages[j] {
				highestJoltages[j] = bank[i]
				highestJoltagesIndexes[j] = i
			}
		}
	}

	sum := 0
	for i, joltages := range highestJoltages {
		sum += joltages * powInt(10, (numberOfDigitsToTake-i-1))
	}

	return sum

	// highestJoltage := 0
	// highestJoltageIndex := 0
	// for i := len(bank) - 2; i >= 0; i-- {
	// 	if bank[i] >= highestJoltage {
	// 		highestJoltage = bank[i]
	// 		highestJoltageIndex = i
	// 	}
	// }
	// secondHighestJoltage := 0
	// for i := len(bank) - 1; i > highestJoltageIndex; i-- {
	// 	if bank[i] > secondHighestJoltage {
	// 		secondHighestJoltage = bank[i]
	// 	}
	// }
	// return highestJoltage*10 + secondHighestJoltage
}

func getTotalOutputJoltage(banks [][]int, numberOfDigitsToTake int) int {
	sum := 0
	for _, bank := range banks {
		sum += getBankJoltage(bank, numberOfDigitsToTake)
	}
	return sum
}

func main() {
	banks, err := readInputFile()
	if err != nil {
		panic(err)
	}

	fmt.Println(getTotalOutputJoltage(banks, 2))
	fmt.Println(getTotalOutputJoltage(banks, 12))
}
