package main

import (
	"github.com/conejoninja/go-escornabot/bot"
	"github.com/conejoninja/go-escornabot/input"
	"machine"
	"time"

	"tinygo.org/x/drivers/easystepper"
)

func main() {
	//motors := easystepper.NewDual(machine.D5, machine.D7, machine.D6, machine.D8, machine.D9, machine.D11, machine.D10, machine.D12, 200, 60)
	motors := easystepper.NewDual(machine.D0, machine.D7, machine.D1, machine.D9, machine.D10, machine.D12, machine.D11, machine.D13, 200, 60)
	buttons := input.NewAnalogButtons(machine.A0)

	escornabot := bot.New(&motors, &buttons, 400)
	for {
		/*println("FORWARD")
		escornabot.Forward()
		time.Sleep(time.Millisecond * 1000)

		println("BACKWARD")
		escornabot.Backward()
		time.Sleep(time.Millisecond * 1000)

		println("LEFT")
		escornabot.Left()
		time.Sleep(time.Millisecond * 1000)

		println("RIGHT")
		escornabot.Right()
		time.Sleep(time.Millisecond * 1000)*/

		println(escornabot.GetInput())
		time.Sleep(100*time.Millisecond)
	}
}
