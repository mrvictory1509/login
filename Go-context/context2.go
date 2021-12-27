package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, cancel := context.WithCancel(context.Background())

	timeoutCtx, _ := context.WithTimeout(parentCtx, time.Second*10)
	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select {
	case <-timeoutCtx.Done():
		fmt.Println("Done")
	}
}
