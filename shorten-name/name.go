package main

import (
	"fmt"
	"strings"
)

func main() {
	fullNameArray := []string{"Nibras Ahmed", "Ahmed Saleem", "Ahsan Yousaf", "Muhammad Husnain"}
	shortNameArray := []string{}
	for _, name := range fullNameArray {
		parts := strings.Split(name, " ")
		first, last := parts[0], parts[len(parts)-1]
		shortNames := string(first[0]) + string(last[0])
		shortNameArray = append(shortNameArray, shortNames)
	}
	fmt.Println(shortNameArray)
}