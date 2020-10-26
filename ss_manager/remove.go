package ss_manager

import (
	"fmt"
	"strconv"
)

func Remove(port int) (bool, error) {
	var recv string

	err := send("remove: {\"server_port\": " + strconv.Itoa(port) + "}")

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
