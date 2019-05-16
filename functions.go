package main
import "fmt"
func main() {
	x,y := 5,6

	fmt.Println(add(x,y))
	fmt.Println("---------------")

	fmt.Println(factorial(5))
}

func add(a,b int) int {
	return a+b
}

func factorial(num int)  int{
	if num == 0{
		return 1
	}
	return num * factorial(num-1)
}
