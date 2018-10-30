// Package webservice contains all the necessary code to serve data state
// updates via JSON as well as the main html page and javascript.
package webservice

import (
	"fmt"
	"io"
	"log"
        "math"
	"net/http"
        "strconv"

	"github.com/gorilla/websocket"
	"github.com/weirdsoul/scaley_beast/scalestate"
)

// websocketHandler writes all updates obtained from planeState to the socket.
// All communication is unidirectional from server to client.
type websocketHandler struct {
	// planeState points to the current plane state, which keeps being
	// updated using UDP packets received from the flight simulator.
	planeState *planestate.PlaneState
	// sequenceNumber is the sequence number of the last update that
	// has been processed.
	// TODO(aeckleder): This is a bug, the sequenceNumber should be
	// per connection and not per handler. Write a test that verifies
	// the correct behavior!
	sequenceNumber planestate.SequenceNumber
}

var upgrader = websocket.Upgrader{}

// jsonResponse represents a single JSON response, which can hold an arbitrary
// number of updates to the plane state. We always transmit all relevant
// updates (updates that have not been superseded by new data) since the last
// update, which we identify by sequence number.
type jsonResponse struct {
	// Updates contains all the updates since the last message.
	Updates []planestate.DataSet
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
			Updates: updates,
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

func reanimateHandler(staticDir string, serialWriter io.Writer,
                      w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
        if err := r.ParseForm(); err != nil {
          fmt.Fprintf(w, "ParseForm() err: %v", err)
          return
        }
        phase_1 := r.FormValue("phase_1")
        phase_2 := r.FormValue("phase_2")
        phase_3 := r.FormValue("phase_3")

        access_code := r.FormValue("access_code")

        p1, _ := strconv.ParseFloat(phase_1, 64)
        p2, _ := strconv.ParseFloat(phase_2, 64)
        p3, _ := strconv.ParseFloat(phase_3, 64)

        var p1_golden float64 = 110.00
        var p2_golden float64 = 330.00
        var p3_golden float64 = 770.00

	if access_code == "afterlife" && math.Abs(p1 - p1_golden) < 50 &&
           math.Abs(p2 - p2_golden) < 50 && math.Abs(p3 - p3_golden) < 50 {
		if _, err := serialWriter.Write([]byte(",")); err != nil {
		   fmt.Fprintf(w, "Error contacting Eddie: %v", err)
                   return
		}
                http.ServeFile(w, r, staticDir + "/success.html")
		return
	} else {
		if _, err := serialWriter.Write([]byte(".")); err != nil {
		   fmt.Fprintf(w, "Error contacting Eddie: %v", err)
                   return
		}
                http.ServeFile(w, r, staticDir + "/failure.html")
		return
        }
    default:
    	 http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
	 return
    }
}

// ServeHTTP serves http on the specified port using the specified plane data.
// It offers a simple JSON service to retrieve plane state and serves static
// http content from the specified directory. This function never returns.
// The serialWriter instance will be used to control the Arduino hardware.
func ServeHTTP(port int, staticDir string, planeState *planestate.PlaneState,
               serialWriter io.Writer) {
	// Serve the static directory as the root.
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.Handle("/ws", createWebsocketHandler(planeState))
	http.HandleFunc("/reanimate", func(w http.ResponseWriter, r* http.Request) {
		reanimateHandler(staticDir, serialWriter, w, r)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
