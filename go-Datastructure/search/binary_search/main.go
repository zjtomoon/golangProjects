package main

import "fmt"

func main() {
	//arr := []int{1,2,3,4,5,6,7,8,9,10}
	arr := []int{1, 2, 7, 9, 3, 10, 8, 5, 4, 6}
	QuickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
	k := 9
	n := BinarySearch(arr, k)
	fmt.Printf("要查找的数字[%d]在位置[%d]", k, n)
}

func BinarySearch(s []int, k int) int {
	low, high := 0, len(s)-1
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

func QuickSort(left int, right int, arr []int) {
	l, r := left, right
	privot := arr[(left+right)/2]
	temp := 0

	for l < r {
		for arr[l] < privot {
			l++
		}
		for arr[r] > privot {
			r--
		}
		if l >= r {
			break
		}

		//交换
		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp

		if arr[l] == privot {
			r--
		}
		if arr[r] == privot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}
