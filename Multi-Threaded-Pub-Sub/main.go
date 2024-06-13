package main

import (
	"strconv"
	"sync"
)

func main() {
	wgPublishers := new(sync.WaitGroup)
	wgSubscribers := new(sync.WaitGroup)

	topic1 := NewTopic()
	topic2 := NewTopic()

	Close := func(topic *Topic) {
		for _, channel := range topic.Channels {
			close(channel)
		}
	}

	publisher1 := NewPublisher("1")
	publisher2 := NewPublisher("2")

	wgPublishers.Add(4)
	go publisher1.Publish("topic1", topic1, wgPublishers)
	go publisher1.Publish("topic2", topic2, wgPublishers)
	go publisher2.Publish("topic1", topic1, wgPublishers)
	go publisher2.Publish("topic2", topic2, wgPublishers)

	for i := 1; i < 5; i++ {
		wgSubscribers.Add(1)
		subscriber := NewSubscriber(strconv.Itoa(i))
		newChannel := make(chan string)
		if i%2 == 0 {
			topic1.Channels = append(topic1.Channels, newChannel)
			go subscriber.subscribe(newChannel, wgSubscribers)
		} else {
			topic2.Channels = append(topic2.Channels, newChannel)
			go subscriber.subscribe(newChannel, wgSubscribers)
		}
	}

	wgPublishers.Wait()

	Close(topic2)
	Close(topic1)

	wgSubscribers.Wait()
}
