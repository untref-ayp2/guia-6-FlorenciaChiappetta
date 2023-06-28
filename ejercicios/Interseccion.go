package ejercicios

import "errors"

// agregar import faltante

// func Interseccion(s1 []string, s2 []string) linkedlist.LinkedList[string]{
// 	panic("Not implemented")
// }

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack struct {
	stack []any
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack) Push(v any) {
	(*s).stack = append((*s).stack, v)
}

// Pop saca un elemento de la pila. O(1)
func (s *Stack) Pop() (any, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := (*s).stack[len((*s).stack)-1]
	(*s).stack = ((*s).stack)[:len((*s).stack)-1]
	return v, nil
}

// Top devuelve el elemento del tope de la pila. O(1)
func (s *Stack) Top() (any, error) {
	if (*s).IsEmpty() {
		return 0, errors.New("la pila esta vacia")
	}
	v := (*s).stack[len((*s).stack)-1]
	return v, nil
}

// IsEmpty verifica si la pila esta vacia. O(1)
func (s *Stack) IsEmpty() bool {
	return len((*s).stack) == 0
}

func (s *Stack) Size() int {
	return len((*s).stack)
}

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	cola []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	(*q).cola = append((*q).cola, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {

	if len((*q).cola) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := ((*q).cola)[0]
	(*q).cola = (*q).cola[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if len((*q).cola) == 0 {
		return nil, errors.New("queue is empty")
	}
	return ((*q).cola)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len((*q).cola) == 0
}

func BuscarMaximoConPila(arreglo []int) int {

	var s1 Stack
	max := -1000000

	for _, numero := range arreglo {

		if numero > max {

			s1.Push(numero)
			max = numero
		}
	}
	return max
}

func RecibirPila(q Queue) *Stack {
	var s1 Stack
	if q.IsEmpty() {
		return &s1
	}
	value, _ := q.Dequeue()
	RecibirPila(q)
	s1.Push(value)
	return &s1
}
