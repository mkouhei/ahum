package ahum

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"text/template"
)

const testConfig = `[global]
host={{.Host}}
port={{.Port}}
clientid={{.ClientId}}
username={{.Username}}
password={{.Password}}
willQoS={{.WillQos}}
version={{.Version}}
cleanSession={{.CleanSession}}
keepalive={{.KeepAlive}}
`

var (
	testConf = "test.conf"
)

type data struct {
	Host         string
	Port         string
	ClientId     string
	Username     string
	Password     string
	WillQos      string
	Version      string
	CleanSession string
	KeepAlive    string
}

func TestReadConfig(t *testing.T) {
	if c, err := readConfig(""); err == nil || c != nil {
		t.Fatalf("%v, want: %v", err, c)
	}

	v := &data{
		"localhost",
		"1883",
		"ahummq",
		"guest",
		"passw0rd",
		"1",
		"3",
		"true",
		"10",
	}
	configWrite(*v)
	if c, err := readConfig(testConf); err != nil {
		t.Fatalf("%v, want: %v", err, c)
	}
	os.Remove(testConf)
}

func configWrite(v data) {
	testTmpl := template.Must(template.New("").Parse(testConfig))
	var w bytes.Buffer
	testTmpl.Execute(&w, &v)
	ioutil.WriteFile(testConf, w.Bytes(), 0644)
}
