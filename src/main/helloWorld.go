package main

import (
	"fmt"
	"strconv"
)

var count int
var chans []chan int

func main() {
	fmt.Println("Hello World !")
	

	var size=10
	chans=make([]chan int,10)

	for i:=0;i<size;i++{
		go test(i)
	}

	for j:=0;j<size;j++{

		<-chans[j]
	}

}

func test(i int){
	count ++;
	fmt.Println(strconv.Itoa(count))
	chans[i]<-1
}
