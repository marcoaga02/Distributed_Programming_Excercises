package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

type NodeNotFoundError string

func (e NodeNotFoundError) Error() string {
	return fmt.Sprintf(string(e))
}

func (l *List[T]) add(newNode *List[T]) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *List[T]) remove(remNode *List[T]) error {
	if l == remNode { // se il nodo da rimuovere è la testa
		// Aggiorna la testa della lista per puntare al secondo nodo
		*l = *remNode.next
		return nil
	}
	current := l
	for current.next != nil {
		if current.next == remNode {
			current.next = remNode.next
			return nil
		}
		current = current.next
	}
	return NodeNotFoundError("Nodo non trovato")

}

func main() {
	head := &List[int]{val: 5}
	first := &List[int]{val: 10}
	second := &List[int]{val: 15}
	third := &List[int]{val: 20}

	head.add(first)
	head.add(second)
	head.add(third)

	if err := head.remove(second); err != nil {
		fmt.Println(err)
	}

	/* if err := head.remove(second); err != nil {
		fmt.Println(err)
	} */

	if err := head.remove(&List[int]{}); err != nil {
		fmt.Println(err)
	}

}
