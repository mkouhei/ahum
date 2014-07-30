package ahum

import (
	"testing"
)

func TestConn(t *testing.T) {
	msg, err := Conn()
	if err != nil {
		t.Fatal(err)
	}
	if msg.Name == nil {
		t.Fatal(err)
	}

}
