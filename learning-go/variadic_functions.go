package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	arr := []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))
}
