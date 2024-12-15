package main

import (
	"bufio"
	"day1/comparisons"
	"day1/inputs"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nums1, nums2, err := inputs.UserInputsToMaps(*scanner)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	score, err := comparisons.GetSimilarityScore(nums1, nums2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(score)
}
