package piserial

import (
	"log"

	"go.bug.st/serial"
)

func Init() serial.Port {
	mode := &serial.Mode{BaudRate: 4800}
	port, err := serial.Open("/dev/ttyAMA0", mode)
	if err != nil {
		log.Fatal(err)
	}

	return port
}
