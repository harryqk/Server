package main

import (
	"fmt"
	"strconv"
)

func ToolTool(){
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++{
			fmt.Println("fun1before" + strconv.Itoa(i))
			chan1 <- i
			fmt.Println("fun1after" + strconv.Itoa(i))
		}
		close(chan1)
	}()

	go func() {
		for i := range chan1{
			fmt.Println("fun2_" + strconv.Itoa(i))
}
chan2 <- 1
}()
	<- chan2
	close(chan2)
}