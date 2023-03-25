package list

// DoubleLinkedList implementa una lista enlazada doble genérica.
type DoubleLinkedList[T comparable] struct {
	head *DoubleLinkedNode[T]
	tail *DoubleLinkedNode[T]
	size int
}

// NewDoubleLinkedList crea una nueva lista vacía.
//
// Uso:
//
//	list := NewDoubleLinkedList[int]() // Crea una nueva lista enlazada doble.
func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	// Implementar
	return nil
}

// Head devuelve el primer nodo de la lista.
//
// Uso:
//
//	head := list.Head() // Obtiene el primer nodo de la lista.
//
// Retorna:
//   - el primer nodo de la lista.
func (l *DoubleLinkedList[T]) Head() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// Tail devuelve el último nodo de la lista.
//
// Uso:
//
//	tail := list.Tail() // Obtiene el último nodo de la lista.
//
// Retorna:
//   - el último nodo de la lista.
func (l *DoubleLinkedList[T]) Tail() *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// Size devuelve el tamaño de la lista.
//
// Uso:
//
//	size := list.Size() // Obtiene el tamaño de la lista.
//
// Retorna:
//   - el tamaño de la lista.
func (l *DoubleLinkedList[T]) Size() int {
	// Implementar
	return -1
}

// IsEmpty evalúa si la lista está vacía.
//
// Uso:
//
//	empty := list.IsEmpty() // Verifica si la lista está vacía.
//
// Retorna:
//   - true si la lista está vacía; false en caso contrario.
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	// Implementar
	return false
}

// Clear elimina todos los nodos de la lista.
//
// Uso:
//
//	list.Clear() // Elimina todos los nodos de la lista.
func (l *DoubleLinkedList[T]) Clear() {
	// Implementar
}

// Prepend inserta un dato al inicio de la lista.
//
// Uso:
//
//	list.Prepend(10) // Inserta el dato 10 al inicio de la lista.
//
// Parámetros:
//   - `data`: el dato a insertar al frente de la lista.
func (l *DoubleLinkedList[T]) Prepend(data T) {
	// Implementar
}

// Append inserta un dato al final de la lista.
//
// Uso:
//
//	list.Append(10) // Inserta el dato 10 al final de la lista.
//
// Parámetros:
//   - `data`: el dato a insertar al final de la lista.
func (l *DoubleLinkedList[T]) Append(data T) {
	// Implementar
}

// Find busca un dato en la lista
//
// Uso:
//
//	node := list.Find(10) // Busca el dato 10 en la lista.
//
// Parámetros:
//   - `data`: el dato a buscar en la lista.
//
// Retorna:
//   - el nodo que contiene el dato buscado; `nil` si no se encuentra.
func (l *DoubleLinkedList[T]) Find(data T) *DoubleLinkedNode[T] {
	// Implementar
	return nil
}

// RemoveFirst elimina el primer nodo de la lista.
//
// Uso:
//
//	list.RemoveFirst() // Elimina el primer nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveFirst() {
	// Implementar
}

// RemoveLast elimina el último nodo de la lista.
//
// Uso:
//
//	list.RemoveLast() // Elimina el último nodo de la lista.
func (l *DoubleLinkedList[T]) RemoveLast() {
	// Implementar
}

// Remove elimina un la primera aparición de un dato en la lista.
//
// Uso:
//
//	list.Remove(10) // Elimina la primera aparición del dato 10 en la lista.
func (l *DoubleLinkedList[T]) Remove(data T) {
	// Implementar
}

// String devuelve una representación en cadena de la lista.
//
// Uso:
//
//	fmt.Println(list) // Imprime la representación en cadena de la lista.
//
// Retorna:
//   - una representación en cadena de la lista.
func (l *DoubleLinkedList[T]) String() string {
	// Implementar
	return ""
}
