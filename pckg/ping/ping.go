package ping

import (
	"fmt"
	"io/ioutil"
	"net"
)

func Ping(ip string, port int, message string) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", target)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, message)

	resp, err := ioutil.ReadAll(conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resp))
}
