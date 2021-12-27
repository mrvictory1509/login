package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Println(msg)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[: len(nameInput)-1]

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s nói: %s\n", nameInput, msg[: len(msg)-1])
		connection.Write([]byte(msg))
	}

	connection.Close()

}
