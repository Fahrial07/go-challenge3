package main

import "fmt"

func main() {
	input := "selamat malam"

	//output setiap karakter
	for _, r := range input {
		fmt.Printf("%c\n", r)
	}

	charactrCoun := make(map[string]int)

	for _, r := range input {
		charactrCoun[string(r)] += 1
	}

	fmt.Println(charactrCoun)
}
