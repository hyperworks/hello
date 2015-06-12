package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const LISTEN_ADDR = ":8080"
const DEFAULT_TEXT = "Hello, World!"

func main() {
	text := DEFAULT_TEXT
	if len(os.Args) > 1 {
		text = strings.Join(os.Args[1:], " ")
	}

	listener, e := net.Listen("tcp", LISTEN_ADDR)
	noError("listen", e)
	defer func() { noError("server close", listener.Close()) }()

	buffer := make([]byte, 1024)
	output := []byte(
		"HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: " + strconv.Itoa(len(text)) + "\r\n" +
			"\r\n" +
			text + "\r\n")

	for {
		conn, e := listener.Accept()
		noError("accept", e)

		go func() {
			defer func() { noError("close", conn.Close()) }()

			n, e := conn.Read(buffer)
			noError("read", e)

			if n > 4 && buffer[n-4] == '\r' && buffer[n-3] == '\n' &&
				buffer[n-2] == '\r' && buffer[n-1] == '\n' {
				_, e := conn.Write(output)
				noError("write", e)
			}
		}()
	}
}

func noError(step string, e error) {
	if e != nil {
		fmt.Println(step, e)
	}
}
