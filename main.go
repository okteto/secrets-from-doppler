package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting hello-world server...")
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, my name is %s, and my favorite color is %s", os.Getenv("MY_NAME"), os.Getenv("MY_COLOR"))
}
