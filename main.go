package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
	// a := 5
	// double(&a)
	// fmt.Println(a)
}
func reverse(arr *[4]int) {
	for index, value := range *arr {
		arr[len(arr)-1-index] = value
	}

}

// func double(num *int) {
// 	*num = *num * 2
// }
