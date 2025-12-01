package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Rotation struct {
	Direction string
	Value     int
}

func readInputFile() ([]Rotation, error) {
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

	var rotations []Rotation
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return nil, err
		}

		rotations = append(rotations, Rotation{
			string(line[:1]),
			value,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, err
}

func rotate(cursor int, rotation Rotation) int {
	if rotation.Direction == "L" {
		return cursor - rotation.Value
	}
	if rotation.Direction == "R" {
		return cursor + rotation.Value
	}
	return cursor
}

func part1(rotations []Rotation) int {
	cursor := 50
	numberOfZeros := 0

	for _, rotation := range rotations {
		cursor = rotate(cursor, rotation)
		// handle overflow / underflow
		for cursor < 0 || cursor > 99 {
			if cursor < 0 {
				cursor += 100
			}
			if cursor > 99 {
				cursor -= 100
			}
		}

		if cursor == 0 {
			numberOfZeros++
		}
	}

	return numberOfZeros
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(rotations []Rotation) int {
	cursor := 50
	numberOfZeros := 0

	for _, rotation := range rotations {
		was_greater_than_0 := cursor > 0
		cursor = rotate(cursor, rotation)

		newZeros := 1
		if cursor != 0 {
			newZeros = Abs(cursor / 100)
		}
		if was_greater_than_0 && cursor < 0 {
			newZeros++
		}

		numberOfZeros += newZeros
		cursor %= 100

		if cursor < 0 {
			cursor += 100
		}
	}

	return numberOfZeros
}

func main() {
	rotations, err := readInputFile()
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(rotations))
	fmt.Println(part2(rotations))
}
