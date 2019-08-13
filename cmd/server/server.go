package main

import (
	"flag"

	"github.com/openware/playground/pkg/fasthttp"
	"github.com/openware/playground/pkg/http"
)

var (
	port = flag.String("port", "8080", "Port to run server")
	fast = flag.Bool("fast", false, "Enables fast mode")
)

func main() {
	flag.Parse()

	if *fast {
		fasthttp.Run(*port)
	} else {
		http.Run(*port)
	}
}
