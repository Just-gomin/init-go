/*
# 인터페이스 분리 원칙(ISP: Interface Segregation Principle)
*/
package main

import "time"

/*
	// # 나쁜 설계
	type Report interface {
		Report() string
		Pages() int
		Author() string
		WrittenDate() time.Time
	}

	func SendReport(r Report) {
		send(r.Report())
	}
*/

// # 좋은 설계
type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func send(message string) {}

func SendReport(r Report) {
	send(r.Report())
}
