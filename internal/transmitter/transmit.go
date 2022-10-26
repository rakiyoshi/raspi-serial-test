package transmitter

import (
	"fmt"

	"github.com/rakiyoshi/raspi-serial-test/pkg/piserial"
	"go.bug.st/serial"
)

type Transmitter struct {
	port serial.Port
}

func (t Transmitter) Write(p []byte) (n int, err error) {
	n, err = t.port.Write(p)
	if err != nil {
		return -1, fmt.Errorf("transmitter: %v", err)
	}
	return
}

func NewTransmitter() Transmitter {
	return Transmitter{piserial.Init()}
}
