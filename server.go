package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MugiOKa/simple-api-server/hello"
	"github.com/MugiOka/xrequestid"
	"github.com/urfave/negroni"
)

func Index(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3000/test")
	defer resp.Body.Close()
	for k, v := range resp.Header {
		fmt.Printf("%v : %v\n", k, v)
	}

	io.WriteString(w, hello.Hello())
}

func Test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hello())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/test", Test)
	n := negroni.Classic() // Includes some default middlewares
	n.Use(xrequestid.New(16))
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
