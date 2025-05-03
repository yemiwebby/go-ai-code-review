package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(3)
}

func Countdown(start int, writer ...io.Writer) {

	output := io.Writer(os.Stdout)
	if len(writer) > 0 {
		output = writer[0]
	}

	for {
		if start == 0 {
			fmt.Fprintln(output, "Go!")
			return
		}

		fmt.Fprintln(output, start)
		time.Sleep(time.Second * 1)
		start--
	}

}
