/* context.WithCancel() */

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genNos(cancelCtx)
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i += 2 {
			select {
			case <-ctx.Done():
				close(ch)
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- i
			}

		}
	}()
	return ch
}
