package main

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

/* --------------- */
// 아무 인터페이스에도 포함될 수 없음
type rw1 struct{}

func (rw *rw1) Read(n int, err error)  {}
func (rw *rw1) Write(n int, err error) {}

/* --------------- */
// Reader, Writer, ReadWriter 모두에 포함될 수 있믕
type rw2 struct{}

func (rw *rw2) Read(n int, err error)  {}
func (rw *rw2) Write(n int, err error) {}
func (rw *rw2) Close()                 {}

/* --------------- */
// 아무 인터페이스에도 포함될 수 없음
type rw3 struct{}

func (rw *rw3) Read(n int, err error) {}

/* --------------- */
// Writer 인터페이스에 포팜될 수 있음
type rw4 struct{}

func (rw *rw4) Write(n int, err error) {}
func (rw *rw4) Close()                 {}
