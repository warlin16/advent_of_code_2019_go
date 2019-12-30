package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// SumOfFuelRequirements gets the sum of fuel requirements
func SumOfFuelRequirements() (int, error) {
	fileContents, err := ioutil.ReadFile("day_one_input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return 0, err
	}
	modules := strings.Split(string(fileContents), "\n")
	sum := 0
	for _, ele := range modules {
		module, _ := strconv.Atoi(ele)
		sum += getModuleMass(module)
	}
	return sum, nil
}

func getModuleMass(module int) int {
	roundedMass := math.Floor(float64(module / 3))
	return int(roundedMass) - 2
}
