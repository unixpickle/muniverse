// Command netwait waits for an incoming TCP connection,
// then exits when the TCP connection ends.
//
// This is used in muniverse to delete old containers even
// if the muniverse client dies unexpectedly.
// Why does this work? Because the OS automatically cleans
// up a process's sockets when it dies.
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var timeout time.Duration
	var port int
	flag.DurationVar(&timeout, "timeout", 0, "max time to wait for a connection")
	flag.IntVar(&port, "port", 1337, "port on which to listen")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		die(err)
	}
	defer listener.Close()

	gotConn := make(chan struct{})
	if timeout != 0 {
		go func() {
			time.Sleep(timeout)
			select {
			case <-gotConn:
			default:
				die("listen timeout exceeded")
			}
		}()
	}

	conn, err := listener.Accept()
	if err != nil {
		die(err)
	}
	defer conn.Close()
	close(gotConn)

	if _, err := io.Copy(os.Stdout, conn); err != nil {
		die(err)
	}
}

func die(arg interface{}) {
	fmt.Fprintln(os.Stderr, "netwait:", arg)
	os.Exit(1)
}
