package ahum

import (
	"github.com/surge/mqtt"
)

type MessageType byte

func Conn() (*mqtt.ConnectMessage, error) {
	m, err := mqtt.CONNECT.New()
	if err != nil {
		return nil, err
	}
	msg := m.(*mqtt.ConnectMessage)
	return msg, nil
}
