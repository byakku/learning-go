package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testString := "aabbccc"
	findDuplicates(testString)
}

func findDuplicates(s1 string) int {

	elements := strings.Split(s1, "")

	duplicateFrequency := make(map[string]int)
  
	for _, value := range elements {
		_, exists := duplicateFrequency[value]

		if exists {
			duplicateFrequency[value] += 1
		} else {
			duplicateFrequency[value] = 0
		}

	}
	fmt.Println("Printing map\n", duplicateFrequency)
	duplicates := []string{}

	for _, value := range duplicateFrequency {
		if value == 1 {
		} else {
			duplicates = append(duplicates, strconv.Itoa(value))
		}
	}
	fmt.Println(duplicates)
	return 1 

}


