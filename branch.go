package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
	var res1 = calculator(10, 20, "+")
	fmt.Println(res1);
	fmt.Println(
		scoreToLetter(100),
		scoreToLetter(98),
		scoreToLetter(88),
		scoreToLetter(58),
		//scoreToLetter(101),
	)
}

func calculator(a, b int, op string) int {
	var res int
	switch op {
	case "+" : res = a + b;
	case "-" : res = a - b;
	case "*" : res = a * b;
	case "/" : res = int((float64(a)/float64(b)));
	default:
		panic("unsupported operator: " + op);
	}
	return res;
}

func scoreToLetter(score int) string {
	var letterGrade string;
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score));
	case score < 60: letterGrade = "F";
	case score < 80: letterGrade = "C";
	case score < 90: letterGrade = "B";
	case score <= 100: letterGrade = "A";
	}
	return letterGrade;
}