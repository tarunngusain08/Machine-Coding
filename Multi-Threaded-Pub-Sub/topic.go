package main

type Topic struct {
	Channels []chan string
}

func NewTopic() *Topic {
	return &Topic{}
}
