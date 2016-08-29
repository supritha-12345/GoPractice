package main

import(
	 "fmt"
//	"os"
)

func factorial(x uint) uint {
	if x == 0{
		return 1
	}
	return x*factorial(x-1)
}

func main(){
	var i uint
	fmt.Println("Enter the number:")
	fmt.Scan(&i)
//	fmt.Println(i)
	var fact uint	
	fact = factorial(i)
	fmt.Println("the factorial of entered number is:")
	fmt.Printf("%v",fact)
}
