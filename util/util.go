package util

import "time"

// Delay is delay
func Delay(t time.Duration) {
	time.Sleep(time.Millisecond * t)
}
