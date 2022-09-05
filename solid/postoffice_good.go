package main

type Person interface {
	Name() string
	PhoneNumber() string
	Address() string
}

type Post interface {
	Sender() Person
	Recipient() Person
	Contents() string
}

type PostReceiver interface {
	Receive(Post)
}

type PostSender interface {
	Send(Post)
}

type NormalSender struct{}

func (ns *NormalSender) Send(post Post) {}

type ExpressSender struct{}

func (es *ExpressSender) Send(post Post) {}

type AbroadSender struct{}

func (as *AbroadSender) Send(post Post) {}
