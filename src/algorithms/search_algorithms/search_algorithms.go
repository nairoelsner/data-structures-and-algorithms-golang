package searchalgorithms

import (
	"errors"
)

func LinearSearch(value int, list []int) (int, error) {
	for i, v := range list {
		if value == v {
			return i, nil
		}
	}
	return -1, errors.New("Value not found!")
}

func BinarySearch(value int, list []int) (int, error) {
	return binarySearchRecusrsion(value, 0, len(list)-1, list)
}

func binarySearchRecusrsion(value int, start int, end int, list []int) (int, error) {
	if start > end {
		return -1, errors.New("Value not Found!")
	}

	median := (end + start) / 2

	if value == list[median] {
		return median, nil
	}

	if value > list[median] {
		return binarySearchRecusrsion(value, median+1, end, list)
	} else {
		return binarySearchRecusrsion(value, start, median-1, list)
	}
}
