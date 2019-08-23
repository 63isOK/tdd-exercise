package peaks

import (
	"fmt"
)

var (
	upPos = 0
)

// PrintPeaks print pos and value of peaks
func PrintPeaks(data []int) string {
	value := []int{}
	pos := findPeaks(data)
	for _, x := range pos {
		value = append(value, data[x])
	}

	return fmt.Sprintf("pos %v peaks %v", pos, value)
}

// findPeaks get pos of peaks
func findPeaks(data []int) []int {
	result := make([]int, 0)

	for i := range data {
		if i == 0 {
			continue
		}

		if data[i] > data[i-1] {
			upPos = i
		} else if data[i] < data[i-1] {
			if upPos != 0 {
				result = append(result, upPos)
				upPos = 0
			}
		}
	}

	return result
}
