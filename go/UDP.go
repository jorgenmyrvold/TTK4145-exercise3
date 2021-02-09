package main

import (
	// "bufio"
	"fmt"
	"net"
	"time"
)

// serverIP: 10.100.23.147:2002

func readFromServer(port string) {

	UDPaddr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp4", UDPaddr)
	if err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buffer[0:n]))
	}
}

func sendToServer(msg string, port string) {
	UDPaddr, err := net.ResolveUDPAddr("udp4", port)
	UDPaddrSend, err := net.ResolveUDPAddr("udp4", "10.100.23.147"+port)
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp4", UDPaddr)
	if err != nil {
		fmt.Println(err)
	}

	connSend, _ := net.DialUDP("udp4", nil, UDPaddrSend)

	data := []byte(msg)
	buffer := make([]byte, 1024)

	for {
		_, err_conn := connSend.Write(data)
		if err_conn != nil {
			fmt.Println(err_conn)
		}

		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer[0:n]))

		time.Sleep(1000 * time.Millisecond)
	}

	
	conn.Close()
}


func main() {
	// readFromServer(":30000")
	// go readFromServer(":2002")
	sendToServer("Hello server!", ":20002")
}
