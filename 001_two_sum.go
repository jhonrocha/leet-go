package main

import "fmt"

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for idx, value := range nums {
		if compl, ok := store[target-value]; ok {
			return []int{compl, idx}
		} else {
			store[value] = idx
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      //Output: [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         //Output: [0,1]
}
