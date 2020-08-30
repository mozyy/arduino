package examples

import (
	"machine"

	"../util"
)

//SevenSegmentDisplay 七段数码管
func SevenSegmentDisplay() {
	numbers := [10]byte{
		0x3f,
		0x06,
		0x5b,
		0x4f,
		0x66,
		0x6d,
		0x7d,
		0x07,
		0x7f,
		0x6f,
	}
	leds := [8]machine.Pin{
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D6,
		machine.D7,
		machine.D8,
		machine.D9,
	}
	// configure leds
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	for {
		for _, num := range numbers {
			for i, led := range leds {
				state := num>>i&1 == 1
				led.Set(!state)
			}
			util.Delay(500)
		}
		for _, led := range leds {
			led.Set(false)
		}
		util.Delay(2000)
	}
}
