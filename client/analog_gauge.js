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
    var foundElement = true;
    /** @type {!Array<!number>} **/
    this.inputSource = [];
    /** @type {!Array<!number>} **/
    this.inputIndex = [];
    /** @type {!Array<!function(number):number>} **/    
    this.speedToDeg = [];

    for (var sig = 0; foundElement; ++sig) {
	/** @type {string} **/
	var indexSuffix = sig > 0 ? (sig + 1).toString() : "";

	var sourceAttr = "source" + indexSuffix;

	// If we don't have a source we are done.
	if (!goog.dom.dataset.has(domElement, sourceAttr)) {
	    console.log("Attribute " + sourceAttr + " not found, assuming that we are done.");
	    break;
	}

	this.inputSource.push(parseInt(goog.dom.dataset.get(domElement, sourceAttr), 10));

	var indexAttr = "index" + indexSuffix;
	this.inputIndex.push(goog.dom.dataset.has(domElement, indexAttr)
			     ? parseInt(goog.dom.dataset.get(domElement, indexAttr), 10): 0);
	
	var transformAttr = "transform" + indexSuffix;
	if (goog.dom.dataset.has(domElement, transformAttr)) {
	    // Data transform is used to convert the input signal x to degrees.
	    this.speedToDeg.push(new Function("x", goog.dom.dataset.get(domElement, transformAttr)));
	} else {
	    console.log("Using fallback function for input signal conversion.");
	    /**Completely linear. Expects the input to be degrees.
	     * @param {number} x Input to be converted.
	     * @return {number} Resulting degrees for the gauge hand.
	     */	    
	    var idFunction = function(x) { return x; };
	    this.speedToDeg.push(idFunction);
	}
    }
};

/**
 * Called to update an instrument with a new piece of data.
 * @param {!Object} msg A JSON object containing all the update information.
 */
browser_instruments.AnalogGauge.prototype.updateInstrument = function(msg) {
    /** @type {?Array<?Object>} **/
    var valueSet = msg["Values"];
    if (valueSet.length == 0) return;
    for (var sig = 0; sig < this.inputSource.length; ++sig) {
	if (msg["Index"] == this.inputSource[sig]) {
	    /** @type {string} **/
	    var indexSuffix = sig > 0 ? (sig + 1).toString() : "";

	    var needleName = "needle" + indexSuffix;

	    var needleElement = goog.dom.getElementByClass(needleName, this.domElement);
	    if (needleElement != null) {
		var gaugeValue = parseFloat(valueSet[this.inputIndex[sig]]);
		if (gaugeValue != null) {
		    var degree = this.speedToDeg[sig](gaugeValue);
		    // TODO(aeckleder): Do not hardcode center point here.
		    needleElement.setAttribute("transform",
					       "rotate(" + degree + " 240 240)");
		}
	    } else {
		console.error("Couldn't find needle element.");
	    }
	}
    }
};
