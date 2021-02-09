package main

import (
	"net"
	"fmt"
	"bufio"
)

func connectToServer(port string) {
	conn, err := net.Dial("tcp", "10.100.23.147"+port)
	if err != nil {
		fmt.Println(err)
	}

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	conn.Close()
}

func main() {
	connectToServer(":34933")  // Fixed length
	// connectToServer(":33546")  // 0 terminated

}