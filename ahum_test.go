package ahum

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	c := &config{"ahummq", "guest", "passw0rd", 1, 3, true, 10}
	r, err := c.Conn()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
