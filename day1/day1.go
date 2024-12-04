package main

import (
	"bufio"
	"fmt"
	"log"
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
	numbers := userInputToArray(*scanner)
	println(numbers)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// fmt.Println("enter in numbers in array one")
	// firstArray := bufio.NewScanner(os.Stdin)
	// firstArrayNumbers := userInputToArray(*firstArray)

	// fmt.Println("enter in numbers in array two")
	// secondArray := bufio.NewScanner(os.Stdin)
	// secondArrayNumbers := userInputToArray(*secondArray)
	// var diffs []int
	// leng := len(firstArrayNumbers)

	// for i := 0; i < leng; i++ {
	// 	x := findIndexOfSmallestValue(firstArrayNumbers)
	// 	firstArrayNumbers[x] = math.MaxInt32
	// 	y := findIndexOfSmallestValue(secondArrayNumbers)
	// 	secondArrayNumbers[y] = math.MaxInt32
	// 	diffs = append(diffs, int(math.Abs(float64(x-y))))
	// }

	// fmt.Println(diffs)
}

func findIndexOfSmallestValue(firstArrayNumbers []int) int {
	smallest := firstArrayNumbers[0]
	index := 0
	for i, num := range firstArrayNumbers {
		if num < smallest {
			smallest = num
			index = i
		}
	}
	return index
}

func userInputToArray(scan bufio.Scanner) []int {

	var nums []int
	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		numsAsStr := strings.Split(scan.Text(), "   ")
		for _, num := range numsAsStr {
			i, err := strconv.Atoi(num)

			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, i)
		}
	}
	return nums
}
