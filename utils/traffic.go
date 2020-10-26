package utils

import (
	"encoding/json"
	"fmt"
	"ss-libev-mu/ss_manager"
)

func GetStats() (map[string]int, error) {
	var stats map[string]int

	str, _, err := ss_manager.Ping()

	if err != nil {
		fmt.Println("get stat error: ", err.Error())
		return nil, err
	}

	str = str[6:len(str)]

	fmt.Println("final json: ", str)

	if err = json.Unmarshal([]byte(str), &stats); err != nil {
		fmt.Println("parse json error: ", err.Error())
		return nil, err
	}

	return stats, nil
}
