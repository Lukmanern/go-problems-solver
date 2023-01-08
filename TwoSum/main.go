package main

import "fmt"

func main() {
	nums := []int{
		-1,0,1,4,6,7,8,
	}
	target := 9
	fmt.Println(twoSum(nums, target)) // (2,6) => 1 + 8 = 9
}

func twoSum(nums []int, target int) []int {
	// Initialize variables to store 
	// the indices of the elements
	var i, j int
	// Initialize a flag to indicate 
	// when to break out of the loop
	var isBreak bool = false
	// Get the length of the array
	var ln int = len(nums)

	// Check if the length of the array is 2
	if ln == 2 {
		// Return the indices if 
		// the length of the array is 2
		return []int{0, 1}
	}

	// Loop over the array elements
	for i = 0; i < ln; i++ {
		// Loop over the remaining elements
		for j = i + 1; j < ln; j++ {
			// Check if the current element and 
			// the next element add up to the target
			if nums[i]+nums[j] == target {
				// Set the flag to true if 
				// the target has been found
				isBreak = true
				break
			}
		}
		// Break out of the outer loop 
		// if the target has been found
		if isBreak {
			break
		}
	}

	// Return the indices of the 
	// elements that add up to the target
	return []int{i, j}
}