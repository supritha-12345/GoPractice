package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum1(s string, c chan int) {
//	sum := 0
	go func(){
	for i:=1; i<=10;i++ {
		c <- i	
		}
//	close(c)
//	wg.Done()
	}()
	wg.Done()
	//c <- sum // send sum to c
}


func sum2(s string, c chan int) {
//      sum := 0
        go func(){
        for i:=11; i<=20;i++ {
                c <- i  
                }
 //       close(c)
//	wg.Done()
        }()
	wg.Done()       
        //c <- sum // send sum to c
}


func sum3(s string, c chan int) {
//      sum := 0
        go func(){
        for i:=21; i<=30;i++ {
                c <- i  
                }
 //       close(c)
//	wg.Done()
        }()
 	wg.Done()       
        //c <- sum // send sum to c
}

func sum4(s string, c chan int) {
//      sum := 0
        go func(){
        for i:=31; i<=40;i++ {
                c <- i  
                }
 //       close(c)
//	wg.Done()
        }()
 	wg.Done()       
        //c <- sum // send sum to c
}

func sum5(s string, c chan int) {
//      sum := 0
        go func(){
        for i:=41; i<=50;i++ {
                c <- i  
                }
 //       close(c)
//	wg.Done()
       }()
        wg.Done()
        //c <- sum // send sum to c
}


func main() {
//	s := []int{7, 2, 8, -9, 4, 0}

//	var wg sync.WaitGroup
	wg.Add(5)
	c := make(chan int)
	go sum1("First:",c)
	go sum2("Second:",c)
	go sum3("Third:",c)
	go sum4("Fourth:",c)
	go sum5("Fifth:",c)
//	go sum(s[len(s)/2:], c)
//	close(c)
	for i := range c{
//	x := <-i// receive from c

	fmt.Println(i)
}

close(c)
//wg.Wait()
}

