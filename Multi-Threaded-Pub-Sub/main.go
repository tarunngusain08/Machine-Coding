package main

import (
	"strconv"
	"sync"
)

func main() {
	wgPublishers := new(sync.WaitGroup)
	wgSubscribers := new(sync.WaitGroup)

	channel1 := make(chan string)
	channel2 := make(chan string)
	signal := make(chan struct{})

	publisher1 := NewPublisher("1")
	publisher2 := NewPublisher("2")

	wgPublishers.Add(4)
	go publisher1.publish("chan1", channel1, wgPublishers)
	go publisher1.publish("chan2", channel2, wgPublishers)
	go publisher2.publish("chan1", channel1, wgPublishers)
	go publisher2.publish("chan2", channel2, wgPublishers)

	go func() {
		wgPublishers.Wait()
		close(channel1)
		close(channel2)
		close(signal)
	}()

	for i := 1; i < 5; i++ {
		wgSubscribers.Add(1)
		subscriber := NewSubscriber(strconv.Itoa(i))
		if i%2 == 0 {
			go subscriber.subscribe(channel1, signal, wgSubscribers)
		} else {
			go subscriber.subscribe(channel2, signal, wgSubscribers)
		}
	}
	wgSubscribers.Wait()
}
