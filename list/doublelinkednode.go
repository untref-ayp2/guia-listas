package list

// DoubleLinkedNode representa un nodo de una lista enlazada doble.
type DoubleLinkedNode[T comparable] struct {
	data T
	next *DoubleLinkedNode[T]
	prev *DoubleLinkedNode[T]
}

// NewDoubleLinkedNode crea un nuevo nodo de lista enlazada doble con el dato especificado.
func NewDoubleLinkedNode[T comparable](data T) *DoubleLinkedNode[T] {
	return &DoubleLinkedNode[T]{data: data}
}

// SetData establece el dato almacenado en el nodo.
func (n *DoubleLinkedNode[T]) SetData(data T) {
	// Implementar
}

// Data devuelve el dato almacenado en el nodo.
func (n *DoubleLinkedNode[T]) Data() T {
	// Implementar
	var x T
	return x
}

// SetNext establece el nodo siguiente al nodo actual.
func (n *DoubleLinkedNode[T]) SetNext(next *DoubleLinkedNode[T]) {
	// Implementar
}

// Next devuelve el nodo siguiente al nodo actual.
func (n *DoubleLinkedNode[T]) Next() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// HasNext devuelve true si el nodo actual tiene asignado un nodo siguiente.
func (n *DoubleLinkedNode[T]) HasNext() bool {
	// Implementar
	return false
}

// SetPrev establece el nodo anterior al nodo actual.
func (n *DoubleLinkedNode[T]) SetPrev(newPrev *DoubleLinkedNode[T]) {
	// Implementar
}

// Prev devuelve el nodo anterior al nodo actual.
func (n *DoubleLinkedNode[T]) Prev() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// HasPrev devuelve true si el nodo actual tiene asignado un nodo previo.
func (n *DoubleLinkedNode[T]) HasPrev() bool {
	// Implementar
	return false
}
