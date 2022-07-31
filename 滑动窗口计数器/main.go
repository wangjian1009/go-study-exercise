package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Metrics struct {
}

func main() {
	rand.Seed(time.Now().UnixNano())

	counter := NewSecWindowCounter(2)

	go func() {
		for i := 0; i < 1000; i++ {
			if rand.Intn(100) < 50 {
				counter.AddEvent(PASS)
			} else {
				counter.AddEvent(ERR)
			}
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
	}()

	for {
		time.Sleep(time.Second)
		pass, err := counter.GetData()
		fmt.Println("统计: ", pass, err)
	}
}
