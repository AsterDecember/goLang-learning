package main
import "fmt"

func main ()  {
	//EvenNum := [5]int{0,2,4,6,8}
	numSlice := []int{5,4,3,2,1}
	sliced := numSlice[3:5]

	for i, value := range numSlice {
		fmt.Println(value,i)
	}


	fmt.Println(sliced)

	//here comes the array dynamically defined
	/**
	*args: type of array, size, capacity (size cant be greater than capacity)
	 */
	slice2 := make([]int,0,10)
	fmt.Println(slice2)
	fmt.Printf("%s \n", "------------ ")
	copy(slice2,numSlice)
	fmt.Println(slice2)
	fmt.Println(numSlice)
	//append
	slice3 := append(numSlice, 3,0,-1)
	fmt.Println(slice3)
}