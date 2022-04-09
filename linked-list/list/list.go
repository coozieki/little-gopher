package list

import "errors"

type node[T comparable] struct {
	val  T
	next *node[T]
}

func FromArray[T comparable](arr []T) *node[T] {
	firstNode := new(node[T])
	currentNode := firstNode
	for index, el := range arr {
		currentNode.val = el
		if index < len(arr)-1 {
			currentNode.next = new(node[T])
			currentNode = currentNode.next
		}
	}
	return firstNode
}

func (n *node[T]) At(index int) (T, error) {
	currentNode := n
	for i := 0; i < index; i++ {
		if currentNode.next == nil {
			return *new(T), errors.New("element not found")
		}
		currentNode = currentNode.next
	}
	return currentNode.val, nil
}

func (n *node[T]) Search(val T) int {
	i := 0
	currentNode := n
	for currentNode.val != val {
		if currentNode.next == nil {
			return -1
		}
		currentNode = currentNode.next
		i++
	}

	return i
}

func (n *node[T]) AsArray() []T {
	currentNode := n
	var res []T
	for currentNode != nil {
		res = append(res, currentNode.val)
		currentNode = currentNode.next
	}

	return res
}

func (n *node[T]) Insert(index int, val T) error {
	if index == 0 {
		newNode := node[T]{val: n.val, next: n.next}
		n.val = val
		n.next = &newNode
		return nil
	}

	currentNode := n
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
		if currentNode == nil {
			return errors.New("index is out of range")
		}
	}
	currentNode.next = &node[T]{
		next: currentNode.next,
		val:  val,
	}

	return nil
}

func (n *node[T]) Delete(index int) error {
	if index == 0 {
		n.val = n.next.val
		n.next = n.next.next
		return nil
	}

	currentNode := n
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.next
		if currentNode == nil {
			return errors.New("index is out of range")
		}
	}
	if currentNode.next != nil {
		if currentNode.next.next != nil {
			currentNode.next = currentNode.next.next
		} else {
			currentNode.next = nil
		}
	}

	return nil
}
