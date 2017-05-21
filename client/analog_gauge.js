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
    /** @type {number} **/
    this.inputSource = goog.dom.dataset.has(domElement, "source")
	? parseInt(goog.dom.dataset.get(domElement, "source"), 10): 3;
    /** @type {number} **/
    this.inputIndex =  goog.dom.dataset.has(domElement, "index")
	? parseInt(goog.dom.dataset.get(domElement, "index"), 10): 0;    
    if (goog.dom.dataset.has(domElement, "transform2")) {
	// Data transform is used to convert the input signal x to degrees.
	this.speedToDeg2 = new Function("x", goog.dom.dataset.get(domElement, "transform2"));
    } else {
	console.log("Using fallback function for input signal 2 conversion.");
	/**Completely linear. Expects the input to be degrees.
	 * @param {number} x Input to be converted.
	 * @return {number} Resulting degrees for the gauge hand.
	 */
	this.speedToDeg2 = function(x) { return x; };
    }
    /** @type {number} **/
    this.inputSource2 = goog.dom.dataset.has(domElement, "source2")
	? parseInt(goog.dom.dataset.get(domElement, "source2"), 10): 3;
    /** @type {number} **/
    this.inputIndex2 =  goog.dom.dataset.has(domElement, "index2")
	? parseInt(goog.dom.dataset.get(domElement, "index2"), 10): 0;
};

/**
 * Called to update an instrument with a new piece of data.
 * @param {!Object} msg A JSON object containing all the update information.
 */
browser_instruments.AnalogGauge.prototype.updateInstrument = function(msg) {
    /** @type {?Array<?Object>} **/
    var valueSet = msg["Values"];
    if (valueSet.length == 0) return;    
    if (msg["Index"] == this.inputSource) {
	var needleElement = goog.dom.getElementByClass("needle", this.domElement);
	if (needleElement != null) {
	    var gaugeValue = parseFloat(valueSet[this.inputIndex]);
	    if (gaugeValue != null) {
		var degree = this.speedToDeg(gaugeValue);
		// TODO(aeckleder): Do not hardcode center point here.
		needleElement.setAttribute("transform",
					   "rotate(" + degree + " 240 240)");
	    }
	} else {
	    console.error("Couldn't find needle element.");
	}
    }
    // The analog gauge supports two independent input sources. Not sure if we will ever
    // need it, but this flexibility comes quite cheap. The second needle is purely
    // optional, we don't even log when it is not there.
    if (msg["Index"] == this.inputSource2) {
	var needleElement2 = goog.dom.getElementByClass("needle2", this.domElement);
	if (needleElement2 != null) {
	    var gaugeValue = parseFloat(valueSet[this.inputIndex2]);
	    if (gaugeValue != null) {
		var degree = this.speedToDeg2(gaugeValue);
		// TODO(aeckleder): Do not hardcode center point here.
		needleElement2.setAttribute("transform",
					    "rotate(" + degree + " 240 240)");
	    }
	}
    }
};
