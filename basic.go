package main

import (
	"fmt"
	"math"
)
func main() {
	fmt.Println("hello world!");
	variableZeroValue();
	variableInitValue();
	variableInfer();
	fmt.Println(mySum(2,4));
	enums();
	triangle();
	fmt.Println("variables outside function a & b: ", a, b);
}

func variableZeroValue() {
	var a int;
	var str string;
	fmt.Println(a, str);
}

func variableInitValue() {
	var a int = 100;
	var str string = "ucsd";
	fmt.Println(a, str);
}

func variableInfer() {
	a, b := 10, 20
	b = 1000;
	fmt.Println(a, b)
}

func mySum(a int, b int) int {
	return a + b
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func triangle() {
	a, b := 3, 4;
	c := int(math.Sqrt(float64(a*a + b*b)));
	fmt.Println("the third line is: ", c);
}

var a, b int = 25, 25;
