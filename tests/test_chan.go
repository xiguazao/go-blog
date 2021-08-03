package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	//timeout := make(chan bool)
	//go func() {
	//	//time.Sleep(time.Second*1)
	//	fmt.Println(123)
	//	timeout <- true
	//}()

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	for cv := range c {
		fmt.Println(cv)
		// do something
	}
	//x := <-c
	//y := <-c
	//x, y := <-c, <-c // 从 c 中接收

	//fmt.Println(x, y, x+y)
}
