package server

import (
	"fileserver/internal/protocol"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn   net.Conn
	writer *protocol.MessageWriter
}

type TcpChatServer struct {
	listener net.Listener
	mutex    sync.Mutex
	clients  map[net.Addr]client
}

func NewTcpChatServer() *TcpChatServer {
	return &TcpChatServer{clients: make(map[net.Addr]client)}
}

func (s *TcpChatServer) ListenAndServe(address string) error {
	l, err := net.Listen("tcp", address)

	if err == nil {
		s.listener = l
	}

	s.start()
	return err
}

func (s *TcpChatServer) Close() {
	s.listener.Close()
}

func (s *TcpChatServer) start() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Println(err)
		} else {
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %d", conn.RemoteAddr(), len(s.clients)+1)
	client := client{conn: conn,
		writer: protocol.NewMessageWriter(conn),
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients[client.conn.RemoteAddr()] = client
	return &client
}

func (s *TcpChatServer) serve(client *client) {
	msgReader := protocol.NewMessageReader(client.conn)
	defer s.remove(client)

	for {
		command, err := msgReader.Read()

		if err != nil && err != io.EOF {
			log.Printf("Read error: %v\n", err)
			client.writer.Write([]byte(err.Error()))
		}

		if err == io.EOF {
			break
		}

		if command != nil {
			log.Println("Executing")
			data, err := command.Handle()
			if err != nil {
				log.Printf("Error: %v\n", err)
				client.writer.Write([]byte(err.Error()))
			} else {
				client.writer.Write(data)
			}
		}
	}
}

func (s *TcpChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clients, client.conn.RemoteAddr())
	client.conn.Close()
}
