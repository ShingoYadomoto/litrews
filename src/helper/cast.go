package helper

import (
	"strconv"
)

func AtoiSlice(stringSlice []string) (intSlice []int, err error) {
	intSlice = make([]int, len(stringSlice))

	for i, v := range stringSlice {
		var value int
		value, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		intSlice[i] = value
	}
	return
}
