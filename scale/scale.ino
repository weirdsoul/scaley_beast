/*
 * The scaley beast
 Arduino pin to HX711 map
 2 -> HX711 CLK
 3 -> DOUT
 5V -> VCC
 GND -> GND
*/
 
#include "HX711.h"
 
#define DOUT  3
#define CLK  2
#define LED 4
 
HX711 scale(DOUT, CLK);
 
float calibration_factor = -23479; // Calibrated using a 1kg weight.

//=============================================================================================
//                         SETUP
//=============================================================================================
void setup() {
  Serial.begin(115200);
  scale.set_scale();
  scale.tare(); //Reset the scale to 0

  // Configure LED
  pinMode(LED, OUTPUT);
}
 
//=============================================================================================
//                         LOOP
//=============================================================================================
void loop() {
 
  scale.set_scale(calibration_factor);
  Serial.print(scale.get_units(5), 3);
  Serial.print("\n");
 
  if(Serial.available())
  {
    char temp = Serial.read();
    if (temp == 't')
      scale.tare();  //Reset the scale to zero
    else if (temp == 'r') {      
      // This is our value for 1 kg and therefore
      // our new scale.
      calibration_factor = scale.get_value(5);
    }
    else if(temp == ',')
        digitalWrite(LED, true);
    else if(temp == '.')
        digitalWrite(LED, false);
  }
}
//=============================================================================================
