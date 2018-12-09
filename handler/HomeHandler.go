package handler

import (
	"fmt"
	"net/http"
)

func HandleHome() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// The if is needed because of the default Mux
		// ref: https://github.com/golang/go/issues/4799
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_, _ = fmt.Fprintf(w, "You shouldn't be here")
	}
}
