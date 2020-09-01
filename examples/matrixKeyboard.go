package examples

import (
	"machine"
	"time"
)

type note struct {
	tone     float64
	duration float64
}

// MatrixKeyboard 矩阵按键
func MatrixKeyboard() {
	colPins := []machine.Pin{
		machine.D6,
		machine.D7,
		machine.D8,
		machine.D9,
	}
	rowPins := []machine.Pin{
		machine.D2,
		machine.D3,
		machine.D4,
		machine.D5,
	}
	// chars := [][]string{
	// 	{"1", "2", "3", "A"},
	// 	{"4", "5", "6", "B"},
	// 	{"7", "8", "9", "C"},
	// 	{"*", "0", "#", "D"},
	// }

	machine.InitPWM()
	bzrPin := machine.PWM{Pin: machine.D10}
	bzrPin.Configure()

	for _, pin := range colPins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.Set(true)
	}
	for _, pin := range rowPins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInput})
		pin.Set(true)
	}
	for {
		for c, col := range colPins {
			col.Set(false)
			for r, row := range rowPins {
				if !row.Get() {
					// println(c, r)
					for !row.Get() {
					}
					// println(chars[r][c])
					// char := chars[r][c]
					// i, err := strconv.Atoi(char)
					// if err != nil {
					// 	break
					// }
					n := float64((c+1)*(r+1)) / float64(16) * float64(65535)
					bzrPin.Set(uint16(n))
					time.Sleep(time.Millisecond * 500)
					bzrPin.Set(0)
					break
				}
			}
			col.Set(true)
		}
	}
}
