package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 8080
	if len(os.Args) > 1 {
		if p, e := strconv.Atoi(os.Args[1]); e != nil {
			panic(e)
		} else {
			port = p
		}
	}

	listener, e := net.Listen("tcp", ":"+strconv.Itoa(port))
	noError("listen", e)

	defer func() {
		noError("close", listener.Close())
	}()

	buffer := make([]byte, 1024)
	output := []byte(
		"HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 12\r\n" +
			"\r\n" +
			"Hello, World\r\n")

	for {
		conn, e := listener.Accept()
		noError("accept", e)

		n, e := conn.Read(buffer)
		noError("read", e)

		os.Stdout.Write(buffer)
		if n > 4 && buffer[n-4] == '\r' && buffer[n-3] == '\n' &&
			buffer[n-2] == '\r' && buffer[n-1] == '\n' {
			_, e := conn.Write(output)
			noError("write", e)
		}
	}
}

func noError(step string, e error) {
	if e != nil {
		fmt.Println(step, e)
	}
}
