package ahum

import "testing"

func TestConn(t *testing.T) {
	c := &config{"localhost", 1882, "ahummq", "guest", "passw0rd", 1, 3, true, 10}
	if r, err := c.Conn(); err == nil || r != -1 {
		t.Fatalf("%v, rc: %v, want: dial tcp 127.0.0.1:1882: connection refused", err, r)
	}
	c = &config{"localhost", 1883, "ahummq", "guest", "passw0rd", 1, 3, true, 10}
	if r, err := c.Conn(); err != nil || r != 59 {
		t.Fatalf("%v, rc: %v, want: 59", err, r)
	}
}
