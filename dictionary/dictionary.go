package dictionary

import (
	"errors"
	"guia6/linkedlist"
)

// Diccionario es un conjunto de entradas formadas por pares únicos (clave: valor)
type Dictionary[K comparable, V any] struct {
	mapa map[K]V
}

/*********Métodos*********/
// Crea un diccionario
// O(1)
func NewDictionary[K comparable, V any]() Dictionary[K, V] {
	dict := Dictionary[K, V]{mapa: make(map[K]V)}
	return dict
}

// Agrega una nueva entrada al diccionario, si existe pisa el valor para esa entrada
// O(1)
func (dict *Dictionary[K, V]) Put(key K, val V) {
	dict.mapa[key] = val
}

// Remueve una entrada del diccionario
// O(1)
func (dict *Dictionary[K, V]) Remove(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	if exists {
		delete(dict.mapa, key)
	}
	return exists
}

// Verifica si existe una entrada para ese valor de clave
// O(1)
func (dict *Dictionary[K, V]) Contains(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	return exists
}

// Devuelve el valor para esa clave
//O(1)
func (dict *Dictionary[K, V]) Get(key K) V {
	return dict.mapa[key]
}

// Devuelve la cantidad de entradas del diccionario
// O(1)
func (dict *Dictionary[K, V]) Size() int {
	return len(dict.mapa)
}

// Devuelve un arreglo con las claves del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetKeys() []K {
	var dictKeys []K
	dictKeys = []K{}
	var key K
	for key = range dict.mapa {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// Devuelve un arreglo con las valores del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetValues() []V {
	var dictValues []V
	dictValues = []V{}
	var key K
	for key = range dict.mapa {
		dictValues = append(dictValues, dict.mapa[key])
	}
	return dictValues
}

func (dict *Dictionary[K, V]) NotaFinal(entrada *Dictionary[string, []int]) *Dictionary[string, float64] {
	d := NewDictionary[string, float64]()
	for _, apellido := range entrada.GetKeys() {
		arreglo := entrada.Get(apellido)
		d.Put(apellido, promedio(arreglo))
	}
	return &d
}

func promedio(arreglo []int) float64 {
	suma := float64(sumatoria(arreglo))
	cantNotas := len(arreglo)

	return suma / float64(cantNotas)

}
func sumatoria(arreglo []int) int {
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	medio := len(arreglo) / 2

	return sumatoria(arreglo[:medio]) + sumatoria(arreglo[medio:])
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

func ColaInvertida(q *Queue) *Queue {
	if q.IsEmpty() {
		return q
	}
	a, _ := q.Dequeue()
	ColaInvertida(q)
	q.Enqueue(a)
	return q
}

func (dict *Dictionary[K, V]) Aprobados(entrada Dictionary[string, []int]) []string {
	aprobados := []string{}
	for _, alumno := range entrada.GetKeys() {
		aux := entrada.Get(alumno)
		if aprobo(aux) {
			aprobados = append(aprobados, alumno)

		}
	}
	return aprobados
}

func aprobo(arreglo []int) bool {
	if len(arreglo) != 4 {
		return false
	}
	if arreglo[0] >= 4 && (arreglo[2] >= 4 || arreglo[3] >= 4) {
		return true

	}
	if arreglo[1] >= 4 && (arreglo[2] >= 4 || arreglo[3] >= 4) {
		return true

	}
	return false
}

func Agregar(g Dictionary[string, []string], palabra, def string) *Dictionary[string, []string] {
	if g.Contains(palabra) {
		aux := g.Get(palabra)
		aux = append(aux, def)
		g.Put(palabra, aux)
	}
	aux := []string{def}
	g.Put(palabra, aux)
	return &g
}

//un método Buscar que dada una palabra muestre por pantalla todas las definiciones de la misma
func Buscar(g Dictionary[string, []string], palabra string) []string {
	return g.Get(palabra)
}

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

func (s1 *Stack) RecibirPila(q Queue) *Stack {
	if q.IsEmpty() {
		return s1
	}
	value, _ := q.Dequeue()
	s1.RecibirPila(q)
	s1.Push(value)
	return s1
}

// 3. Escriba una función que toma 2 slices 
// y devuelve una nueva lista que es el resultado de la interseccion de los 2 slices anteriores,
//  la complejidad del metodo debe ser O(n)


func Interseccion(s1 []string, s2 []string) *linkedlist.LinkedList[string]{
	

}