package main

import (
	"fmt"
	"sync"
)

type Subscriber struct {
	Name string
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{Name: name}
}

func (s *Subscriber) subscribe(channel <-chan string, signal <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-channel:
			if !ok {
				fmt.Println("Channel closed, subscriber", s.Name, "exiting")
				return // Exit the goroutine if channel is closed
			}
			fmt.Println(val + "sub" + s.Name)
		case <-signal:
			fmt.Println("Received signal to stop, subscriber", s.Name, "exiting")
			return // Exit the goroutine if signal received
		}
	}
}
