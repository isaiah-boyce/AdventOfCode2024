package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstArrayNumbers, secondArrayNumbers, err := userInputsToMaps(*scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	score := 0

	for k, v := range firstArrayNumbers {
		// we need to take the number from the first array,
		// look at the second array and find how many times that number occurs in the second array
		// then multiply the number from the first by the occurances in the second array.
		// then add that to diffs.
		score += k * secondArrayNumbers[k] * v
		//score += int(math.Abs(float64(firstArrayNumbers[i] - secondArrayNumbers[i])))
	}

	fmt.Println(score)
}

func userInputsToMaps(scan bufio.Scanner) (map[int]int, map[int]int, error) {

	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for scan.Scan() {
		line := scan.Text()

		parts := strings.Fields(line)
		num1, err := strconv.Atoi(parts[0])
		if err == nil {
			map1[num1] = map1[num1] + 1
		}
		num2, err := strconv.Atoi(parts[1])
		if err == nil {
			map2[num2] = map2[num2] + 1
		}
	}
	return map1, map2, nil
}
