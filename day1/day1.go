package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums1, nums2, err := userInputsToMaps(*scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	score, err := getSimilarityScore(nums1, nums2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(score)
}

func getSimilarityScore(nums1 map[int]int, nums2 map[int]int) (int, error) {
	score := 0
	if nums1 == nil {
		return 0, errors.New("no numbers provided")
	}
	for k, v := range nums1 {
		score += k * nums2[k] * v
	}
	return score, nil
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
