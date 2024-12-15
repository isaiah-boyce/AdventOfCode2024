package inputs

import (
	"bufio"
	"strconv"
	"strings"
)

func UserInputsToMaps(scan bufio.Scanner) (map[int]int, map[int]int, error) {

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
