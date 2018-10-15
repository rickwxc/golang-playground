package main

import (
    "net"
	"bytes"
	"bufio"
    "os"
    "fmt"
)

type Channel struct {
	id int
	conn net.Conn
}

var channels map[string]*Channel


func main() {

	const port = ":8081"
	fmt.Println("Launching chatting server at port", port)

	// listen on all interfaces
	listener, err := net.Listen("tcp", port)
	checkError(err)

	channels = make(map[string]*Channel)

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

	size := len(channels) + 1
	msize := fmt.Sprintf("%v", size);

	channels[msize] = &Channel{
		size,
		conn,
	}

    var buf [512]byte
    for {
        // read upto 512 bytes
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }

		var s = string(buf[0:])
		fmt.Println("From client: ", s, n)
		boardcast(size, s)
    }

}
func boardcast(chan_id int, s string){
	var b bytes.Buffer

	for id, channel := range channels {
		fmt.Println("Key:", id, "Value:", channel)
		b.WriteString(fmt.Sprintf("%v, %v\n", id, channel))
	}

	for _, channel := range channels {
		w := bufio.NewWriter(channel.conn)
		fmt.Fprint(w, fmt.Sprintf("from %v: %v-->\n", chan_id, s))
		//fmt.Fprint(w, b.String())
		w.Flush()
	}

}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

