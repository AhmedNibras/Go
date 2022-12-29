package main

import (
	"fmt"
)

func main() {
	// initialize fullNameArray and shortNameArray
	fn := []string{"Nibras Ahmed", "Ahmed Saleem", "Ahsan Yousaf", "Muhammad Husnain"}
	sn := []string{}

	for _, n := range fn {
		// initialize space
		s := 0
		// loop through the name
		for i, c := range n {
			// if space is found, break the loop
			if c == ' ' {
				s = i
				break
			}
		}
		// get the short name		
		shortNames := string(n[0]) + string(n[s+1])
		// append the short name to the shortNameArray
		sn = append(sn, shortNames)
	}

	// print the shortNameArray
	fmt.Println(sn)
}
