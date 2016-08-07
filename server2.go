package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler はリクエストされた URL の Path 要素を返します。
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter は今までの呼び出し数を返します
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		fmt.Fprintf(w, "parameter error: %s", err)
		return
	}
	var cycles int
	var err error
	if cycles, err = strconv.Atoi(r.Form.Get("cycles")); err != nil {
		log.Print(err)
		cycles = 5
	}
	Lissajous(w, cycles)
}
