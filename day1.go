package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./day1-input.txt")
	if err != nil {
		panic("Error while opening input file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		token := scanner.Text()
		for _, char := range token {
			if char >= 48 && char <= 57 {

			}
		}
	}
	fmt.Println(sum)
}
