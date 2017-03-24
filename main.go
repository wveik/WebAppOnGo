package main

import (
	"fmt"
	"net/http"
)

func indexHandler (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main()  {
	fmt.Println("Listening on port :3000")

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", nil)

}
