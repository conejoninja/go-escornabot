package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/easystepper"
)

func main() {
	motors := easystepper.NewDual(machine.D5, machine.D7, machine.D6, machine.D8, machine.D9, machine.D11, machine.D10, machine.D12, 200, 60)

	for {
		println("CLOCKWISE")
		motors.Move(200,200)
		motors.Off()
		time.Sleep(time.Millisecond * 1000)

		println("COUNTERCLOCKWISE")
		motors.Move(-200,-200)
		motors.Off()
		time.Sleep(time.Millisecond * 1000)

		println("COUNTERCLOCKWISE")
		motors.Move(200,-200)
		motors.Off()
		time.Sleep(time.Millisecond * 1000)

		println("COUNTERCLOCKWISE")
		motors.Move(-200,200)
		motors.Off()
		time.Sleep(time.Millisecond * 1000)
	}
}
