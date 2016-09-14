/* vim:set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab */
package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
)
type Config struct {
	SSH                  string // SSH client executable
}

var CONF=Config{}

func main() {
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.StringVar(&CONF.SSH, "ssh", "ssh", "path to SSH client utility.")
	flags.Parse(os.Args[1:])
	switch cmd := flags.Arg(0); cmd {
	case "ssh":
		ssh()
	case "help":
	    usageLong(flags)
	case "":
		usageShort()
	default:
	    fmt.Printf("unknown cmd is %s\n",cmd);
	}
	os.Exit(1)
}

func ssh(){
	fmt.Printf("ssh options is [%s] \n ",CONF.SSH);
}

func usageShort() {
	binName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage: %s [<options>] {help|ssh} [<args>]\n", binName)
}

func usageLong(flags *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, `Usage: %s [<options>] <command> [<args>]

MyBox2 utility.

Commands:
   ssh [ssh-command]   Login to Mybox VM via SSH.
   version             Display version information.

Options:
`, os.Args[0])
	flags.PrintDefaults()
}
