package examples

import "machine"

func Uart() {
	uart := machine.UART0
	uart.Configure(machine.UARTConfig{TX: machine.UART_TX_PIN,RX: machine.UART_RX_PIN})

	uart.Write([]byte("b byte\n"))

}