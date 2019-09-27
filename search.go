package main

import "fmt"

func search(arr []int, val int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			fmt.Println(val)
		} else if val > 3 {
			fmt.Println("That value is not Correct")
			break
		}

	}
}

func main() {
	arr := []int{1, 2, 3}
	search(arr, 2)
}
