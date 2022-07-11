package main

import (
	"fmt"
	"server/serverHello"
	"server/serverInc"
)

func main() {
	var number int
L:
	for {
		fmt.Println("Enter:")
		fmt.Println("1 for start Hello server;")
		fmt.Println("2 for start Increment server;")
		fmt.Println("3 for exit.")
		fmt.Scan(&number)
		switch number {
		case 1:
			serverHello.ConnectHello()
		case 2:
			serverInc.ConnectInc()
		case 3:
			break L
		}
	}
}
