package main

import (
	"fmt"
	"time"
)

// с помощью таймера
func sleep(duration time.Duration) {
	start := time.Now()
	// находимся в цикле до тех пор, пока время,
	//прошедшее с запуска функции меньше времени запрашиваемого
	for {
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	sleep(time.Second)
}
