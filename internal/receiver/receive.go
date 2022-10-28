package receiver

import (
	"io"

	"github.com/rakiyoshi/raspi-serial-test/pkg/piserial"
	"go.bug.st/serial"
)

type Receiver struct {
	port serial.Port
	io.Reader
}

func (r Receiver) Read(p []byte) (int, error) {
	return r.port.Read(p)
}

func NewReceiver() io.Reader {
	return io.Reader(piserial.Init())
}
