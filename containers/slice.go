package main

import (
	"fmt"
)

//------how to create a slice
//1) use ":"/colon to get a slice from an array
//2)


func updateSlice(s []int) {
	s[0] = -1
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s := arr[2:6]
	fmt.Println("arr[2:6] =", s)
	fmt.Println("arr[:6] =", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println("arr =", arr) //slice is passed by ref

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println("arr =", arr)

	fmt.Println("Reslice:")
	fmt.Println(s2) //[-1 1 -1 3 4 5 6 7]
	s2 = s2[:5]
	fmt.Println(s2) //[-1 1 -1 3 4]
	s2 = s2[2:]
	fmt.Println(s2) //[-1 3 4]

	fmt.Println("Extending slice:")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 10)
	fmt.Println("s3 =", s3, s4, s5)
	fmt.Println(arr, cap(arr))
}
