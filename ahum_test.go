package ahum

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	r, err := Conn()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
