package ss_manager

import (
	"encoding/json"
	"fmt"
	"ss-libev-mu/model"
)

func Add(user model.User) (bool, error) {
	var recv string
	jsons, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	err = send("add: " + string(jsons))

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
