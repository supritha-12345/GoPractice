package main

import (
	"fmt"
	"os"
	"bufio"
//	"unicode/utf8"
)

func main(){

//read a file
	file, err := os.Open("/home/supritha/goboot/helloworlproj/src/vowels/sample.txt")
	if err!=nil {
		fmt.Println("error message:", err)
	}
	var lines[] string
	reader := bufio.NewScanner(file)
	for reader.Scan(){
		lines = append(lines,reader.Text())	
	}
	//fmt.println("Lines are:", lines)

//display a file
	for i,line := range lines {
		fmt.Println(i,line)
	}


	var count int
//	fmt.Printf("Length:%x",len(lines))
//	fmt.Printf("char count:%x",utf8.RuneCountInString(lines))

	 for _, c := range lines {
	   fmt.Println(c)
		for _, i := range c {
//			fmt.Println(string(i))
			if string(i) == string('a') {
				count++
			}
			if string(i) == string('e') {
                                count++
                        }
			if string(i) == string('i') {
                                count++
                        }
			if string(i) == string('o') {
                                count++
                        }
			if string(i) == string('u') {
                                count++
                        }
			if string(i) == string('A') {
                                count++
                        }
                        if string(i) == string('E') {
                                count++
                        }
                        if string(i) == string('I') {
                                count++
                        }
                        if string(i) == string('O') {
                                count++
                        }
                        if string(i) == string('U') {
                                count++
                        }

		}
	}
	fmt.Printf("The total no of vowels in the file text is: %v",count)

	//write the file lines into another file
	file1,err1 := os.Create("/home/supritha/goboot/helloworlproj/src/vowels/newfile.txt")
	if err1!=nil {
		fmt.Println("error while writing the new file:", err1)
	}
	writer := bufio.NewWriter(file1)
	
//	var text[] string
	for _,writeline := range lines {
		fmt.Fprintln(writer, writeline)
	}
	writer.Flush()

//close the file
	file.Close()
}
