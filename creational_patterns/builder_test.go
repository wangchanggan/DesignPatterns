package creational_patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestBuilder(t *testing.T) {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(1024).
		WithTimeOut(30 * time.Second).
		Build()

	fmt.Println(server)
}
