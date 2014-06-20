package greeter

import (
	"time"
)

func Greeting() string {

	msg := "this is a test " + time.Now().String()
	return msg
}
