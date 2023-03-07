package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	// Создаём канал и спустя определенное время отправляем о завершении через сигнал
	stop := make(chan struct{})

	go func() {
		select {
		case <-stop:
			fmt.Println("received stop signal, stopping function 1")
			return
		}
	}()

	stop <- struct{}{}
	time.Sleep(time.Millisecond * 300)

	//  через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("received ctx.Done signal, stopping function 2")
			return
		}
	}(ctx)

	cancel()
	time.Sleep(time.Millisecond * 300)

	// закрываем канал
	dataChan := make(chan int)

	go func() {
		for {
			if _, ok := <-dataChan; !ok {
				fmt.Println("channel closed, stopping function 3")
				return
			}
		}
	}()

	dataChan <- 1234
	close(dataChan)
	time.Sleep(time.Millisecond * 300)

	// Использую time.After
	dataChan = make(chan int)

	go func() {
		for {
			select {
			case <-dataChan:
				fmt.Println("got a value from dataChan")
			case <-time.After(time.Second):
				fmt.Println("got a value from time.After, stopping function 4")
				return
			}
		}
	}()

	dataChan <- 1234
	time.Sleep(time.Second * 2)
}
