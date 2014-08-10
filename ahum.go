package ahum

import (
	"fmt"
	"io"
	"net"

	"github.com/surge/mqtt"
)

type MessageType byte

func (c *config) Conn() (int64, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	msg := mqtt.NewConnectMessage()
	msg.SetWillQos(c.WillQos)
	msg.SetVersion(c.Version)
	msg.SetCleanSession(c.CleanSession)
	msg.SetClientId([]byte(c.ClientId))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte(c.Username))
	msg.SetPassword([]byte(c.Password))

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
