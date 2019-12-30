package main

import (
	"fmt"
	"io/ioutil"
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
	return module/3 - 2
}
