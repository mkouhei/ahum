package ahum

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	r, n, err := Conn()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
	fmt.Println(n)
}
