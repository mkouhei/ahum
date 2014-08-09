package ahum

type config struct {
	ClientId     string
	Username     string
	Password     string
	WillQos      byte
	Version      byte
	CleanSession bool
	KeepAlive    uint16
}
