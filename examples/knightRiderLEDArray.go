package examples

import (
	"machine"
	"time"
)

// KnightRiderLEDArray 跑马灯
func KnightRiderLEDArray() {
	leds := []machine.Pin{
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
		machine.D6,
	}
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	index, delta := 0, 1
	for {
		for i, led := range leds {
			led.Set(i == index)
		}
		index += delta
		if index == 0 || index == len(leds)-1 {
			delta *= -1
		}

		time.Sleep(time.Millisecond * 75)
	}

}
