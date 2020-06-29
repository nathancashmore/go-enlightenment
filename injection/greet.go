/*
An object that implements the io.Writer interface to output a greeting for a given name.

This may then be used by any object that implements the interface.  Below is an example
to provide a default greeting for a HTTP server.

  func HttpGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
  }

  func main() {
	_ = http.ListenAndServe(":5000", http.HandlerFunc(HttpGreetHandler))
  }

*/
package injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
