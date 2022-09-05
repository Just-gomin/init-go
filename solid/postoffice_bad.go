package main

type SendType int

const (
	normal SendType = 1 + iota
	express
	abroad
)

type Person struct {
	name        string
	phonenumber string
	address     string
}

type Package struct {
	sender    Person
	recipient Person
	contents  string
	weight    float64
}

func (mail *Mail) ReceivePackage() {}
func (mail *Mail) SendPackage(sendType SendType) {
	switch sendType {
	case normal:
	case express:
	case abroad:
	}
}

type Mail struct {
	sender    Person
	recipient Person
	contents  string
}

func (mail *Mail) ReceiveMail() {}
func (mail *Mail) SendMail(sendType SendType) {
	switch sendType {
	case normal:
	case express:
	case abroad:
	}
}
