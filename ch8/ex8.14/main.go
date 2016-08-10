// Exercise 8.14:
// Change the chat server's protocol so that each client provides its name
// on entering. Use that name instead of the network address when prefixing
// each message with its sender's identity.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stderr, "%s: too few arguments.\nusage: %[1]s host port\n", os.Args[0])
		os.Exit(1)
	}
	host := os.Args[1]
	port := os.Args[2]
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chat server started at %s:%s\n", host, port)

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	// Ask for nickname
	who, err := askName(conn)
	if err != nil {
		return
	}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func askName(conn net.Conn) (string, error) {
	fmt.Fprint(conn, "请输入你的名字（直接回车，默认为IP地址）: ")
	name, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return "", err
	}
	if len(name) <= 1 {
		return conn.RemoteAddr().String(), nil
	} else {
		return string(name[:len(name)-1]), nil
	}
}
