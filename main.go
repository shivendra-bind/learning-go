package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shivendra-bind/learning-go/di"
	"github.com/shivendra-bind/learning-go/hello"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func main() {
	fmt.Println(hello.Hello("SB", ""))
	sleeper := &DefaultSleeper{}
	di.Countdown(os.Stdout, sleeper)
}
