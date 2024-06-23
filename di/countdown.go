package di

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {

		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}
