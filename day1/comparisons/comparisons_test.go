package comparisons

import (
	"testing"
)

func TestGetSimilarityScore(t *testing.T) {
	nums1 := map[int]int{1: 1, 2: 1, 3: 1}
	nums2 := map[int]int{1: 1, 2: 1, 3: 1}
	expected := 6
	score, err := GetSimilarityScore(nums1, nums2)
	if err != nil {
		t.Fatalf("wasnt able to get score")
	}
	if score != expected {
		t.Fatalf(`expected %d to be %d, got %d`, score, expected, score)
	}
}

func TestEmptyMaps(t *testing.T) {
	nums1 := map[int]int{1: 1, 2: 1, 3: 1}
	nums2 := map[int]int{}
	score, err := GetSimilarityScore(nums1, nums2)
	if err == nil || score != 0 {
		t.Fatalf("this should error")
	}
}
