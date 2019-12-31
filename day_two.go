package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// var staticCodes = []string{"1", "9", "10", "3", "2", "3", "11", "0", "99", "30", "40", "50"}

// OperatingSystem is the system for our rocket ship
func OperatingSystem() []string {
	fileContents, err := ioutil.ReadFile("./txt_inputs/day_two_input.txt")
	if err != nil {
		log.Fatal("Something went wrong reading the file", err.Error())
	}
	codes := strings.Split(string(fileContents), ",")
	// codes := staticCodes
	for i := 0; i < len(codes); i += 4 {
		ele := codes[i]
		if ele == "99" {
			return codes
		}
		posOne, posTwo, finalPos := getNextThreeIndices(i, codes)
		if ele == "1" {
			codes = opcode(posOne, posTwo, finalPos, codes, 1)
		} else if ele == "2" {
			codes = opcode(posOne, posTwo, finalPos, codes, 2)
		}
	}
	return codes
}

func opcode(firstCode, secondCode, position int, codes []string, code int) []string {
	firstVal, _ := strconv.Atoi(codes[firstCode])
	secondVal, _ := strconv.Atoi(codes[secondCode])
	if code == 1 {
		codes[position] = strconv.Itoa(firstVal + secondVal)
	} else {
		codes[position] = strconv.Itoa(firstVal * secondVal)
	}
	return codes
}

func getNextThreeIndices(idx int, codes []string) (int, int, int) {
	posOne, _ := strconv.Atoi(codes[idx+1])
	posTwo, _ := strconv.Atoi(codes[idx+2])
	posThree, _ := strconv.Atoi(codes[idx+3])
	return posOne, posTwo, posThree
}

// GetNounAndVerb is part 2 of the puzzle
func GetNounAndVerb() {
	fileContents, err := ioutil.ReadFile("./txt_inputs/day_two_input.txt")
	if err != nil {
		log.Fatal("Something went wrong reading the file", err.Error())
	}
	codes := strings.Split(string(fileContents), ",")
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			codeCopies := make([]string, len(codes))
			copy(codeCopies, codes)
			codeCopies[1] = strconv.Itoa(noun)
			codeCopies[2] = strconv.Itoa(verb)
			for i := 0; i < len(codeCopies); i += 4 {
				ele := codeCopies[i]
				if ele == "99" {
					if codeCopies[0] == "19690720" {
						fmt.Println("this is the noun and verb", noun, verb)
					}
					break
				}
				posOne, posTwo, finalPos := getNextThreeIndices(i, codeCopies)
				if ele == "1" {
					codeCopies = opcode(posOne, posTwo, finalPos, codeCopies, 1)
				} else if ele == "2" {
					codeCopies = opcode(posOne, posTwo, finalPos, codeCopies, 2)
				}
			}
		}
	}
}
