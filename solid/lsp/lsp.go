/*
	# 리스코프 치환 원칙(LSP: Liskov Substitution Principle)
*/

package main

/*
	// # 리스코프 치환 원칙 예시 코드
	type T interface {
		Something()
	}

	type S struct{}

	func (s *S) Something() {}

	type U struct{}

	func (u *U) Something() {}

	func q(t T) {
		t.Something()
	}

	var y = &S{}
	var u = &U{}

	func main(){
		q(y)
		q(u)
	}
*/

// # 나쁜 설계
type Report interface {
	Report() string
}
type MarketingReport struct {
	report string
}

func (m *MarketingReport) Report() string {
	return m.report
}

func SendReport(r Report) {
	if _, ok := r.(*MarketingReport); ok { // r이 마케팅 보고서일 경우 패닉
		panic("Can't send MarketingReport")
	}
}

func main() {
	var report = &MarketingReport{}

	SendReport(report) // 패닉 발생
}
