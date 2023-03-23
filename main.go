package main

import "fmt"

func main() {
	kata := "selamat malam"
	hit := make(map[rune]int)
	for _, huruf := range kata {
		fmt.Println(string(huruf))
		hit[huruf]++
	}
	for k, v := range hit {
		fmt.Print(string(k), ": ", v, " ")
	}
}
