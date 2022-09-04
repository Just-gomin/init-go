/*
	# 개방 폐쇄 원칙(OCP : Open-Closed Principle)
*/

package main

/*
	// #나쁜 설계
	// 함수 내에 Switch Case 문으로 각각을 동작을 구분하고 있습니다.
	// 새로운 전송 방법을 추가하기 위해선 기존에 작성된 내용에 새로운 타입을 추가 해야합니다.
	// 따라서 변경에 열려
	func SendReport(r *Report, method SendType, receiver string) {
		switch method {
		case Email:
			// 이메일로 전송
		case Fax:
			// 팩스로 전송
		case PDF:
			// pdf 파일로 전송
		case Printer:
			// 프린터로 인쇄
		}
	}
*/

type Report interface {
	Report() string
}

type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct{}

func (e *EmailSender) Send(r Report) {}

type FaxSender struct{}

func (f *FaxSender) Send(r Report) {}

type PDFSender struct{}

func (p *PDFSender) Send(r Report) {}

type PrinterSender struct{}

func (p *PrinterSender) Send(r Report) {}
