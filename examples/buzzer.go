package examples

import (
	"machine"
	"time"
)

const (
	// 定义					 音名  简谱(.在左表示低一个八度, 在右表示高一个八度)
 	NOTE_B0 =   31	// B0 ....7
	
 	NOTE_C1 =   33	// C1 ...1
 	NOTE_CS1 =  35	// C1#
 	NOTE_D1 =   37	// D1 ...2
 	NOTE_DS1 =  39	// D1#
 	NOTE_E1 =   41	// E1 ...3
 	NOTE_F1 =   44	// F1 ...4
 	NOTE_FS1 =  46	// F1#
 	NOTE_G1 =   49	// G1 ...5
 	NOTE_GS1 =  52	// G1#
 	NOTE_A1 =   55	// A1 ...6
 	NOTE_AS1 =  58	// A1#
 	NOTE_B1 =   62	// B1 ...7

 	NOTE_C2 =   65	// C2 ..1
 	NOTE_CS2 =  69	// C2#
 	NOTE_D2 =   73	// D2 ..2
 	NOTE_DS2 =  78	// D2#
 	NOTE_E2 =   82	// E2 ..3
 	NOTE_F2 =   87	// F2 ..4
 	NOTE_FS2 =  93	// F2#
 	NOTE_G2 =   98	// G2 ..5
 	NOTE_GS2 =  104	// G2#
 	NOTE_A2 =   110	// A2 ..6
 	NOTE_AS2 =  117	// A2#
 	NOTE_B2 =   123	// B2 ..7

 	NOTE_C3 =   131	// C3 .1
 	NOTE_CS3 =  139	// C3#
 	NOTE_D3 =   147	// D3 .2
 	NOTE_DS3 =  156	// D3#
 	NOTE_E3 =   165	// E3 .3
 	NOTE_F3 =   175	// F3 .4
 	NOTE_FS3 =  185	// F3#
 	NOTE_G3 =   196	// G3 .5
 	NOTE_GS3 =  208	// G3#
 	NOTE_A3 =   220	// A3 .6
 	NOTE_AS3 =  233	// A3#
 	NOTE_B3 =   247	// B3 .7

 	NOTE_C4 =   262	// C4(中央C) 1
 	NOTE_CS4 =  277	// C4#
 	NOTE_D4 =   294	// D4 2
 	NOTE_DS4 =  311	// D4#
 	NOTE_E4 =   330	// E4 3
 	NOTE_F4 =   349	// F4 4
 	NOTE_FS4 =  370	// F4#
 	NOTE_G4 =   392	// G4 5
 	NOTE_GS4 =  415	// G4#
 	NOTE_A4 =   440	// A4 6
 	NOTE_AS4 =  466	// A4#
 	NOTE_B4 =   494	// B4 7

 	NOTE_C5 =   523	// C5 1.
 	NOTE_CS5 =  554	// C5#
 	NOTE_D5 =   587	// D5 2.
 	NOTE_DS5 =  622	// D5#
 	NOTE_E5 =   659	// E5 3.
 	NOTE_F5 =   698	// F5 4.
 	NOTE_FS5 =  740	// F5#
 	NOTE_G5 =   784	// G5 5.
 	NOTE_GS5 =  831	// G5#
 	NOTE_A5 =   880	// A5 6.
 	NOTE_AS5 =  932	// A5#
 	NOTE_B5 =   988	// B5 7.

 	NOTE_C6 =   1047	// C6 1..
 	NOTE_CS6 =  1109	// C6#
 	NOTE_D6 =   1175	// D6 2..
 	NOTE_DS6 =  1245	// D6#
 	NOTE_E6 =   1319	// E6 3..
 	NOTE_F6 =   1397	// F6 4..
 	NOTE_FS6 =  1480	// F6#
 	NOTE_G6 =   1568	// G6 5..
 	NOTE_GS6 =  1661	// G6#
 	NOTE_A6 =   1760	// A6 6..
 	NOTE_AS6 =  1865	// A6#
 	NOTE_B6 =   1976	// B6 7..

 	NOTE_C7 =   2093	// C7 1...
 	NOTE_CS7 =  2217	// C7#
 	NOTE_D7 =   2349	// D7 2...
 	NOTE_DS7 =  2489	// D7#
 	NOTE_E7 =   2637	// E7 3...
 	NOTE_F7 =   2794	// F7 4...
 	NOTE_FS7 =  2960	// F7#
 	NOTE_G7 =   3136	// G7 5...
 	NOTE_GS7 =  3322	// G7#
 	NOTE_A7 =   3520	// A7 6...
 	NOTE_AS7 =  3729	// A7#
 	NOTE_B7 =   3951	// B7 7...

 	NOTE_C8 =   4186	// C8 1....
 	NOTE_CS8 =  4435	// C8#
 	NOTE_D8 =   4699	// D8 2....
 	NOTE_DS8 =  4978	// D8#
)	

type Tune struct {
	tune uint16
	durt int
}

func   Buzzer() {
	// speaker := machine.PA30
	// speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// speaker.Set(true)
	machine.InitPWM( )
	bzrPin :=    machine.PWM{Pin: machine.D11}
	bzrPin.Configure()
	// 生日快乐
	tunes := []Tune {
		{NOTE_G4,8}, {NOTE_G4, 8},{NOTE_A4,4},{NOTE_G4,4},{NOTE_C5, 4},{NOTE_B4,4},{0,4},
		{NOTE_G4,8}, {NOTE_G4, 8},{NOTE_A4,4},{NOTE_G4,4},{NOTE_D5, 4},{NOTE_C5,4},{0,4},
		{NOTE_G4,8}, {NOTE_G4, 8},{NOTE_G5,4},{NOTE_E5,4},{NOTE_C5, 4},{NOTE_B4,4},{NOTE_A4,2},{0,8},
		{NOTE_F5,8}, {NOTE_F5, 8},{NOTE_E5,4},{NOTE_C5,4},{NOTE_D5, 4},{NOTE_C5,2},{0,4},
	}

	for {
		for _,tune:= range tunes {
			bzrPin.Set(tune.tune)
			//e.g. quarter note = 1000 / 4, eighth note = 1000/8, etc.
			noteDuration := 1000 / tune.durt
			time.Sleep(time.Millisecond * time.Duration(noteDuration))
			bzrPin.Set(0)
			pauseDuration := noteDuration / 3
			time.Sleep(time.Millisecond * time.Duration(pauseDuration))
		}
		time.Sleep(time.Second)
	}
}
