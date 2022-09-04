/*
	# 단일 책임 원칙 (SRP : Single Responsibility Principle)
*/

package main

import "fmt"

/*
	// #나쁜 설계
	// 각각의 Report 객체들은 보고서라는 책임과 보고서 전송이라는 책임을 함께 지고 있습니다.
	// 따라서 새로운 Report가 생길 때마다 보고서 전송 메서드도 새롭게 만들어져야합니다.
	type FinanceReport struct {
		report string
	}

	func (r *FinanceReport) SendReport(email string) {
		// Finance Report에 의존 하는 보고 함수
	}

	type MarketingReport struct {
		report string
	}

	func (r *MarketingReport) SendReport(email string) {
		// Marketing Report에 의존 하는 보고 함수
	}
*/

// 좋은 설계
type Report interface {
	Report() string
}

type ReportSender struct {
}

func (s *ReportSender) SendReport(report Report) {
	var content string = report.Report()
	fmt.Println(content)
}

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type MarketingReport struct {
	report string
}

func (r *MarketingReport) Report() string {
	return r.report
}
