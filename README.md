# go-enlightenment
A journey into the world of GOlang

Following the TDD examples laid down by quii in [learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

## Useful Bits
Need some local docs ?

```go get golang.org/x/tools/cmd/godoc```

Need to ensure return types are checked ?

```go get -u github.com/kisielk/errcheck```

Need to test main method to increase coverage ?

```go get -u github.com/kami-zh/go-capturer```

```
func TestMainFunc(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		main()
	})

	assertOutput(out, expectedDefaultOutput, t)
}
```

## Sound Bytes
"Everything you need to know is printed to your terminal - all you have to do is be patient enough to read it."