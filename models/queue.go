package models

type Message struct {
	Type    string
	Content string
	Email   string
}

type MessageQueue struct {
	queue chan Message
}

func NewQueue(size int) *MessageQueue {
	return &MessageQueue{
		queue: make(chan Message, size),
	}
}
