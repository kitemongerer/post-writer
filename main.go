package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		io.WriteString(w, "Posts only plz!")
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	io.WriteString(w, fmt.Sprintf("Received data: %s\n", string(bytes)))
}

func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8080", nil)
}
