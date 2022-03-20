package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func countingDownHandler(conn net.Conn) {
	defer func() {
		io.WriteString(conn, "Your connection will be closed by server")
		conn.Close()
	}()

	io.WriteString(conn, "Enter number : ")
	input := bufio.NewScanner(conn)
	count := Scan(input)
	if count > 20 {
		return
	}
	
	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n")
		time.Sleep(time.Second)
		count--
		if count < 0 {
			io.WriteString(conn, "Enter number : ")
			count = Scan(input)
			if count == 0 {
				// cannot scan
				break
			}
		}
	}
}

func Scan(input *bufio.Scanner) int {
	if ok := input.Scan(); !ok {
		log.Println("Cannot scan value from conn")
		log.Printf("Connection is closed by client")
		return 0
	}
	count, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("Connot convert value from Text to int. Value : ", input.Text())
		return 0
	}

	return count
}

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		log.Println("New client is connected : ", conn.RemoteAddr())
		if err != nil {
			log.Panicln(err)
		}
		// Logic
		go countingDownHandler(conn)
	}
}
