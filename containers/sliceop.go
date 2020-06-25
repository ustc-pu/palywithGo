package main
import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("create slice:")
	var s []int
	fmt.Println(s)
	for i := 0; i < 5; i++ {
		s = append(s, 2*i + 1)
	}
	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copy slices:")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("delet from slice:")
	s3 = append(s2[:3], s2[4:]...)
	printSlice(s3)

	fmt.Println("pop from front:")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("pop from back:")
	back := s2[len(s2) - 1]
	s2 = s2[:len(s2)-1]
	fmt.Println(back)
	printSlice(s2)
}
