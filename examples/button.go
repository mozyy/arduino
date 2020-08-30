package examples

import (
	"machine"

	"../util"
)

// Button 使用按钮控制灯
func Button() {
	led := machine.D10
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// led2 := machine.D12
	// led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// button := machine.D8

	state := false

	util.ButtonClick(machine.D8, func() {
		state = !state
		led.Set(state)
	})

	led2()

}

func led2() {
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(400)
		led.Low()
		delay(400)
	}
}
