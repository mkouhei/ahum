package ahum

import (
	"fmt"

	"github.com/miguel-branco/goconfig"
)

const (
	defaultWillQos   = byte(1)
	defaultVersion   = byte(3)
	defaultKeepAlive = uint16(10)
)

type config struct {
	Host         string
	Port         uint32
	ClientId     string
	Username     string
	Password     string
	WillQos      byte
	Version      byte
	CleanSession bool
	KeepAlive    uint16
}

func readConfig(confPath string) (*config, error) {
	conf := &config{}
	c, err := goconfig.ReadConfigFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("readConfig: %v", err)
	}

	conf.Host, err = c.GetString("global", "host")
	if err != nil {
		return nil, fmt.Errorf("readConfig: host: %v", err)
	}
	var p int64
	p, err = c.GetInt64("global", "port")
	if err != nil {
		return nil, fmt.Errorf("readConfig: port: %v", err)
	}
	conf.Port = uint32(p)

	conf.ClientId, err = c.GetString("global", "clientid")
	if err != nil {
		return nil, fmt.Errorf("readConfig: clientid: %v", err)
	}

	conf.Username, err = c.GetString("global", "username")
	if err != nil {
		return nil, fmt.Errorf("readConfig: username: %v", err)
	}

	conf.Password, err = c.GetString("global", "password")
	if err != nil {
		return nil, fmt.Errorf("readConfig: password: %v", err)
	}

	var wq, ver int64
	wq, err = c.GetInt64("global", "willQoS")
	if err != nil {
		conf.WillQos = defaultWillQos
	}
	conf.WillQos = byte(wq)

	ver, err = c.GetInt64("global", "version")
	if err != nil {
		conf.Version = defaultVersion
	}
	conf.Version = byte(ver)

	conf.CleanSession, err = c.GetBool("global", "cleanSession")
	if err != nil {
		conf.CleanSession = true
	}

	var ka int64
	ka, err = c.GetInt64("global", "keepalive")
	if err != nil {
		conf.KeepAlive = defaultKeepAlive
	}
	conf.KeepAlive = uint16(ka)

	return conf, nil
}
