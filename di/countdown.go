package di

import (
	"fmt"
	"io"
	"time"
)

func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {

		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, "Go!")
}
