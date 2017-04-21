package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[0:1], c)
	fmt.Println(<-c)
	go sum(a[:len(a) / 2], c)
	fmt.Println(<-c)
	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y)
}