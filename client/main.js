goog.require('goog.dom');
goog.require('goog.events');
goog.require('goog.json');
goog.require('goog.net.WebSocket');
goog.require('goog.net.WebSocket.MessageEvent');

/**
 * Main entry point into the client application.
 * @export
 */
function main() {
    console.debug("Connecting to websocket.");
    var ws = new goog.net.WebSocket(true);
    goog.events.listen(ws, goog.net.WebSocket.EventType.MESSAGE,
		       /** @param {!goog.net.WebSocket.MessageEvent} e **/
		       function(e) {
			   var con = goog.dom.getElement("console");
			   if (con != null) {
			       var msg = goog.json.parse(e.message);
			       goog.dom.append(con, goog.json.serialize(msg));
			   } else {
			       console.error("Couldn't find console.");
			   }
			   ws.send("ACK");
		       });
    ws.open("ws://localhost:8080/ws");
}

