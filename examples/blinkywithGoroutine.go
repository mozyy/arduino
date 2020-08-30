package examples

import (
	"machine"
	"time"
)

// BlinkywithGoroutine 用 goroutine 使两个灯闪烁
func BlinkywithGoroutine() {
	led1() // goroutine
	led22()
}

func led1() {
	led := machine.D10
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			led.High()
			delay(500)
			led.Low()
			delay(500)
		}
	}()
}

func led22() {
	led := machine.D12
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(400)
		led.Low()
		delay(400)
	}
}

func delay(t int64) {
	time.Sleep(time.Duration(1000000 * t))
}
