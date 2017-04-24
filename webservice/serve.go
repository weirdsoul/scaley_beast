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
	// planeState points to the current plane state, which keeps being
	// updated using UDP packets received from the flight simulator.
	planeState *planestate.PlaneState
	// sequenceNumber is the sequence number of the last update that
	// has been processed.
	sequenceNumber planestate.SequenceNumber
}

var upgrader = websocket.Upgrader{}

// jsonResponse represents a single JSON response, which can hold an arbitrary
// number of updates to the plane state. We always transmit all relevant
// updates (updates that have not been superseded by new data) since the last
// update, which we identify by sequence number.
type jsonResponse struct {
	// Updates contains all the updates since the last message.
	updates []planestate.DataSet
}

func (h *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer ws.Close()

	// A never ending stream of Hello world.
	for {
		s, updates := h.planeState.GetControlDataSince(h.sequenceNumber, true)
		h.sequenceNumber = s

		response := &jsonResponse{
			updates: updates,
		}
		if ws.WriteJSON(&response); err != nil {
			log.Println("ws.WriteMessage: ", err)
			return
		}
		// We ignore the actual message.
		messageType, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("ws.ReadMessage: ", err)
			return
		}
		if messageType != websocket.TextMessage {
			log.Printf("ws.ReadMessage: messageType(%v) != TextMessage)\n",
				messageType)
			return
		}
	}
}

// createWebsocketHandler creates a handler that serves the web socket used
// for client server communication.
func createWebsocketHandler(planeState *planestate.PlaneState) *websocketHandler {
	return &websocketHandler{
		planeState:     planeState,
		sequenceNumber: 0,
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
