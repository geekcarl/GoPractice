package main

import (
	"fmt"
	"strconv"
)
var size = 20
var count int

func main() {
	fmt.Println("Hello World !")
	
	chans := make([]chan int,size)

	for i:=0;i<size;i++{
		chans[i] = make(chan int)
		go test(chans[i],i)
	}

	for _,c:=range(chans){

		v:= <-c
		fmt.Println("out-"+strconv.Itoa(v))
	}
}

func test(c chan int,index int){
	c<-index
	count ++;
	fmt.Println(strconv.Itoa(count))
}
