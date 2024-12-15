package comparisons

import "errors"

func GetSimilarityScore(nums1 map[int]int, nums2 map[int]int) (int, error) {
	score := 0
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0, errors.New("Insufficient numbers to get similarity score")
	}
	for k, v := range nums1 {
		score += k * nums2[k] * v
	}
	return score, nil
}
