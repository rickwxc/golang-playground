/* ThreadedEchoServer
 */
package main

import (
    "net"
    "os"
    "fmt"
)

func main() {

	const port = ":8081"
	fmt.Println("Launching multi-thread echo server at port", port)

	// listen on all interfaces
	listener, err := net.Listen("tcp", port)
	checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        // run as a goroutine
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    // close connection on exit
    defer conn.Close()

    var buf [512]byte
    for {
        // read upto 512 bytes
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }

		fmt.Println("From client: ", string(buf[0:]))
        // write the n bytes read
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return
        }
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
