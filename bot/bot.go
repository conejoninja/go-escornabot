package bot

import (
	"time"

	"github.com/conejoninja/go-escornabot/input"
	"tinygo.org/x/drivers/easystepper"
	"tinygo.org/x/drivers/buzzer"
)

const (
	FORWARD  Direction = 0
	BACKWARD Direction = 1
	LEFT     Direction = 2
	RIGHT    Direction = 3
	STOP     Direction = 4
)

const (
	IDLE   Status = 0
	INPUT  Status = 1
	MOVING  Status      = 2
)

type Direction uint8
type Status uint8

type Bot struct {
	motors     *easystepper.DualDevice
	moveSteps  int32
	input      input.Inputer
	bzr *buzzer.Device
	status     Status
	memory     [20]Direction
	memoryStep uint8
}

func New(motors *easystepper.DualDevice, input input.Inputer,bzr *buzzer.Device, steps int32) *Bot {
	return &Bot{
		motors:    motors,
		input:     input,
		bzr:bzr,
		moveSteps: steps,
		status:    INPUT,
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

func (b *Bot) Loop() {
	then := time.Now()
	for {
		switch b.status {
		case IDLE:
			break
		case INPUT:
			for b.memoryStep = 0;b.memoryStep<20;b.memoryStep++ {
				b.memory[b.memoryStep] = STOP
			}
			b.memoryStep = 0
			for b.status == INPUT {
				button := b.GetInput()
				if button != input.NONE {
					b.memory[b.memoryStep] = Direction(button)
					if b.memory[b.memoryStep] == STOP {
						b.status = MOVING
					}
					b.memoryStep++
					then = time.Now()
				}
				if  b.memoryStep > 0 && time.Since(then) > 10*time.Second {
					b.memory[b.memoryStep] = STOP
					b.status = MOVING
				}
				time.Sleep(100 * time.Millisecond)
			}
			break
		case MOVING:
			for b.memoryStep = 0; b.memoryStep < 20; b.memoryStep++ {
				switch b.memory[b.memoryStep] {
				case FORWARD:
					b.Forward()
					break
				case BACKWARD:
					b.Backward()
					break
				case LEFT:
					b.Left()
					break
				case RIGHT:
					b.Right()
					break
				case STOP:
					b.memoryStep = 20
					break
				}
				b.Beep()
				time.Sleep(100*time.Millisecond)
			}

			b.HappySound()
			b.status = INPUT

			break
		}
	}
}

func (b *Bot) Beep() {
	b.bzr.Tone(buzzer.G4, 0.25)
	time.Sleep(30 * time.Millisecond)

}

func (b *Bot) HappySound() {
	b.bzr.Tone(buzzer.G4, 0.25)
	time.Sleep(30 * time.Millisecond)
	b.bzr.Tone(buzzer.G4, 0.5)
	time.Sleep(100 * time.Millisecond)

}
