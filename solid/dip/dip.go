/*
# 의존 관계 역전 원칙(DIP: Dependency Inversion Principle)
*/
package main

import "fmt"

/*
	// # 나쁜 설계
	type Mail struct {
		alarm Alarm
	}

	type Alarm struct {
	}

	func (a *Alarm) Alarm() {}

	func (m *Mail) OnRecv() {
		m.alarm.Alarm()
	}
*/

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct{}

func (a *Alarm) OnFire() {
	fmt.Println("알람~!")
}

func main() {
	var mail = &Mail{}
	var listener EventListener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
