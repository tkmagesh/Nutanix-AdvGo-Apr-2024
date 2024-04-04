/* context.WithTimeout() */

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	rootValCtx := context.WithValue(rootCtx, "root-key", "root-value")
	childValCtx := context.WithValue(rootValCtx, "child-key", "child-value")
	timeoutCtx, cancel := context.WithTimeout(childValCtx, 5*time.Second)
	defer cancel()
	ch := genNos(timeoutCtx)
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos(ctx context.Context) <-chan int {
	fmt.Printf("[genNos] value of root-key : %v\n", ctx.Value("root-key"))
	fmt.Printf("[genNos] value of child-key : %v\n", ctx.Value("child-key"))
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i += 2 {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				fmt.Println("cancellation signal received")
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
