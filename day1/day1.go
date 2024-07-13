package day1

import (
	"bufio"
	"fmt"
	"os"
)

func ResolveDay1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic("Error while opening input file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var numbers []int
		var numToSum int
		token := scanner.Text()
		for _, char := range token {
			if char >= 48 && char <= 57 {
				numbers = append(numbers, int(char)-48)
			}
		}
		if len(numbers) == 0 {
			numToSum = 0
		} else if len(numbers) == 1 {
			numToSum = concatNumbers(numbers[0], numbers[0])
		} else {
			numToSum = concatNumbers(numbers[0], numbers[len(numbers)-1])
		}
		sum = sum + numToSum
	}
	fmt.Println(sum)
}

func concatNumbers(num1, num2 int) int {
	return num1*10 + num2
}
