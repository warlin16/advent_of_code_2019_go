package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func getCodesFromFile() []string {
	fileContents, err := ioutil.ReadFile("./txt_inputs/day_two_input.txt")
	if err != nil {
		log.Fatal("Something went wrong reading the file", err.Error())
	}
	codes := strings.Split(string(fileContents), ",")
	return codes
}

func intCodeParser(codes []string) []string {
	for i := 0; i < len(codes); i += 4 {
		ele := codes[i]
		if ele == "99" {
			break
		}
		posOne, _ := strconv.Atoi(codes[i+1])
		posTwo, _ := strconv.Atoi(codes[i+2])
		finalPos, _ := strconv.Atoi(codes[i+3])
		firstVal, _ := strconv.Atoi(codes[posOne])
		secondVal, _ := strconv.Atoi(codes[posTwo])
		if ele == "1" {
			codes[finalPos] = strconv.Itoa(firstVal + secondVal)
		} else if ele == "2" {
			codes[finalPos] = strconv.Itoa(firstVal * secondVal)
		}
	}
	return codes
}

// OperatingSystem is the system for our rocket ship part 1 of the puzzle
func OperatingSystem() []string {
	codes := getCodesFromFile()
	codes[1] = "12"
	codes[2] = "2"
	return intCodeParser(codes)
}

// GetNounAndVerb is part 2 of the puzzle
func GetNounAndVerb() {
	codes := getCodesFromFile()
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			codeCopies := make([]string, len(codes))
			copy(codeCopies, codes)
			codeCopies[1] = strconv.Itoa(noun)
			codeCopies[2] = strconv.Itoa(verb)
			codeCopies = intCodeParser(codeCopies)
			if codeCopies[0] == "19690720" {
				fmt.Println("the noun and verb are:", noun, verb)
			}
		}
	}
}
