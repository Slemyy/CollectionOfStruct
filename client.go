package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln("Error:", err.Error())
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nText to send: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprint(conn, text)
		request, _ := bufio.NewReader(conn).ReadString('\n')

		if request == "" {
			fmt.Println("The connection to the server has been lost.")
			break
		} else if request == "The connection is closed.\n" {
			fmt.Print(request)
			break
		}

		fmt.Print(request)
	}
}
