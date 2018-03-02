package main

import (
    //"fmt"
    "net/http"
	//"os"
	//"io/ioutil"
	//"fmt"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there!")
	//b, err := ioutil.ReadFile("index.html")
	b, err := ioutil.ReadFile("register.html")
	if err != nil {
		panic(err)
	}

	w.Write(b)
}

func main() {
    http.HandleFunc("/", handler)
    //http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)

}
