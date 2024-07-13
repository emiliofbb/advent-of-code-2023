package day1

import (
	"bufio"
	"fmt"
	"os"
)

func Resolve_day1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic("Error while opening input file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var numeros []int
		var num_sumar int
		token := scanner.Text()
		for _, char := range token {
			if char >= 48 && char <= 57 {
				numeros = append(numeros, int(char)-48)
			}
		}
		if len(numeros) == 0 {
			num_sumar = 0
		} else if len(numeros) == 1 {
			num_sumar = concat_numbers(numeros[0], numeros[0])
		} else {
			num_sumar = concat_numbers(numeros[0], numeros[len(numeros)-1])
		}
		sum = sum + num_sumar
	}
	fmt.Println(sum)
}

func concat_numbers(num1, num2 int) int {
	return num1*10 + num2
}
