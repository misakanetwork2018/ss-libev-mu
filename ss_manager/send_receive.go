package ss_manager

import (
	"io"
	"net"
)

var conn *net.UnixConn

func send(data string) error {
	_, err := conn.Write([]byte(data))

	return err
}

func receive() (string, []byte, error) {
	buffer := make([]byte, 2048)
	len := 0

	for {
		n, err := conn.Read(buffer)
		if n > 0 {
			len += n
		}
		if err != nil {
			if err != io.EOF {
				return "", buffer, err
			}

			break
		}
	}

	return string(buffer[:len]), buffer, nil
}
