package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := flag.String("i", "0.0.0.0", "listening interface (default 0.0.0.0)")
	p := flag.Int("p", 8080, "listening port (default 8080)")
	flag.Parse()

	http.HandleFunc("/", handler)

	u := fmt.Sprintf("%s:%d", *h, *p)
	log.Println("start server at:", u)
	log.Fatal(http.ListenAndServe(u, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.String())
}
