package main

import (
	"fmt"
	"net"
	"runtime"
)

const (
	NEW_CONN   = 1
	CONN_CLOSE = 2
)

type Connection struct {
	id         uint32
	connection net.Conn
	output     chan string
	open       bool
	server     *Server
}

type ConnectionEvent struct {
	eventType  int
	connection *Connection
}

type ConnectionMsg struct {
	message    string
	connection *Connection
}

type Server struct {
	host string
	port string

	events   chan ConnectionEvent
	messages chan ConnectionMsg

	connections map[uint32]*Connection
}

func (server *Server) init(host string, port string) {
	server.host = host
	server.port = port
	server.events = make(chan ConnectionEvent, 128)
	server.messages = make(chan ConnectionMsg, 128)
	server.connections = make(map[uint32]*Connection)
}

func (server *Server) getAdr() string {
	return fmt.Sprintf("%s:%s", server.host, server.port)
}

func (server *Server) listen() {
	listener, err := net.Listen("tcp", server.getAdr())
	if err != nil {
		fmt.Printf("Error creating listener: %s\n", err)
		return
	}
	var nextId uint32 = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error while accepting incoming connection: %s\n", err)
		} else {
			nextId++

			connection := new(Connection)
			connection.connection = conn
			connection.id = nextId
			connection.output = make(chan string)
			connection.open = true
			connection.server = server
			server.events <- ConnectionEvent{eventType: NEW_CONN, connection: connection}
		}
	}
}

func (server *Server) handleConnections() {
	for {
		select {
		case evt := <-server.events:
			conn := evt.connection
			switch evt.eventType {
			case NEW_CONN:
				server.connections[conn.id] = conn
				go conn.handleInput()
				go conn.handleOutput()
			case CONN_CLOSE:
				delete(server.connections, conn.id)
			default:
			}
		case msg := <-server.messages:
			srcConn := msg.connection
			for id, conn := range server.connections {
				if id != srcConn.id {
					conn.output <- fmt.Sprintf("%d:%s", srcConn.id, msg.message)
				}
			}
		default:
			runtime.Gosched()
		}
	}
}

func (conn *Connection) handleInput() {
	buf := make([]byte, 1024)
	for {
		read, err := conn.connection.Read(buf)
		if err != nil {
			break
		}
		conn.server.messages <- ConnectionMsg{connection: conn, message: string(buf[0:read])}
	}
	conn.close()
}

func (conn *Connection) handleOutput() {
	buf := make([]byte, 1024)
	for message := range conn.output {
		bytesWritten := copy(buf, message)
		_, error := conn.connection.Write(buf[0:bytesWritten])
		if error != nil {
			fmt.Printf("error writing to the connection Id %d\n", conn.id)
			break
		}
	}
	conn.close()
}

func (conn *Connection) close() {
	if conn.open {
		conn.server.events <- ConnectionEvent{eventType: CONN_CLOSE, connection: conn}
		conn.connection.Close()
		close(conn.output)
		conn.open = false
	}
}

func main() {

	server := Server{}
	server.init("localhost", "3000")

	fmt.Println("Listening on: " + server.getAdr())

	go server.handleConnections()
	go server.listen()

	// hack to make program wait indefinitely
	// wg := new(sync.WaitGroup)
	// wg.Add(1)
	// wg.Wait()

	// Alternatively,
	select {}

}
