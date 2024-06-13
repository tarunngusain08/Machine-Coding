package main

import (
	"fmt"
	"sync"
)

type subscriber struct {
	Name string
}

type Subscriber interface {
	subscribe(channel <-chan string, wg *sync.WaitGroup)
}

func NewSubscriber(name string) Subscriber {
	return &subscriber{Name: name}
}

func (s *subscriber) subscribe(channel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-channel:
			if !ok {
				fmt.Println("Channel closed, subscriber", s.Name, "exiting")
				return // Exit the goroutine if channel is closed
			}
			fmt.Println(val + "->subscriber" + s.Name)
		}
	}
}
