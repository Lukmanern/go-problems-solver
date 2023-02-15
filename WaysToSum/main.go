package main

import "fmt"

func main() {
	fmt.Println(wayToSum(8, 2))
	fmt.Println(NumberOfways(8, 2))
	fmt.Println(wayToSum(10, 3))
	fmt.Println(NumberOfways(10, 3))
	fmt.Println(wayToSum(5, 2))
	fmt.Println(NumberOfways(5, 2))
	fmt.Println(wayToSum(12, 3))
	fmt.Println(NumberOfways(12, 3))
	fmt.Println(wayToSum(29, 7))
	fmt.Println(NumberOfways(29, 7))
}

// my version
func wayToSum(N int, K int) int {
	if N == 0 {
		return 1
	}
	if N < 0 {
		return 0
	}
	count := 0
	for i := 1; i <= K; i++ {
		count += wayToSum(N-i, i)
	}
	return count
}

// geeksforgeeks version
// see https://www.geeksforgeeks.org/ways-to-sum-to-n-using-natural-numbers-up-to-k-with-repetitions-allowed/
func NumberOfways(N, K int) int {

    // Initialize a slice
    dp := make([]int, N+1)
    
    // Update dp[0] to 1
    dp[0] = 1
    
    // Iterate over the range [1, K+1]
    for row := 1; row < K+1; row++ {
    
        // Iterate over the range [1, N+1]
        for col := 1; col < N+1; col++ {
        
            // If col is greater than or equal to row
            if col >= row {
            
                // Update current dp[col] state
                dp[col] = dp[col] + dp[col-row]
            }
        }
    }
    
    // Return the total number of ways
    return dp[N]
}