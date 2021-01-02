package main

import (
	"fmt"
	"net/http"
)

func mainHelloWorld() {

	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello world")
	})

	http.ListenAndServe("127.0.0.1:3000", nil)

}
