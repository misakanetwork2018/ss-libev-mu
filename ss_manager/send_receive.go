package ss_manager

import (
	"net"
)

var conn *net.UnixConn

func send(data string) error {
	_, err := conn.Write([]byte(data))

	return err
}

func receive() (string, []byte, error) {
	buffer := make([]byte, 20480)

	n, err := conn.Read(buffer)

	return string(buffer[:n]), buffer, err
}
