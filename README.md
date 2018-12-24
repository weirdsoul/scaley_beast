# Scaley beast

An Arduino scale with a web interface. Built using an Arduino Nano with a HX711 load cell amplifier, 
four load cells and a piece of wood, this scale was used as part of an escape room game we played at our home 
Halloween 2018. The Arduino needs to be plugged into a PC running a simple webserver written in Go.

The PC software consists of a Go webserver, a few pieces of Javascript code using the Google Closure library and some html
and graphics produced using Gimp. It will render a scale gauge on the website it serves, which gets updates from the Arduino
through a web socket.

I used Bazel as the build system, because it allows to build the whole PC side application with a single command.
Just look into the BUILD file for the respective targets.

The Arduino code is compiled using the Arduino IDE.
