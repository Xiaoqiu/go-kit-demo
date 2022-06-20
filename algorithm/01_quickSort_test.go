package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
