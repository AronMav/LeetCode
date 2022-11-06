package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV")) //1994
}

func romanToInt(s string) int {

	template := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	previous := 0
	sum := 0
	for _, v := range s {
		nember := template[string(v)]
		if previous < nember {
			sum += nember - previous*2
		} else {
			sum += nember
		}
		previous = nember
	}
	return sum
}
