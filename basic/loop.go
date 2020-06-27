package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		intToBin(5),
		intToBin(32),
		)
	printFile("abc.txt");
	//forever()
}

func intToBin(num int) string {
	var res string;
	for ; num > 0; num /= 2 { //omit the init value
		var low int = num % 2;
		res = strconv.Itoa(low) + res;
	}
	return res;
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		fmt.Println(scanner.Text());
	}
}

//func forever() {
//	for {
//		fmt.Println("aaa")
//	}
//}