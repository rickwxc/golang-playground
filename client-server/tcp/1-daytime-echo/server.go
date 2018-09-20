/* DaytimeServer */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	const port = ":8080"
	fmt.Println("Launching daytime echo server at port", port)

	// listen on all interfaces
	ln, err := net.Listen("tcp", port)
	checkError(err)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		fmt.Println("echo: ", daytime)
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
