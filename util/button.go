package util

import (
	"machine"
	"time"
)

// ButtonClick is pull up button click event
func ButtonClick(button machine.Pin, callback func()) {
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	go func() {
		buttenState := true
		lastButtonState := true
		debounceDelay := time.Millisecond * 50
		for {
			state := button.Get()
			if state != lastButtonState {
				time.Sleep(debounceDelay)
				state := button.Get()
				if state != buttenState {
					buttenState = state
				}
				if buttenState == false {
					callback()
				}
			}
			lastButtonState = state
			time.Sleep(debounceDelay)
		}
	}()
}
