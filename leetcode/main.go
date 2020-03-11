package main

import (
	"fmt"

	"github.com/TCLP/golang-dev/leetcode/bit"
	"github.com/TCLP/golang-dev/leetcode/sort"
)

func bitReverseRun() {
	fmt.Println(bit.BitReverse(10))
}

func bitSingleNumberRun() {
	var nums []int = []int{1, 3, 2, 4, 1, 2, 3}
	fmt.Println(bit.SingleNumber(nums))
}

func bitHangmingWeightRun() {
	fmt.Println(bit.HangmingWeight(3))
}

func sortQuickSortRun() {
	var arr []int = []int{5, 3, 1, 4, 2}
	sort.QuickSort(arr, 0, 4)
	fmt.Println(arr)
}

func main() {
	fmt.Println("this is leetcode")

	bitReverseRun()
	bitSingleNumberRun()
	bitHangmingWeightRun()
	sortQuickSortRun()
}
