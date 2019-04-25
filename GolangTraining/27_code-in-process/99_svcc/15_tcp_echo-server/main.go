package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// NewScanner returns a new Scanner to read from r.
	// The split function defaults to ScanLines.
	scanner := bufio.NewScanner(conn)
	// Scan advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method.
	for scanner.Scan() {
		// Text returns the most recent token generated by a call to Scan
		// as a newly allocated string holding its bytes.
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Printf("TYPE: %T\n", ln)
		ln = fmt.Sprint("FROM SERVER: " + ln)
		fmt.Fprintln(conn, ln)
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
}

/*
start waitGroup.go (go run waitGroup.go) then ...
telnet localhost 9000
*/
