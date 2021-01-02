package scan

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	ip      string
	lock    *semaphore.Weighted
	threads int
}

type connection struct {
	IP      string
	Port    int
	Timeout time.Duration
}

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func ScanPort(ip string, port int, timeout time.Duration, wg *sync.WaitGroup) {

	/* that was debug stuff to test multithreading success
	fmt.Println(ip, port, timeout)
	time.Sleep(500 * time.Millisecond)
	wg.Done()

	return*/
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout, wg)
		} else {
			//fmt.Println(port, "closed")
			wg.Done()
		}
		return
	}

	conn.Close()

	fmt.Printf("%d \t %s \n", port, knownPorts[port])
	wg.Done()
}

func Do(ip string, start int, end int, timeout int, threads int) {
	ps := &PortScanner{
		ip:      ip,
		lock:    semaphore.NewWeighted(Ulimit()),
		threads: threads}

	ps.Split(start, end, time.Duration(timeout)*time.Millisecond)
}

func Run(connections chan connection, wg *sync.WaitGroup) {
	for el := range connections {
		ScanPort(el.IP, el.Port, el.Timeout, wg)
	}
}

func (ps *PortScanner) Split(start, end int, timeout time.Duration) {
	connections := make(chan connection, ps.threads)

	wg := sync.WaitGroup{}

	for i := 0; i < ps.threads; i++ {
		go Run(connections, &wg)
	}

	for port := start; port <= end; port++ {
		wg.Add(1)
		go func(port int) {
			con := connection{IP: ps.ip, Port: port, Timeout: timeout}
			//fmt.Println(port, " do")
			connections <- con
		}(port)

	}

	//close(connections) //closing here will throw wn error
	wg.Wait()
}
