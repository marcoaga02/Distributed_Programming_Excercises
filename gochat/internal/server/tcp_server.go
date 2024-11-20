package server

import (
	"gochat/internal/protocol"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn   net.Conn
	name   string
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
		msg, err := msgReader.Read()

		if err != nil && err != io.EOF {
			log.Printf("Read error: %v\n", err)
		}

		if err == io.EOF {
			break
		}

		if msg != nil {
			switch v := msg.(type) {
			case protocol.SendMessage:
				bmsg := protocol.NotifyMessage{
					Author:  client.name,
					Content: v.Content,
				}
				go s.Broadcast(bmsg)
			case protocol.ChangeNameMessage:
				client.name = v.NewName
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

func (s *TcpChatServer) Broadcast(msg any) {
	for _, c := range s.clients {
		_, err := c.writer.Write(msg)
		if err != nil {
			log.Printf("Error write: %v", err)
		}
	}
}
