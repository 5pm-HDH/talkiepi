package talkiepi

import (
	"crypto/tls"

	"github.com/5pm-HDH/gumble/gumble"
	"github.com/5pm-HDH/gumble/gumbleopenal"
	"github.com/dchote/gpio"
)

// Raspberry Pi GPIO pin assignments (CPU pin definitions)
const (
	OnlineLEDPin       uint = 5
	ParticipantsLEDPin uint = 6
	TransmitLEDPin     uint = 26
	ButtonPin          uint = 2
)

type Talkiepi struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	ConnectAttempts uint

	Stream *gumbleopenal.Stream

	ChannelName    string
	IsConnected    bool
	IsTransmitting bool

	GPIOEnabled     bool
	OnlineLED       gpio.Pin
	ParticipantsLED gpio.Pin
	TransmitLED     gpio.Pin
	Button          gpio.Pin
	ButtonState     uint
}
