package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cubeMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func ResolveDay2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		panic("Error while opening input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumIds := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplited := strings.Split(line, ":")
		gameIdArray := strings.Split(lineSplited[0], " ")
		gameId, err := strconv.Atoi(strings.Trim(gameIdArray[len(gameIdArray)-1], ":"))
		if err != nil {
			panic("Error while getting game ID")
		}
		subsets := strings.Split(lineSplited[1], ";")
		isInvalid := false
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
					if cubeNumValue > cubeMap["red"] {
						isInvalid = true
						break
					}
				} else if cubeType == "green" {
					if cubeNumValue > cubeMap["green"] {
						isInvalid = true
						break
					}
				} else if cubeType == "blue" {
					if cubeNumValue > cubeMap["blue"] {
						isInvalid = true
						break
					}
				}
			}
		}
		if !isInvalid {
			sumIds += gameId
		}
	}
	fmt.Println(sumIds)
}
