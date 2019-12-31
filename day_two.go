package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var staticCodes = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

// OperatingSystem is the system for our rocket ship
func OperatingSystem() {
	fileContents, err := ioutil.ReadFile("./txt_inputs/day_two_input.txt")
	if err != nil {
		log.Fatal("Something went wrong reading the file", err.Error())
	}
	codes := strings.Split(string(fileContents), ",")
	for _, ele := range codes {
		fmt.Println("this is the current ele", ele)
		// currentCode := strconv.Atoi(ele)
		if ele == "1" {
			// do opcode 1
			fmt.Println("to be determined")
		} else if ele == "2" {
			// do opcode 2
			fmt.Println("to be determined")
		}
	}
}

func opcodeOne(codes []int) []int {
	// firstPos, secondPos := codes[1], codes[2]
	// codes[3] = codes[firstPos] + codes[secondPos]
	// return codes
}

func opcodeTwo(codes []int) []int {
}
