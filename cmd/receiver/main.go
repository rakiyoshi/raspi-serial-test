package main

import (
	"bufio"
	"fmt"

	"github.com/rakiyoshi/raspi-serial-test/internal/receiver"
)

func main() {
	r := bufio.NewScanner(receiver.NewReceiver())
	for r.Scan() {
		fmt.Println(r.Text())
	}
}
