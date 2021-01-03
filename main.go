package main

import (
	"github.com/DionTech/portscan/pckg/flood"
	"github.com/DionTech/portscan/pckg/scan"
	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help bool `short:"h" long:"help" description:"Display usage" global:"false"`
		Scan struct {
			IP      string `short:"i" long:"ip" description:"define the ip to scan" required:"true" nonempty:"false"`
			Start   int    `short:"s" long:"start" description:"define the start port" required:"false" nonempty:"false"`
			End     int    `short:"e" long:"end" description:"define the end port" required:"false" nonempty:"false"`
			Timeout int    `long:"timeout" description:"define the timeout" required:"false" nonempty:"false"`
			Threads int    `short:"t" long:"threads" description:"amount of threads being used" required:"false" nonempty:"false"`
		} `command:"scan" description:"make a port scan" nonempty:"false"`
		Flood struct {
			LocalIP string `short:"l" long:"local-ip" description:"define the local ip address being used to flood" required:"false" nonempty:"false"`
			IP      string `short:"i" long:"ip" description:"define the ip to flood" required:"true" nonempty:"false"`
			Port    int    `short:"p" long:"port" description:"define the port to flood" required:"true" nonempty:"false"`
			Size    int    `short:"s" long:"size" description:"define the size of how many connections to establish" required:"true" nonempty:"false"`
		} `command:"flood" description:"make a port flooding; be careful what you do" nonempty:"false"`
	}{}

	gocmd.HandleFlag("Flood", func(cmd *gocmd.Cmd, args []string) error {
		localIP := "127.0.4.1"
		if flags.Flood.LocalIP != "" {
			localIP = flags.Flood.LocalIP
		}

		flood.Do(localIP, flags.Flood.IP, flags.Flood.Port, flags.Flood.Size)

		return nil
	})

	gocmd.HandleFlag("Scan", func(cmd *gocmd.Cmd, args []string) error {
		//defining some efault values
		start := 1
		end := 65535
		timeout := 500
		threads := 500

		if flags.Scan.Start != 0 {
			start = flags.Scan.Start
		}

		if flags.Scan.End != 0 {
			end = flags.Scan.End
		}

		if flags.Scan.Threads != 0 {
			threads = flags.Scan.Threads
		}

		scan.Do(flags.Scan.IP, start, end, timeout, threads)

		return nil
	})

	gocmd.New(gocmd.Options{
		Name:        "portscan",
		Version:     "0.1.0",
		Description: "scanning ports done with go",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
