/**
Minimum Possible Unfairness Value
*/

package main

import (
	"fmt"
	"sort"
)

const (
	max        = 100000
	maxVal     = 1000000001
	unfairness = 999999999999
)

func main() {

	var n, k, min, max, newUnfairness int

	fmt.Println("Enter number of packets")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	values := make([]int, n)

	fmt.Println("Enter number of childern")
	_, err = fmt.Scanf("%d", &k)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Enter how many candies for %d \n", i+1)
		fmt.Scanf("%d", &values[i])
	}

	fmt.Println("Printing the set: ", values)
	sort.Ints(values)
	fmt.Println("Sorted set :", values)

	for j := 0; j < n-k; j++ {
		z := k + j - 1
		if z > n-1 {
			break
		}
		min = values[j]
		max = values[z]
		newUnfairness = max - min
		if newUnfairness < unfairness {
			fmt.Println("Minimum Possible Unfairness Value ", unfairness)
		}
	}

}
