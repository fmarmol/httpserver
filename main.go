package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	Verbose bool
}

func main() {
	inter := flag.String("i", "0.0.0.0", "listening interface (default 0.0.0.0)")
	p := flag.Int("p", 8080, "listening port (default 8080)")
	v := flag.Bool("v", false, "verbose mode")
	flag.Parse()

	h := Handler{Verbose: *v}
	http.Handle("/", h)

	u := fmt.Sprintf("%s:%d", *inter, *p)
	log.Println("start server at:", u)
	log.Fatal(http.ListenAndServe(u, nil))
}
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.String())
	if h.Verbose {
		defer r.Body.Close()
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error:", err)
			// TODO set header properly
			return
		}
		log.Println(string(data))
	}

}
