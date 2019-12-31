package main

import (
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
		posOne, posTwo, finalPos := getNextThreeIndices(i, codes)
		if ele == "1" {
			codes = opcode(posOne, posTwo, finalPos, codes, 1)
			continue
		} else if ele == "2" {
			codes = opcode(posOne, posTwo, finalPos, codes, 2)
			continue
		} else if ele == "99" {
			return codes
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
