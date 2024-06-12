package main

import "sync"

type Publisher struct {
	Name string
}

func NewPublisher(name string) *Publisher {
	return &Publisher{Name: name}
}

func (p *Publisher) publish(message string, channel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	channel <- message + "pub" + p.Name
}
