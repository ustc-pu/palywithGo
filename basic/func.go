package main
import(
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func calculator(a, b int, op string) (int, error) {
	switch(op) {
	case "+": return a+b, nil
	case "-": return a-b, nil
	case "*": return a*b, nil
	case "/":
		res, _ := div(a, b)
		return res, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s\n", op)
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b;
}

/*
the arguments of this function are:
1 func named op,
2 a of type int, and 3 b of type int
-- op func(a, b int) int,
means op function has two arguments arg1 and arg2
return type is int
--
call op func to get the return value
 */
func apply(op func(arg1, arg2 int) int, a, b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function with multi args: %s" +
		"(%d, %d)\n", opName, a, b)
	fmt.Print("the result is: ")
	return op(a, b)
}

func myPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int { // various num of arguments
	res := 0;
	for i := range numbers {
		res += numbers[i]
	}
	return res
}

func swapbyPointer(a, b *int) {
	*a, *b = *b, *a
}

func swapbyReturn(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(div(13,3))
	if res, err := calculator(16, 2, "?"); err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(res)
	}
	res := apply(myPow, 5, 3)
	fmt.Println(res)

	res2 := apply(
		func(a, b int) int { //anonymous function
			return int(math.Pow(float64(a), float64(b)))
	}, 5, 3)
	fmt.Println(res2)
	fmt.Println(sum(1,3,5,7,9))

	x, y := 3, 4
	fmt.Printf("before swap, the values are: %d, %d\n", x, y)
	swapbyPointer(&x, &y)
	fmt.Printf("after swap, the values are: %d, %d\n", x, y)
	//fmt.Println(x, y)

	a, b := 7, 9
	fmt.Printf("before swap, the values are: %d, %d\n", a, b)
	a, b = swapbyReturn(a, b)
	fmt.Printf("after swap, the values are: %d, %d\n", a, b)
	//fmt.Println(a, b)
}