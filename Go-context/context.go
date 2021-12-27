package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	canceledChan := make(chan bool)

	go func() {
		select {
			case <- ctx.Done():
				fmt.Println(ctx.Err())
				canceledChan <- true
				return
		}
	}()

	isCanceledChan := <- canceledChan
	if (isCanceledChan) {
		close(canceledChan)
		return
	}
	time.Sleep(time.Second * 100)
	fmt.Println("end")
}