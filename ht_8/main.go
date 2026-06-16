package main

import (
	"errors"
	"fmt"
)

// Hard
// Реализация бинарного поиска. Возвращаем индекс искомого элемента в переданном массиве. Если элемента нет, то возвращаем -1.
func BinarySearch(lst []int, target int) (int, error) {
	var n int
	left, right := 0, len(lst)-1
	for left <= right {
		n = (left + right) / 2
		if lst[n] == target {
			return n, nil
		}
		if lst[n] < target {
			left = n + 1
		} else {
			right = n - 1
		}

	}
	return -1, errors.New("Указанного элемента нет в списке.")
}

func main() {
	lst := []int{1, 2, 5, 7, 8, 9, 11, 24, 56, 78, 102}

	idx, err := BinarySearch(lst, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Индекс элемента: %d\n", idx)
	}

	idx, err = BinarySearch(lst, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Индекс элемента: %d\n", idx)
	}
}
