/*
Module that provides a Store object that facilitates Fetch and Cancel methods to simulate obtaining and cancelling the retrieval
of a product from the store.

The context package provides functions to derive new Context values from existing ones.
These values form a tree: when a Context is canceled, all Contexts derived
from it are also canceled.
*/

package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())

		if err != nil {
			return // log error
		}
		_, _ = fmt.Fprint(writer, data)
	}
}
