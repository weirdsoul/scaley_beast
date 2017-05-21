goog.provide('browser_instruments.AnalogGauge');

goog.require('goog.dom');
goog.require('goog.dom.dataset');

/**
 * Construct an AnalogGauge instance associated with the specified DOM element..
 * @constructor
 * @param {!Element} domElement The associated DOM element.
 */
browser_instruments.AnalogGauge = function(domElement) {
    this.domElement = domElement;
    if (goog.dom.dataset.has(domElement, "transform")) {
	// Data transform is used to convert the input signal x to degrees.
	this.speedToDeg = new Function("x", goog.dom.dataset.get(domElement, "transform"));
    } else {
	console.log("Using fallback function for input signal conversion.");
	/**Completely linear. Expects the input to be degrees.
	 * @param {number} x Input to be converted.
	 * @return {number} Resulting degrees for the gauge hand.
	 */
	this.speedToDeg = function(x) { return x; };
    }
};

/**
 * Called to update an instrument with a new piece of data.
 * @param {!Object} msg A JSON object containing all the update information.
 */
browser_instruments.AnalogGauge.prototype.updateInstrument = function(msg) {
    // TODO(aeckleder): We only care about speed right now. But we want a generic gauge,
    // so do not hardcode this here.
    if (msg["Index"] != "3") return;
    var needleElement = goog.dom.getElementByClass("needle", this.domElement);
    if (needleElement != null) {
	/** @type {?Array<?Object>} **/
	var valueSet = msg["Values"];
	if (valueSet.length > 0) {
	    var gaugeValue = parseFloat(valueSet[0]);
	    if (gaugeValue != null) {
		var degree = this.speedToDeg(gaugeValue);
		// TODO(aeckleder): Do not hardcode center point here.
		needleElement.setAttribute("transform",
					   "rotate(" + degree + " 240 240)");
	    }
	} else {
	    console.error("Couldn't find value set.");
	}
    } else {
	console.error("Couldn't find needle element.");
    }    
};
