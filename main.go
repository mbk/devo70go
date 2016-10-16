package main

import (
	"context"
	"fmt"
)

func greeter(cancelFunc context.CancelFunc, message chan string) {


	select {
	case text := <- message:
		fmt.Printf("Received a message %s", text)
		cancelFunc()
	}
}

func hello(ctx context.Context, message string) {
	sendChannel := make(chan string)
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc() //good practice to ensure we always call it when done.
	go greeter(cancelFunc, sendChannel)
	sendChannel <- "Hello world"

	select {
	case <-cancelCtx.Done():
		//Which will exit the program. Not using the context would have done the same - but as a demo....
		return
	}
}

func main() {
	hello(context.Background(), "Hello World")
}
