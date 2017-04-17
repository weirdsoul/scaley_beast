// Package webservice contains all the necessary code to serve data state
// updates via JSON as well as the main html page and javascript.
package webservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/weirdsoul/browser_instruments/planestate"
)

// ServeHTTP serves http on the specified port using the specified plane data.
// It offers a simple JSON service to retrieve plane state and serves static
// http content from the specified directory. This function never returns.
func ServeHTTP(port int, staticDir string, planeState planestate.PlaneState) {
	// Serve the static directory as the root.
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
