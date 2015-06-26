package main

import (
	"flag"
	"fmt"
	"github.com/sharpyfox/mock-bidder/http_handlers"
	"github.com/sharpyfox/mock-bidder/utils"
	"log"
	"net/http"
	"strconv"
)

func main() {
	pr := flag.Float64("pr", 0.9, "non empty response probability")
	portPtr := flag.Int("p", 7040, "port to start http server")
	showVersion := flag.Bool("version", false, "print version string")
	flag.Parse()

	if *showVersion {
		fmt.Println(utils.Version("mock-bidder"))
		return
	}

	h := http_handlers.RequestsHandler{Probability: float32(*pr)}
	http.HandleFunc("/auctions", func(w http.ResponseWriter, r *http.Request) {
		h.HandleResponse(w, r)
	})
	log.Println("INFO  Starting server on 0.0.0.0:" + strconv.Itoa(*portPtr))
	http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil)
}
