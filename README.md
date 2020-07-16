# go-enlightenment : A journey into the world of GOlang

## Overview [![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/nathancashmore/go-enlightenment)](https://pkg.go.dev/mod/github.com/nathancashmore/go-enlightenment)

Following the TDD examples laid down by quii in [learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

## Useful Bits
Need some local docs ?

    go get golang.org/x/tools/cmd/godoc

Need to ensure return types are checked ?

    go get -u github.com/kisielk/errcheck

Need to test main method to increase coverage ?

    go get -u github.com/kami-zh/go-capturer

    func TestMainFunc(t *testing.T) {
        out := capturer.CaptureStdout(func() {
            main()
        })
    
        assertOutput(out, expectedDefaultOutput, t)
    }

Need latest documentation uploaded to pkg.go.dev ?

    curl https://proxy.golang.org/github.com/nathancashmore/go-enlightenment/@v/vX.X.X.info
    GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/nathancashmore/go-enlightenment@vX.X.X


## Sound Bytes
"Everything you need to know is printed to your terminal - all you have to do is be patient enough to read it."

## Install
    go get github.com/nathancashmore/go-enlightenment

## License

MIT.
