package examples

import (
	"machine"
	"time"
)

// Potentiometer 旋转电位器控制led灯
func Potentiometer() {
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

	machine.InitADC()
	ldr := machine.ADC{machine.ADC0}
	ldr.Configure()

	for {
		n := int(float64(ldr.Get()) / 65535 * float64(len(leds)))
		println(ldr.Get(), n, uint16(len(leds)))
		for i, led := range leds {
			led.Set(i < n)
		}

		time.Sleep(time.Millisecond * 100)
	}

}
