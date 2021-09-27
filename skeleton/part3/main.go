// Solution to part 3 of the Whispering Gophers code lab.
//
// This program listens on the host and port specified by the -listen flag.
// For each incoming connection, it launches a goroutine that reads and decodes
// JSON-encoded messages from the connection and prints them to standard
// output.
//
// You can test this program by running it in one terminal:
// 	$ part3 -listen=localhost:8000
// And running part2 in another terminal:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const TCPNetwork = "tcp"

var listenAddr = flag.String("listen", "localhost:8000", "host:port to listen on")

type Message struct {
	Body string
}

func main() {
	flag.Parse()
	listen, err := net.Listen(TCPNetwork, *listenAddr)
	if err != nil {
		log.Fatalf("problem listening from the address %v", err)
	}

	for {
		c, err := listen.Accept()
		if err != nil {
			log.Fatalf("problem accept a new connection from the listener %v", err)
		}

		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	decoder := json.NewDecoder(c)
	for {
		var m Message
		if err := decoder.Decode(&m); err != nil {
			log.Fatalf("problem decode information %v", err)
		}
		fmt.Fprintf(os.Stdout, "%#v", m)
	}
}
