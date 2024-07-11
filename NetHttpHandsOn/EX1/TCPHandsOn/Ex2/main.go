package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		scan_obj := bufio.NewScanner(conn)

		for scan_obj.Scan() {
			ln := scan_obj.Text()
			fmt.Println(ln)
		}

		defer conn.Close()
	}

}
