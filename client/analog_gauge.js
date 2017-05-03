goog.provide('browser_instruments.AnalogGauge');

goog.require('goog.dom');

/** 
 * Construct an AnalogGauge instance associated with the specified DOM element..
 * @constructor
 * @param {!Element} domElement The associated DOM element.
 */
browser_instruments.AnalogGauge = function(domElement) {
    this.domElement = domElement;
};

/**
 * Convert airspeed in knots to degrees of needle movement.
 * @param {number} x
 * @return {number}
 */
browser_instruments.AnalogGauge.speedToDeg = function(x) {
    return (
	    -2.2126e-10 * Math.pow(x,6) + 1.1855e-7 * Math.pow(x,5) +
	    -2.218e-5 * Math.pow(x,4) + 1.5602e-3 * Math.pow(x,3) +
	    -1.6063e-2 * Math.pow(x,2) + 7.2573e-1 * x +
	    5.9149e-1);
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
		var degree = browser_instruments.AnalogGauge.speedToDeg(gaugeValue);
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
