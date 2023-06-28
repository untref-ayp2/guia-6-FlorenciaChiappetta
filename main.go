package main

import (
	"fmt"
	"guia6/linkedlist"
)

func main() {
	// d := dictionary.NewDictionary[string, int]()
	// d.Put("Leo", 60)
	// d.Put("Leo", 61)
	// d.Put("Fabi", 36)
	// d.Put("Leo", 60)
	// d.Put("Fede", 36)
	// d.Put("Vale", 40)
	// d.Put("Leo", 50)
	// fmt.Println("Clave: valor en el diccionario (String, Int)")
	// for _, key := range d.GetKeys() {
	// 	fmt.Printf("%v:\t%d\n",key, d.Get(key))
	// }
	// fmt.Println("--------------------")
	// d.Remove("Fede")
	// fmt.Println("Borre a fede")
	// fmt.Println("Al pedirle al diccionario el valor para fede me da como resultado: ", d.Get("Fede"))
	// fmt.Println("--------------------")
	// fmt.Println("Clave: valor en el diccionario sin fede(String, Int)")
	// for _, key := range d.GetKeys() {
	// 	fmt.Printf("%v:\t%d\n",key, d.Get(key))
	// }
	// fmt.Println("--------------------")
	// ds := dictionary.NewDictionary[string, set.Set[int]]()
	// s := set.NewSet[int]()
	// s.Add(100)
	// s.Add(222)
	// s.Add(333)
	// ss := set.NewSet[int]()
	// ss.Add(1)
	// ss.Add(2)
	// ss.Add(3)
	// ds.Put("uno", *s)
	// ds.Put("dos", *ss)
	// ds.Put("tres", *ss)
	// fmt.Println("Clave: valor en el diccionario (String, Set[Int])")
	// var aux set.Set[int]
	// for _, key := range ds.GetKeys() {
	// 	aux = ds.Get(key)
	// 	fmt.Printf("%v:\t%v\n",key, aux.String())
	// }
	// fmt.Println("--------------------")
	// dl := dictionary.NewDictionary[string, linkedlist.LinkedList[string]]()
	// l:= linkedlist.NewLinkedList[string]()
	// l.InsertAt("A",0)
	// l.InsertAt("B",1)
	// l.InsertAt("C",2)
	// dl.Put("Lista 1",*l)
	// ll:= linkedlist.NewLinkedList[string]()
	// ll.InsertAt("D",0)
	// ll.InsertAt("E",1)
	// ll.InsertAt("F",2)
	// dl.Put("Lista 2", *ll)
	// fmt.Println("Clave: valor en el diccionario (String, LinkedList[String])")
	// var auxl linkedlist.LinkedList[string]
	// for _, key := range dl.GetKeys() {
	// 	auxl = dl.Get(key)
	// 	fmt.Printf("%v:\t%v\n",key, auxl.String())
	// }

	// d2 := dictionary.NewDictionary[string, []string]()
	// arreglo1 := []string{"Pedro", "Juana"}
	// d2.Put("Miercoles", arreglo1)
	// arreglo2 := []string{"Luis", "Pedro"}
	// d2.Put("Jueves", arreglo2)
	// arreglo3 := []string{"Luis", "Juana"}
	// d2.Put("Viernes", arreglo3)
	// // arreglo4 := []string{"Lorena", "Luis", "Ana"}
	// // d.Put("Sabado", arreglo4)
	// fmt.Println(d2)
	// fmt.Println(ejercicios.Ejercicio2(d2))

	// fmt.Println("\n", "--------------------------------------------------------------------")

	// d := dictionary.NewDictionary[string, []string]()
	// arregloI := []string{"Miercoles", "Viernes"}
	// d.Put("Ana", arregloI)
	// arregloII := []string{"Viernes", "Jueves"}
	// d.Put("Juan", arregloII)
	// arregloIII := []string{"Miercoles", "Jueves"}
	// d.Put("Jose", arregloIII)
	// // arreglo4 := []string{"Lorena", "Luis", "Ana"}
	// // d.Put("Sabado", arreglo4)
	// fmt.Println(d)
	// fmt.Println(ejercicios.Ejercicio2Invertido(d))

	// d := dictionary.NewDictionary[string, *dictionary.Dictionary[int, []string]]()
	// d2 := dictionary.NewDictionary[int, []string]()
	// ar := []string{"Individual F US Open", "Doble Mixto Wimbledon"}
	// ab := []string{"Dobles US Open", "Dobles F Wimbledon"}
	// d2.Put(1990, ar)
	// d2.Put(1988, ab)
	// d.Put("Gabriela Sabattini", &d2)
	// fmt.Println(d2)
	// fmt.Println(ejercicios.Deportistas(&d))

	// d := linkedlist.NewLinkedList()
	// d.Insert(1)
	// d.Insert(2)
	// d.Insert(3)
	// fmt.Println(d.String())

	// d := dictionary.NewDictionary[string, []int]()
	// a := []int{4, 4, 8, 6}
	// d.Put("Perez", a)

	// b := []int{5, 5, 6, 7}
	// d.Put("Sanchez", b)
	// c := []int{9, 8, 7, 6}
	// d.Put("Juan", c)
	// e := []int{10, 2, 8, 6}
	// d.Put("Flores", e)
	// fmt.Println(d.NotaFinal(&d))

	// fmt.Println("\n", "--------------------------------------------------------------------")
	// var q1 dictionary.Queue
	// var s dictionary.Stack
	// q1.Enqueue(1)
	// q1.Enqueue(2)
	// q1.Enqueue(3)
	// q1.Enqueue(4)
	// q1.Enqueue(5)
	// fmt.Println(s.RecibirPila(q1))

	// dict := dictionary.NewDictionary[string, []int]()
	// notas1 := []int{4, 10, 0, 2}
	// dict.Put("Juan", notas1)
	// notas2 := []int{2, 7, 7, 0}
	// dict.Put("Ana", notas2)
	// notas3 := []int{2, 2, 7, 10}
	// dict.Put("Pedro", notas3)
	// notas4 := []int{5, 7, 4, 0}
	// dict.Put("Jose", notas4)
	// fmt.Println(dict.Aprobados(dict))

	// l := linkedlist.NewLinkedList[int]()
	// l.Append(1)
	// l.Append(2)
	// l.Append(3)
	// l.Append(4)
	// fmt.Print(l)
	// // l.Revertir()
	// // fmt.Print(l)

	// a := []int{1, 2, 3, 4, 5, 7, 90, 9}
	// linkedlist.Maximo(a)

	// arreglo := []int{1, 0, 2, 3, 4, 5, 7, 90, 9, 1293, 9902, 70}
	// fmt.Println(linkedlist.MaxandMin(arreglo))

	arreglo1 := []int{9, 11, 23, 24, 4}

	fmt.Println(linkedlist.ArregloEsMagico(arreglo1))
}
