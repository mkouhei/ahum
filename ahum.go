package ahum

import (
	"io"

	"github.com/surge/mqtt"
)

type MessageType byte

func Conn() (io.Reader, int, error) {
	m, err := mqtt.CONNECT.New()
	if err != nil {
		return nil, -1, err
	}
	msg := m.(*mqtt.ConnectMessage)
	msg.SetWillQos(1)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("ahummq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("guest"))
	msg.SetPassword([]byte("passw0rd"))

	r, n, err := msg.Encode()
	if err != nil {
		return nil, -1, err
	}
	return r, n, nil
}
