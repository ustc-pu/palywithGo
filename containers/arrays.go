package main
import (
	"fmt"
)

func printArray(arr *[5]int) {
	arr[0] = -1
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr1, arr2, arr3)

	var grid [3][4]int //3*4 matrix
	fmt.Println(grid)

	//how to walk through an array
	//1) the common way in any language
	for i:=0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//2) range key word
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	//array arg pass by value, a new copy
	//pass a ptr if you want to change the value
	num := [...]int{0,1,2,3,4}
	printArray(&num)
	fmt.Println(num)

}
