// function that greets the user by name, or by saying "Hello, World!" if no name is given.

package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the user:")
 	var text string
	text, _ = reader.ReadString('\n')
//	fmt.Printf("Hello, %v!", text)
	if len(text) == 1  {
//		fmt.Printf("Hello, %v", text)
		 fmt.Println("Hello World")

	}else{
		fmt.Printf("Hello, %v\n",text)
	}
} 
