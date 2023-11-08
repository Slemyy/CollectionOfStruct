package main

import (
	"CollectionOfStruct/handlers"
	"log"
	"net"
	"strings"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalln("Error:", err.Error())
		return
	}

	defer listener.Close() // При завершении потока main, для закрытия потока.
	log.Println("Server started. Listening on port", listener.Addr().String()[5:]+"...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error:", err.Error())
			return
		}

		var mut sync.Mutex
		go handleClient(conn, &mut)
	}

}

func handleClient(conn net.Conn, mut *sync.Mutex) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr() // Получение адреса удаленного узла

	log.Printf("Connection established with: %s\n", remoteAddr)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Connection to %s is closed.\n", remoteAddr)
			return
		}

		clientMessage := string(buffer[:n])
		if clientMessage == "\r\n" {
			_, err := conn.Write([]byte("Not enough arguments. Use: --file <file.json> --query <query>.\n"))
			if err != nil {
				log.Printf("Error: %v\n", err)
				break
			}

			continue
		}

		log.Printf("Received from %s: %s", remoteAddr, clientMessage)

		args := strings.Fields(clientMessage)
		if args[0] == "exit" {
			log.Printf("Close the connection with the client: %s", remoteAddr)
			_, err := conn.Write([]byte("The connection is closed.\n"))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}
			break
		}

		if len(args) < 4 {
			_, err := conn.Write([]byte("Not enough arguments. Use: --file <file.json> --query <query>.\n"))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}

			continue

		} else if args[0] != "--file" || args[2] != "--query" {
			_, err := conn.Write([]byte("Not enough arguments. Use: --file <file.json> --query <query>.\n"))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}

			continue
		}

		query := strings.Join(args[3:], " ")

		if query[0] == '\'' || query[0] == '"' || query[0] == '<' {
			query = query[1:] // Убираем лишние элементы
		}

		if query[len(query)-1] == '\'' || query[len(query)-1] == '"' || query[len(query)-1] == '>' {
			query = query[:len(query)-1]
		}

		ans, err := handlers.HandleQuery(args[1], query, mut)
		if err != nil {
			log.Printf("(%s) Error: %v\n", remoteAddr, err)

			response := "Error: " + err.Error() + "\n"
			_, err := conn.Write([]byte(response))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}

			continue
		}

		// Отправка ответа клиенту
		if ans != "" {
			log.Printf("[✔] (%s) Request processed successfully.", remoteAddr)
			_, err = conn.Write([]byte("[✔] Request processed successfully: (" + ans + ")" + "\n"))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}

			continue

		} else {
			log.Printf("[✔] (%s) Request processed successfully.", remoteAddr)
			_, err := conn.Write([]byte("[✔] Request processed successfully\n"))
			if err != nil {
				log.Printf("(%s) Error: %v\n", remoteAddr, err)
				break
			}

			continue
		}
	}
}
