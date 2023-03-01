package main

import (
	"net/http"
	"strconv"
)

func Begin(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		errorHandler(err)
	}
}
