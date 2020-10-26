package ss_manager

import (
	"fmt"
	"net"
	"os"
)

func Init(managerAddress string, bindAddress string) {
	var err error

	err = os.Remove(bindAddress)

	if err != nil {
		fmt.Println("delete sock failed: ", err.Error())
		return
	}

	rUnixAddr, _ := net.ResolveUnixAddr("unixgram", managerAddress)
	lUnixAddr, _ := net.ResolveUnixAddr("unixgram", bindAddress)

	conn, err = net.DialUnix("unixgram", lUnixAddr, rUnixAddr)

	if err != nil {
		fmt.Println("Error dialing: ", err.Error())
		return
	}
}
