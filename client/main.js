goog.require('browser_instruments.AnalogGauge');
goog.require('goog.array');
goog.require('goog.dom');
goog.require('goog.events');
goog.require('goog.net.WebSocket');
goog.require('goog.net.WebSocket.MessageEvent');

/**
 * Main entry point into the client application.
 * @export
 */
function main() {
    console.debug("Connecting to websocket.");

    var analogGaugeElements = goog.dom.getElementsByClass("AnalogGauge");
    console.log("Num gauges: " + analogGaugeElements.length);

    /** @type {!Array<!browser_instruments.AnalogGauge>} **/
    var analogGauges = [];

    goog.array.forEach(analogGaugeElements,
		       function (domElement) {
			   analogGauges.push(
			       new browser_instruments.AnalogGauge(domElement));
			   console.log("Created analog gauge.");
		       });


    var ws = new goog.net.WebSocket(true);
    goog.events.listen(ws, goog.net.WebSocket.EventType.MESSAGE,
		       /** @param {!goog.net.WebSocket.MessageEvent} e **/
		       function(e) {
			   var msg = JSON.parse(e.message);
			   goog.array.forEach(analogGauges, function (gauge) {
			       if (msg != null) {
				   /** @type {?Array<?Object>} **/
				   var updates = msg["Updates"];
				   if (updates != null) {
				       updates.forEach(function(item, index, array) {
					   if (item != null) {
					       gauge.updateInstrument(item);
					   }
				       });
				   }
			       }
			   });
			   ws.send("ACK");
		       });
    var webSocket = "ws://" + location.host + "/ws";
    console.log("Opening web socket at " + webSocket);
    ws.open(webSocket);
}

