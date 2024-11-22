package cmd

import "golang.org/x/text/message"

type Printer interface {
	Sprint(a ...interface{}) string
	Print(a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	Printf(key message.Reference, a ...interface{}) (n int, err error)
}
