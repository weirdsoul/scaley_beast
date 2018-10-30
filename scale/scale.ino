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

float previous_weights[5]; // The last weight values.
float rolling_weight = 0; // Rolling average of the weight.

//=============================================================================================
//                         SETUP
//=============================================================================================
void setup() {
  Serial.begin(115200);
  scale.set_scale();
  // Artificial delay before determining the zero point.
  // This seems to be necessary as the value fluctuates shortly
  // after startup.
  delay(1000);
  scale.tare(10); //Reset the scale to 0

  rolling_weight = 0;
  for (int i = 0; i < 5; ++i) {
    previous_weights[i] = 0;
  }
  
  // Configure LED
  pinMode(LED, OUTPUT);
}
 
//=============================================================================================
//                         LOOP
//=============================================================================================
void loop() {
 
  scale.set_scale(calibration_factor);
  Serial.print(rolling_weight / 5, 3);
  Serial.print("\n");

  // Update rolling weight.
  rolling_weight -= previous_weights[0];
  for (int i = 1; i < 5; ++i) {
    previous_weights[i - 1] = previous_weights[i];
  }
  
  previous_weights[4] = scale.get_units(1);
  rolling_weight += previous_weights[4];
  
  if(Serial.available())
  {
    char temp = Serial.read();
    if (temp == 't')
      scale.tare(10);  //Reset the scale to zero
    else if (temp == 'r') {      
      // This is our value for 1 kg and therefore
      // our new scale.
      calibration_factor = scale.get_value(10);
    }
    else if(temp == ',')
        digitalWrite(LED, true);
    else if(temp == '.')
        digitalWrite(LED, false);
  }
}
//=============================================================================================
