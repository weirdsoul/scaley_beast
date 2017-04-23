// Package webservice contains all the necessary code to serve data state
// updates via JSON as well as the main html page and javascript.
package webservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/weirdsoul/browser_instruments/planestate"
)

// websocketHandler writes all updates obtained from planeState to the socket.
// All communication is unidirectional from server to client.
type websocketHandler struct {
	planeState *planestate.PlaneState
}

var upgrader = websocket.Upgrader{}

type jsonResponse struct {
	Message string
}

func (*websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	ws.EnableWriteCompression(false)

	defer ws.Close()
	// A never ending stream of Hello world.
	for x := 0; x < 10; x++ {
		response := &jsonResponse{
			Message: "Hello world!",
		}
		if ws.WriteJSON(&response); err != nil {
			log.Println("ws.WriteMessage: ", err)
		}
	}
}

// createWebsocketHandler creates a handler that serves the web socket used
// for client server communication.
func createWebsocketHandler(planeState *planestate.PlaneState) *websocketHandler {
	return &websocketHandler{
		planeState: planeState,
	}
}

// ServeHTTP serves http on the specified port using the specified plane data.
// It offers a simple JSON service to retrieve plane state and serves static
// http content from the specified directory. This function never returns.
func ServeHTTP(port int, staticDir string, planeState *planestate.PlaneState) {
	// Serve the static directory as the root.
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.Handle("/ws", createWebsocketHandler(planeState))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
