package main

import (
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
			Threads int    `short:"t" long:"threads" threads" required:"false" nonempty:"false"`
		} `command:"scan" description:"make a port scan" nonempty:"false"`
	}{}

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
