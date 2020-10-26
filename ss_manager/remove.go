package ss_manager

import (
	"fmt"
)

func Remove(port string) (bool, error) {
	var recv string

	err := send("remove: {\"server_port\": " + port + "}")

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	recv, _, err = receive()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	fmt.Println("Receive from ss-manager: " + recv)

	return recv == "ok", nil
}
