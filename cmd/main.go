package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/sinute/gw2w"
)

var addr = flag.String("addr", ":7788", "bind address")

func main() {
	flag.Parse()

	log.Println("Running...")
	http.Handle("/", http.TimeoutHandler(&gw2w.Handler{}, 3*time.Second, ""))
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Println(err)
	}
}
