package main

import "./examples"

// tinygo flash -target arduino -port COM3 arduino.go
// tinygo flash -scheduler coroutines -target arduino -port COM3 arduino.go
func main() {
	// examples.BlinkywithGoroutine()
	// examples.Button()
	// examples.KnightRiderLEDArray()
	// examples.Potentiometer()
	// examples.SevenSegmentDisplay()
	// examples.MatrixKeyboard()
	examples.Buzzer()
}
