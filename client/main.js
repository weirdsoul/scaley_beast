goog.require('goog.dom');
goog.require('goog.events');
goog.require('goog.json');
goog.require('goog.log');
goog.require('goog.net.WebSocket');
goog.require('goog.net.WebSocket.MessageEvent');

/**
 * Main entry point into the client application.
 * @export
 */
function main() {
    var log = goog.log.getLogger("main");
    
    var ws = new goog.net.WebSocket(true);
    goog.events.listen(ws, goog.net.WebSocket.EventType.MESSAGE,
		       /** @param {!goog.net.WebSocket.MessageEvent} e **/
		       function(e) {
			   var console = goog.dom.getElement("console");
			   if (console != null) {
			       var msg = goog.json.parse(e.message);
			       goog.dom.append(console, e.message);
			       goog.log.info(log, "Received message " + msg);
			   } else {
			       goog.log.error(log, "Couldn't find console.");
			   }
			   ws.send("ACK");
		       });
    ws.open("ws://localhost:8080/ws");
}

