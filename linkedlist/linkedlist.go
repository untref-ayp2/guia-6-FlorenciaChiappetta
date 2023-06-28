package linkedlist

import (
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T int] struct {
	value int
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T int](value int) *node[int] {
	return &node[int]{value: 0, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T int] struct {
	head *node[int] // puntero al primer nodo
	tail *node[int] // puntero al último nodo
	size int
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T int]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// La posición 0 agrega el nodo al principio de la lista O(1)
// La posición size agrega el nodo al final de la lista O(1)
// Insertar en otra posición. O(n)
func (l *LinkedList[T]) InsertAt(value int, position int) {
	if position < 0 || position > l.size {
		return
	}
	newNode := newNode(value)
	// Insertar al principio
	// O(1)
	if position == 0 {
		newNode.next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
		l.size++
		return
	}
	// Insertar al final
	// O(1)
	if position == l.size {
		l.tail.next = newNode
		l.tail = newNode
		l.size++
		return
	}
	// Insertar en position
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	newNode.next = current.next
	current.next = newNode
	l.size++
}

// RemoveAt elimina el nodo en la posición recibida
// Si la posición es inválida, no hace nada
// La posición 0 elimina el primer nodo de la lista O(1)
// Eliminar en otra posición. O(n)
func (l *LinkedList[T]) RemoveAt(position int) {
	if position < 0 || position >= l.size {
		return
	}
	// Eliminar el primer nodo
	// O(1)
	if position == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	current.next = current.next.next
	l.size--
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *LinkedList[T]) Search(value int) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) int {
	if position < 0 || position >= l.size {
		var t int // la variable t se inicializ con un valor nulo
		return t
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}

	return current.value
}

// Size devuelve la cantidad de nodos en la lista
// O(n)
func (l *LinkedList[T]) Size() int {
	return l.size
}

// Contains verifica si el valor recibido está en la lista
// O(n)
func (l *LinkedList[T]) Contains(value int) bool {
	return l.Search(value) != -1
}

// IsEmpty verifica si la lista está vacía
// O(1)
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "lista: ["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

// func (l *LinkedList[T]) Insert(value int) {
// 	nodo := newNode[T](value)
// 	if l.head == nil {
// 		l.head = nodo
// 		l.tail = nodo
// 		l.tail.next = nil
// 	}
// 	posActual := l.head
// 	for posActual != nil {
// 		if posActual.value == value {
// 			return
// 		}
// 		if posActual.value < value && value < posActual.next.value {
// 			nodo.next = posActual.next.next
// 			posActual.next = nodo
// 		}
// 		if posActual.value > value {
// 			nodo.next = l.head
// 			l.head = nodo
// 		}

// 		posActual = posActual.next
// 	}

// }

func (l *LinkedList[T]) Append(value int) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode

}

// Prepend agrega un nuevo nodo, con el valor recibido,
// al inicio de la lista
// O(1)
func (l *LinkedList[T]) Prepend(value int) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode

}

// Escribir una función que devuelva el elemento máximo de una lista de enteros
// que utilice la técnica de división y conquista.
// Calcular el orden de la solución aplicando el Teorema del Maestro.
func Maximo(arreglo []int) int {
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	medio := len(arreglo) / 2
	max_izquierda := Maximo(arreglo[:medio])
	max_derecha := Maximo(arreglo[medio:])

	return max(max_izquierda, max_derecha)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func MaxandMin(arreglo []int) (int, int) {
	if len(arreglo) == 1 {
		return arreglo[0], arreglo[0]
	}
	medio := len(arreglo) / 2
	minizq, maxizq := MaxandMin(arreglo[:medio])
	minderecha, maxderecha := MaxandMin(arreglo[medio:])

	return min(minizq, minderecha), max(maxizq, maxderecha)

}

func ArregloEsMagico(arreglo []int) bool {

	return ValidarArreglo(arreglo, 0, len(arreglo))
}

func ValidarArreglo(Arreglo []int, posinicial int, posfinal int) bool {
	if Arreglo[posinicial] == posinicial {
		return true
	} else if posinicial <= posfinal {
		ValidarArreglo(Arreglo, posinicial+1, posfinal)
	}
	return false
}
