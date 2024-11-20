package server

type ChatServer interface {
	ListenAndServe(address string) error
	Broadcast(message any) error
	Close()
}
