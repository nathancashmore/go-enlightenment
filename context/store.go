/*
Provide a Store object that facilitates Fetch and Cancel methods to simulate obtaining and cancelling the retrieval
of a product from the store.

The context package provides functions to derive new Context values from existing ones.
These values form a tree: when a Context is canceled, all Contexts derived
from it are also canceled.
*/

package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			_, _ = fmt.Fprint(writer, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
