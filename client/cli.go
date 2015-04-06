package client

import (
	"flag"
	"fmt"
	"os"
	// "tunnel/version"
)

const usage1 string = `Usage: %s [OPTIONS] <local port or address>
Options:
`

const usage2 string = `
Examples:
	tunnel 80
	tunnel -subdomain=example 8080
	

Advanced usage: [OPTIONS] <command> [command args] [...]
Commands:
	tunnel start [tunnel] [...]    Start tunnels by name from config file
	
Examples:
	tunnel start wordpress-site railsapp nodeapp 
	
`

type Options struct {
	config    string
	logto     string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}

	config := flag.String(
		"config",
		"",
		"Path to tunnel configuration file. (default: $HOME/.tunnel)")

	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")

	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an ezzytunnel.com account")

	httpauth := flag.String(
		"httpauth",
		"",
		"username:password HTTP basic auth creds protecting the public tunnel endpoint")

	subdomain := flag.String(
		"subdomain",
		"",
		"Request a custom subdomain from the ezzytunnel server. (HTTP only)")

	hostname := flag.String(
		"hostname",
		"",
		"Request a custom hostname from the ezzytunnel server. (HTTP only) (requires you change DNS entries and point CNAME to us)")

	protocol := flag.String(
		"proto",
		"http+https",
		"The protocol of the traffic over the tunnel {'http', 'https', 'tcp'} (default: 'http+https')")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		httpauth:  *httpauth,
		subdomain: *subdomain,
		protocol:  *protocol,
		authtoken: *authtoken,
		hostname:  *hostname,
		command:   flag.Arg(0),
	}

	switch opts.command {
	case "start":
		opts.args = flag.Args()[1:]
	// case "version":
	// 	fmt.Println(version.MajorMinor())
	// 	os.Exit(0)
	// case "help":
	// 	flag.Usage()
	// 	os.Exit(0)
	case "":
		err = fmt.Errorf("Error: Specify a local port to establish tunnel to." +
			"\n\nExample: To expose you localhost port 3000, run " +
			"'tunnel 3000'")
		return

	default:
		if len(flag.Args()) > 1 {
			err = fmt.Errorf("You may only specify one port to tunnel to on the command line, got %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
	}

	return
}
