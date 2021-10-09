package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	k := 9
	n := BinarySearch(arr,k)
	fmt.Printf("要查找的数字[%d]在位置[%d]",k,n)
}

func BinarySearch(s []int,k int) int {
	low,high := 0,len(s)-1
	for low < high {
		mid := (low + high) / 2
		if s[mid] < k {
			low = mid + 1
		} else if s[mid] > k {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}