package main

import (
	"fmt"

	"github.com/surge/mqtt"
)

type MessageType byte

func main() {
	msg := mqtt.NewConnectMessage()
	fmt.Println(msg)
}
