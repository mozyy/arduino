package examples

import (
	"machine"
	"time"
)

func Buzzer() {
	// speaker := machine.PA30
	// speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// speaker.Set(true)
	machine.InitPWM()
	bzrPin := machine.PWM{Pin: machine.D11}
	bzrPin.Configure()

	for {
		for i := 0; i < 65535; i += 100 {
			bzrPin.Set(uint16(i))
			time.Sleep(time.Millisecond * 20)
		}
	}
}
