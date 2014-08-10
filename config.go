package ahum

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
