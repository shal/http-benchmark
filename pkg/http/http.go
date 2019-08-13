package http

import (
	"fmt"
	"net/http"
)

// Version returns http repsonce with current benchmark version.
func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{ "version": "1.0.0" }`)
}

// Run starts the webserer.
func Run(port string) {
	http.Handle("/", http.HandlerFunc(Version))

	fmt.Printf("Listening on port %s...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
