/**
Bill Gates is on one of his philanthropic journeys to a village in Utopia.
He has N packets of candies and would like to distribute one packet to each of
the K children in the village (each packet may contain different number of
candies). To avoid any fighting among the children, he would like to pick
K out of N packets, such that unfairness is minimized.

Suppose the K packets have (x1, x2, x3,....xk) candies in them, where xi
denotes the number of candies in the ith packet, then we define unfairness as

max(x1,x2,...xk) - min(x1,x2,...xk)

where max denotes the highest value amongst the elements, and min denotes
the least value amongst the elements. Can you figure out the minimum
unfairness and print it?
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
			return
		}
	}

}
