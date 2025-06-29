package main

import (
	"fmt"
	"net/http"
)

type ApiHandler struct{}

func (ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Stranger Things\nBreaking Bad\nWednesdaay\nDARK ")
}

func main() {

	http.ListenAndServe(":8000", &ApiHandler{})

}
