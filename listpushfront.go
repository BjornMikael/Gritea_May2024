package piscine

// NodeL represents a node in the linked list
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents the linked list itself
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushFront inserts a new element at the beginning of the list
func ListPushFront(l *List, data interface{}) {
	// Create a new node with the given data
	newNode := &NodeL{
		Data: data,
		Next: l.Head,
	}

	// Update the head of the list to be the new node
	l.Head = newNode

	// If the list was empty, set the tail to the new node as well
	if l.Tail == nil {
		l.Tail = newNode
	}
}
