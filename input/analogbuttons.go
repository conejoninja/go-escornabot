package input

import "machine"

type AnalogButtons struct {
	pin     machine.ADC
	limits  [10]uint16
	pressed bool
}

func NewAnalogButtons(pin machine.Pin) AnalogButtons {
	machine.InitADC()
	analogPin := machine.ADC{pin}
	analogPin.Configure()

	return AnalogButtons{
		pin:    analogPin,
		limits: [10]uint16{49000, 50000, 32000, 33500, 56000, 57000, 43000, 45000, 52000, 54000},
	}
}

func (ab *AnalogButtons) Get() Button {
	value := ab.pin.Get()
	for i := 0; i < 5; i++ {
		if value >= ab.limits[i*2] && value <= ab.limits[i*2+1] {
			if ab.pressed {
				return NONE
			}
			ab.pressed = true
			return Button(i)
		}
	}
	ab.pressed = false
	return NONE
}
