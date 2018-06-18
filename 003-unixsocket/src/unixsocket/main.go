package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func handleErr(e error) string {
	if e != nil {
		fmt.Printf("\n Error : %v \n", e)
	} else {
		return ""
	}
	return ""
}

func startServerMode(wg *sync.WaitGroup) {
	fmt.Printf("\n Running Server Mode .. \n")

	buf := make([]byte, 1024)
	// First we need to create a socket file which can listen on incoming connections.
	err := os.Remove("./unixsock") // delete existing socket file
	handleErr(err)
	listener, err := net.Listen("unix", "./unixsock")
	handleErr(err)
	//wg.Add(1)
	func(listener net.Listener) {
		//defer wg.Done()
		for { // this is the part where the echo happens
			conn, err := listener.Accept()
			handleErr(err)
			nr, err := conn.Read(buf) // Read from the connection to buf
			fmt.Printf("\n Got -> %v", string(buf[:nr]))
			handleErr(err)
			//conn.Write(buf[:nr]) // Write back to the connection, hence echo server
		}
	}(listener)
	//wg.Wait()
}

func startClientMode(wg *sync.WaitGroup) {
	fmt.Printf("\n Running Client Mode .. \n")
	//buf := make([]byte, 1024)

	conn, err := net.Dial("unix", "./unixsock")
	handleErr(err)

	//wg.Add(1)
	func(conn net.Conn) {
		//defer wg.Done()
		for {
			_, err := conn.Write([]byte("YOLLLOOOOOO !"))
			handleErr(err)
		}
	}(conn)

	defer conn.Close()
	//wg.Wait()
}

func main() {
	flagMode := flag.String("mode", "server", "Start in Server Mode else CLient Mode")
	var wgServer sync.WaitGroup
	var wgClient sync.WaitGroup

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("\n Please give -mode flag with server or client argument .. \n")
		return
	}

	if strings.ToLower(*flagMode) == "server" {
		startServerMode(&wgServer)
	} else {
		startClientMode(&wgClient)
	}
}
