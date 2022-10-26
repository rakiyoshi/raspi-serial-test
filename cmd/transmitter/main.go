package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rakiyoshi/raspi-serial-test/internal/transmitter"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Started transmitter")

	w := transmitter.NewTransmitter()
	for input.Scan() {
		fmt.Fprintf(w, "%s\n", input.Text())
	}
}
