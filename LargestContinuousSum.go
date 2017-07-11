// Largest continuous sum of an array of integers
// Author: Abhinaya Ramachandran
package main

import "fmt"
	
	
func LargestContinuousSum(arr []int) int {

	// Check for the edge case
	if len(arr) == 0{
		return 0
	}
	
	// Let the largest and current sum be
	// the first element of the array
	largestSum,currentSum := arr[0], arr[0]
	
	for i := 1;i < len(arr);i++   {
	
		// For each number that follows 
		// check if its addition makes currentSum greater than itself
		if currentSum + arr[i] > arr[i] { currentSum = currentSum + arr[i] } else {currentSum = arr[i]}
		
		// Update largestSum if the currentSum is greater than it
		if currentSum > largestSum { largestSum = currentSum} 
	}
	return largestSum
}

func main() {

	fmt.Println("The Largest Continuous Sum of ")
	fmt.Println("{1,2,-1} is ",LargestContinuousSum([]int{1,2,-1}))
	fmt.Println("{1,2,3,-7,-2} is ",LargestContinuousSum([]int{1,2,3,-7,-2}))
	fmt.Println("{1,2,1,-7,-3,17} is ",LargestContinuousSum([]int{1,2,1,-7,-3,17}))
	fmt.Println("{-2,-1,1,2,-1,3,-1,0} is ",LargestContinuousSum([]int{-2,-1,1,2,-1,3,-1,0}))
	fmt.Println("{-1,-2,-5,-4,-5} is ",LargestContinuousSum([]int{-1,-2,-5,-4,-5}))
	fmt.Println("{} is ",LargestContinuousSum([]int{}))
	fmt.Println("{1} is ",LargestContinuousSum([]int{1}))
	
}
