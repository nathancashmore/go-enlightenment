package main

import (
	"github.com/kami-zh/go-capturer"
	"reflect"
	"testing"
)

type CountdownSpy struct {
	Calls []string
}

func (s *CountdownSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func (s *CountdownSpy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

const write = "write"
const sleep = "sleep"

const expectedDefaultOutput = `3
2
1
Go!`

func TestCountdown(t *testing.T) {
	spy := &CountdownSpy{}

	Countdown(spy, spy)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	assertCountdownOperations(spy.Calls, want, t)
}

func TestCountdownMain(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		main()
	})

	assertOutput(out, expectedDefaultOutput, t)
}

func assertOutput(got, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("called %q times but wanted %q", got, want)
	}
}

func assertCountdownOperations(got []string, want []string, t *testing.T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted calls %v got %v", want, got)
	}
}
