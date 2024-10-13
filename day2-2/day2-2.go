package day2_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ResolveDay2_2() {
	file, err := os.Open("./day2-2/input.txt")
	if err != nil {
		panic("Error while opening input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineValues := strings.Split(line, ":")[1]
		subsets := strings.Split(lineValues, ";")
		maxRedValue := 0
		maxGreenValue := 0
		maxBlueValue := 0
		for _, subset := range subsets {
			subset = strings.Trim(subset, " ")
			cubes := strings.Split(subset, " ")
			for i := 0; i < len(cubes); i = i + 2 {
				cubeNum := cubes[i]
				cubeType := strings.Trim(cubes[i+1], ",")
				cubeNumValue, err := strconv.Atoi(cubeNum)
				if err != nil {
					panic("Error while translate string to int")
				}
				if cubeType == "red" {
					if cubeNumValue > maxRedValue {
						maxRedValue = cubeNumValue
					}
				} else if cubeType == "green" {
					if cubeNumValue > maxGreenValue {
						maxGreenValue = cubeNumValue
					}
				} else if cubeType == "blue" {
					if cubeNumValue > maxBlueValue {
						maxBlueValue = cubeNumValue
					}
				}
			}
		}
		sumPower += maxRedValue * maxGreenValue * maxBlueValue
	}
	fmt.Println(sumPower)
}
