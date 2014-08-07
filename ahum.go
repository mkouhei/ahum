package ahum

import (
	"io"
	"net"

	"github.com/surge/mqtt"
)

type MessageType byte

func Conn() (int64, error) {
	conn, err := net.Dial("tcp", "localhost:1883")
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	msg := mqtt.NewConnectMessage()
	msg.SetWillQos(1)
	msg.SetVersion(3)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("ahummq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("guest"))
	msg.SetPassword([]byte("passw0rd"))

	r, n, err := msg.Encode()
	if err != nil {
		return -1, err
	}

	m, err := io.CopyN(conn, r, int64(n))
	if err != nil {
		return -1, err
	}
	return m, nil
}
