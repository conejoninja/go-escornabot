package main

import (
	"machine"

	"github.com/conejoninja/go-escornabot/bot"
	"github.com/conejoninja/go-escornabot/input"

	"tinygo.org/x/drivers/easystepper"
	"tinygo.org/x/drivers/buzzer"
)

func main() {
	//motors := easystepper.NewDual(machine.D5, machine.D7, machine.D6, machine.D8, machine.D9, machine.D11, machine.D10, machine.D12, 200, 60)
	motors := easystepper.NewDual(machine.D0, machine.D7, machine.D1, machine.D9, machine.D10, machine.D12, machine.D11, machine.D13, 200, 60)
	buttons := input.NewAnalogButtons(machine.A1)
	bzrPin := machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr := buzzer.New(bzrPin)

	escornabot := bot.New(&motors, &buttons, &bzr, 542)

	escornabot.Loop()
}
