package main

import "sync"

type publisher struct {
	Name string
}

type Publisher interface {
	Publish(message string, topic *Topic, wg *sync.WaitGroup)
}

func NewPublisher(name string) Publisher {
	return &publisher{Name: name}
}

func (p *publisher) Publish(message string, topic *Topic, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, channel := range topic.Channels {
		channel <- message + "->publisher" + p.Name
	}
}
