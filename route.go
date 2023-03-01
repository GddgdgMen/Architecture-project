package main

import "net/http"

func main() {
	http.HandleFunc("/time", getTime)

	Begin(8795)
}
