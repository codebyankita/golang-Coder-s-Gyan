package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	// var nums1 = make([]int,2,5)
	// var nums1 = make([]int,0,5)
    nums1 :=[]int{}
	fmt.Println(nums1)
	fmt.Println(cap(nums1))
	
	fmt.Println(nums1 == nil)
	nums1 =append(nums1, 1)
	// nums1 = append(nums1, 2)
	// nums1 = append(nums1, 3)
	// nums1 = append(nums1, 4)
	// nums1 = append(nums1, 5)
	// nums1 = append(nums1, 6)
	// nums1 = append(nums1, 7)
	// nums1 = append(nums1, 8)
	// nums1 = append(nums1, 9)
	// nums1 = append(nums1, 10)

	fmt.Println(nums1)
	fmt.Println(cap(nums1))

	var numsa = []int{1, 2, 3, 4, 5}
	fmt.Println(numsa[0:1])
	fmt.Println(numsa[:2])
	fmt.Println(numsa[1:])
	
 var nums2 = []int{1, 2, 3}
	var nums3 = []int{1, 2, 4}

	fmt.Println(slices.Equal(nums2, nums3))

	var nums4 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums4)

}

// package main

// slice -> dynamic
// most used construct in go
// + useful methods
// func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)

	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

// }