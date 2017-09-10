package main

import (
	"fmt"
)

/*
* Given an integer array of positive numbers, 
* return the minimum number of jumps needed to hop the array.
* Each element of the given slice represents the number of legal hops possible.
* For example arr[2] = 2 means, from index 2, only two successive hops are possible.
*/


// Minimum number of jumps needed to hop an array using dynamic programming

func ArrayHopping(arr []int) {

	// The first element is 0, no hops can be performed
	
	if len(arr) == 0 || arr[0] <= 0 {
		fmt.Println("Array cannot be hopped")
		return 
	}


	/* jump maintains the number of hops needed to reach the i-th element.
	 * jump_from maintains the corresponding index from which the hop is made 
	*/
	
	jump := []int{}
	jump_from := make([]int , len(arr))
	jump = append(jump, 0)
	
	
	
	for i := 1 ; i < len(arr); i ++{
		for j := 0 ; j < i ; j ++{
			if i-j <= arr[j] {
				hop := jump[j] + 1
				
				/* If the previously calculated hop value is greater than the current,
				*  replace it with the new value, else ignore 
				*/
				
				// Assigning the value for the first time
				if  len(jump) < i + 1 {
					jump = append(jump, hop)
					jump_from[i] = j
				}else if jump[i] > hop{
				// Replacing the existing value
					jump[i] = hop
					jump_from[i] = j
				} 
				
			}
		}
	}
	
	
	
	if len(jump) == len(arr)  {
	
		fmt.Println("The minimum number of hops needed is " ,jump[len(arr)-1])
		fmt.Println("The sequence is")
		seq := []int{}
		ind := len(jump_from) -1 
		for ind > 0 {
			seq = append(seq , jump_from[ind])
			ind = jump_from[ind]
		}
		
		for i:= len(seq) - 1 ; i >= 0 ; i -- {
			fmt.Print("-->")
			fmt.Println(seq[i])
		}
		return 
	}else{
	
		fmt.Println("Array Cannot be Hopped")
		return
	}
	
	
}


func main() {

	ArrayHopping([]int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9})
	ArrayHopping([]int{2,3,1,1,2,4,2,0,0,1})
	ArrayHopping([]int{0,1,2,3})
	ArrayHopping([]int{1,1,1,0,2})
	ArrayHopping([]int{})
	
}
