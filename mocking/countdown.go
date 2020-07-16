/*
Module to display a countdown.

N.B> Use of Sleeper interface to allow mocking in test

See "github.com/kami-zh/go-capturer" for simplification of output capture in testing of main

*/
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
