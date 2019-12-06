package input

const (
	FORWARD  Button = 0
	BACKWARD Button = 1
	LEFT     Button = 2
	RIGHT    Button = 3
	OK       Button = 4
	NONE     Button = 5
)

type Button uint8

type Inputer interface {
	Get() Button
}
