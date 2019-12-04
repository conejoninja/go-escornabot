package bot

import (
	"github.com/conejoninja/go-escornabot/input"
	"tinygo.org/x/drivers/easystepper"
)

const (
	FORWARD  Direction = 0
	BACKWARD Direction = 1
	LEFT     Direction = 2
	RIGHT    Direction = 3
)

type Direction uint8

type Bot struct {
	motors    *easystepper.DualDevice
	moveSteps int32
	input     input.Inputer
}

func New(motors *easystepper.DualDevice, input input.Inputer, steps int32) *Bot {
	return &Bot{
		motors:    motors,
		input:     input,
		moveSteps: steps,
	}
}

func (b *Bot) GetInput() input.Button {
	return b.input.Get()
}

func (b *Bot) Move(direction Direction) {
	switch direction {
	case FORWARD:
		b.motors.Move(b.moveSteps, -b.moveSteps)
		break
	case BACKWARD:
		b.motors.Move(-b.moveSteps, b.moveSteps)
		break
	case LEFT:
		b.motors.Move(-b.moveSteps, -b.moveSteps)
		break
	case RIGHT:
		b.motors.Move(b.moveSteps, b.moveSteps)
		break
	}
	b.motors.Off()
}

func (b *Bot) Forward() {
	b.Move(FORWARD)
}

func (b *Bot) Backward() {
	b.Move(BACKWARD)
}

func (b *Bot) Left() {
	b.Move(LEFT)
}

func (b *Bot) Right() {
	b.Move(RIGHT)
}
