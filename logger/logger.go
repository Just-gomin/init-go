package logger

import "fmt"

/* -------------------------- */
/*           Colors           */
/* -------------------------- */
const red string = "\033[31m"
const green string = "\033[32m"
const yellow string = "\033[33m"
const blue string = "\033[34m"
const purple string = "\033[35m"
const cyan string = "\033[36m"
const white string = "\033[37m"

/* -------------------------- */
/*       Format String        */
/* -------------------------- */
const normalString = "[%s] [%-20s] : %s"

const warnString = yellow + "[" + "%s" + "]" + white + " " + // specialchanr
	cyan + "[" + "%-20s" + "]" + white + " " + // title
	": %s\n" // body

const errorString = red + "[" + "%s" + "]" + white + " " + // specialchanr
	cyan + "[" + "%-20s" + "]" + white + " " + // title
	": %s\n" // body

/* -------------------------- */
/*       Print Function       */
/* -------------------------- */
func Normal(specialChar string, title string, body string) {
	fmt.Printf(normalString, specialChar, title, body)
}

func Warning(specialChar string, title string, body string) {
	fmt.Printf(warnString, specialChar, title, body)
}

func Error(specialChar string, title string, content string) {
	fmt.Printf(errorString, specialChar, title, content)
}
