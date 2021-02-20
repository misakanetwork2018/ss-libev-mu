package ss_manager

import (
	"fmt"
	"net"
	"os"
)

var (
	managerAddress string
	bindAddress    string
)

func send(data string) error {
	_, _, err := exec(func(conn *net.UnixConn) ([]byte, int, error) {
		n, err := conn.Write([]byte(data))
		return nil, n, err
	})

	return err
}

func receive() (string, []byte, error) {
	buffer, n, err := exec(func(conn *net.UnixConn) ([]byte, int, error) {
		buffer := make([]byte, 20480)
		n, err := conn.Read(buffer)
		return buffer, n, err
	})

	return string(buffer[:n]), buffer, err
}

// 由于ss-manage经常断线，而且本地udp基本上没有延迟损耗，所以干脆每次查询都连一次
func exec(customFunc func(conn *net.UnixConn) ([]byte, int, error)) ([]byte, int, error) {
	var err error

	if _, err = os.Stat(bindAddress); err == nil {
		if err = os.Remove(bindAddress); err != nil {
			fmt.Println("delete sock failed: ", err.Error())
			return nil, 0, err
		}
	}

	rUnixAddr, _ := net.ResolveUnixAddr("unixgram", managerAddress)
	lUnixAddr, _ := net.ResolveUnixAddr("unixgram", bindAddress)

	conn, err := net.DialUnix("unixgram", lUnixAddr, rUnixAddr)

	if err != nil {
		fmt.Println("Error dialing: ", err.Error())
		return nil, 0, err
	}

	defer conn.Close()

	return customFunc(conn)
}

func Init(manager string, bind string) {
	managerAddress = manager
	bindAddress = bind
}
