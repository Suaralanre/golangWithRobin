package handler

import (
	"fmt"
	"net/http"
)

func (h *HTTPHandler) ABC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
	h.Logger.Print("an error occurred!")
}
