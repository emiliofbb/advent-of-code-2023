package day1_2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numbersStrings = [][]string{
	{"2", "two"},
	{"3", "three"},
	{"4", "four"},
	{"5", "five"},
	{"6", "six"},
	{"7", "seven"},
	{"8", "eight"},
	{"9", "nine"},
}

func ResolveDay1_2() {
	file, err := os.Open("./day1-2/input.txt")
	if err != nil {
		panic("Error while opening input file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var numToSum int
		token := scanner.Text()
		firstIdx, lastIdx := searchFirstLastIndex(token, "1", "one")
		first, last := 1, 1
		var tmpFirst, tmpLast int
		var numActual int
		for idx, strsToSearch := range numbersStrings {
			numActual = idx + 2
			tmpFirst, tmpLast = searchFirstLastIndex(token, strsToSearch[0], strsToSearch[1])
			if tmpFirst != -1 && tmpFirst < firstIdx {
				firstIdx = tmpFirst
				first = numActual
			} else if tmpFirst != -1 && firstIdx == -1 {
				firstIdx = tmpFirst
				first = numActual
			}
			if tmpLast > lastIdx {
				lastIdx = tmpLast
				last = numActual
			}
		}
		if firstIdx == -1 {
			numToSum = 0
		} else {
			numToSum = concatNumbers(first, last)
		}
		sum = sum + numToSum
	}
	fmt.Println(sum)
}

func concatNumbers(num1, num2 int) int {
	return num1*10 + num2
}

func searchFirstLastIndex(str, numStr, numSpell string) (first, last int) {
	indexFirstNumStr := strings.Index(str, numStr)
	indexFirstNumSpell := strings.Index(str, numSpell)
	indexLastNumStr := strings.LastIndex(str, numStr)
	indexLastNumSpell := strings.LastIndex(str, numSpell)
	if indexFirstNumStr != -1 && indexFirstNumSpell == -1 {
		first = indexFirstNumStr
	} else if indexFirstNumStr == -1 && indexFirstNumSpell != -1 {
		first = indexFirstNumSpell
	} else if indexFirstNumStr == -1 && indexFirstNumSpell == -1 {
		first = -1
	} else if indexFirstNumStr < indexFirstNumSpell {
		first = indexFirstNumStr
	} else {
		first = indexFirstNumSpell
	}
	if indexLastNumStr > indexLastNumSpell {
		last = indexLastNumStr
	} else if indexLastNumSpell > indexLastNumStr {
		last = indexLastNumSpell
	} else {
		last = -1
	}
	return first, last
}
