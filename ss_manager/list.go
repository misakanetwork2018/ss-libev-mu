package ss_manager

import "fmt"

func List() (string, []byte, error) {
	var (
		recv   string
		buffer []byte
	)

	err := send("list")

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
