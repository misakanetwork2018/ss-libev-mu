package ss_manager

import (
	"fmt"
)

func Ping() (string, []byte, error) {
	var (
		recv   string
		buffer []byte
	)

	err := send("ping")

	if err != nil {
		fmt.Println(err.Error())
		return "", buffer, err
	}

	recv, buffer, err = receive()

	if err != nil {
		fmt.Println(err.Error())
		return "", buffer, err
	}

	fmt.Println("Receive from ss-manager: " + recv)

	return recv, buffer, nil
}
