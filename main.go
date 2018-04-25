package main

import (
	"fmt"
	"time"
)

type Ball struct {
	count int
}

func main() {

	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball)
	time.Sleep(time.Second * 1)
	<-table

}
func player(action string, table chan *Ball) {
	for {
		ball := <-table
		ball.count++
		fmt.Println(action, ball.count)
		time.Sleep(time.Millisecond * 100)
		table <- ball
	}
}
