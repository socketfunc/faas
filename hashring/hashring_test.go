package hashring

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	addrs := []string{
		"192.168.13.1:8080",
		"192.168.13.2:8080",
		"192.168.13.3:8080",
		"192.168.13.4:8080",
		"192.168.13.5:8080",
	}
	ring := New(addrs)
	fmt.Println(ring.GetNode("test"))

	addrs = []string{
		"192.168.13.10:8080",
		"192.168.13.2:8080",
		"192.168.13.3:8080",
		"192.168.13.4:8080",
		"192.168.13.5:8080",
	}
	diff := ring.Update(addrs)
	fmt.Println(diff)

	fmt.Println(ring.GetNode("test"))
}
