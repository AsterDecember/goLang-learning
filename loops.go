package main

import "fmt"

func main () {

	/*for i:=  1; i<=10; i++ {
		fmt.Println(i)
	}

	x:=1
	for x<=10{
		fmt.Println(x)
		x++
	}

	for i:=1; i<5; i++{
		for j:=1; j<i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}*/

	EvenNum := [5] int {0,2,4,6,8}

	//index can be also _ (underscore) if you dont want to use the index of array
	for index, value := range EvenNum {
		fmt.Printf("value: %d",value)
		fmt.Printf(" index: %d\n",index)
	}

}