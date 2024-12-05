package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	firstArrayNumbers, secondArrayNumbers := userInputToArray(*scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if len(firstArrayNumbers) != len(secondArrayNumbers) {
		return
	}
	sort.Ints(firstArrayNumbers)
	sort.Ints(secondArrayNumbers)
	diffs := 0

	for i := range firstArrayNumbers {
		diffs += int(math.Abs(float64(firstArrayNumbers[i] - secondArrayNumbers[i])))
	}

	fmt.Println(diffs)
}

func userInputToArray(scan bufio.Scanner) ([]int, []int) {

	var array1 []int
	var array2 []int
	for scan.Scan() {
		line := scan.Text()

		parts := strings.Fields(line)
		num1, err := strconv.Atoi(parts[0])
		if err == nil {
			array1 = append(array1, num1)
		}
		num2, err := strconv.Atoi(parts[1])
		if err == nil {
			array2 = append(array2, num2)
		}
	}
	return array1, array2
}
