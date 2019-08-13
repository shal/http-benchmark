package fasthttp

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Version returns http repsonce with current benchmark version.
func Version(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	fmt.Fprintf(ctx, `{ "version": "1.0.0" }`)
}

// Run starts the webserer.
func Run(port string) {
	fmt.Printf("Listening on port %s...\n", port)

	err := fasthttp.ListenAndServe(":"+port, Version)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
