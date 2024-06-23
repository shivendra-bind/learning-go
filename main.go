package main

import (
	"fmt"
	"os"

	"github.com/shivendra-bind/learning-go/di"
	"github.com/shivendra-bind/learning-go/hello"
)

func main() {
	fmt.Println(hello.Hello("SB", ""))
	di.Countdown(os.Stdout)
}
