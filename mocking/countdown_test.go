package main

import (
	"bytes"
	"github.com/kami-zh/go-capturer"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const expectedDefaultOutput = `3
2
1
Go!`

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	assertSleepCalls(sleeper.Calls, 3, t)
	assertOutput(buffer.String(), expectedDefaultOutput, t)
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

func assertSleepCalls(got int, want int, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("called %d times but wanted %d", got, want)
	}
}
