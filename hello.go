package main

import "fmt"

func main () {
	fmt.Printf("hello, world\n")
	var name string = "Miguel Gomez"
    const pi float64 = 3.14151612
    win := true

	fmt.Println(len(name))
	fmt.Println(name + "Ml")
    fmt.Println(win)

	fmt.Printf("%.3f \n",pi)
	fmt.Printf("%T \n", name)
	fmt.Printf("%t \n", win)
	//print binaries
	fmt.Printf("%b \n", 25)
	//prints character to asc code
	fmt.Printf("%c \n", 35)
	//prints hex code
    fmt.Printf("%x \n", 50)

    fmt.Printf("%e \n", pi)

}