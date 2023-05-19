package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 消费者
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("Consumer exit.")
				return
			case num := <-ch:
				fmt.Printf("Consume %d\n", num)
			}
		}
	}(ctx)

	// 生产者
	ticker := time.NewTicker(time.Second)
	prodNum := 1
	for range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("Producer exit.")
			return
		default:
			ch <- prodNum
			fmt.Printf("Produce %d\n", prodNum)
			prodNum += 1
		}
	}
}
