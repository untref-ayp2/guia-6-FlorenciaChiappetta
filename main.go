package main

import (
	"fmt"
	"guia6/set"
	"guia6/dictionary"
	"guia6/linkedlist"
)

func main() {
	d := dictionary.NewDictionary[string, int]()
	d.Put("Leo", 60)
	d.Put("Leo", 61)
	d.Put("Fabi", 36)
	d.Put("Leo", 60)
	d.Put("Fede", 36)
	d.Put("Vale", 40)
	d.Put("Leo", 50)
	fmt.Println("Clave: valor en el diccionario (String, Int)")
	for _, key := range d.GetKeys() {
		fmt.Printf("%v:\t%d\n",key, d.Get(key))
	}
	fmt.Println("--------------------")
	d.Remove("Fede")
	fmt.Println("Borre a fede")
	fmt.Println("Al pedirle al diccionario el valor para fede me da como resultado: ", d.Get("Fede"))
	fmt.Println("--------------------")
	fmt.Println("Clave: valor en el diccionario sin fede(String, Int)")
	for _, key := range d.GetKeys() {
		fmt.Printf("%v:\t%d\n",key, d.Get(key))
	}
	fmt.Println("--------------------")
	ds := dictionary.NewDictionary[string, set.Set[int]]()
	s := set.NewSet[int]()
	s.Add(100)
	s.Add(222)
	s.Add(333)
	ss := set.NewSet[int]()
	ss.Add(1)
	ss.Add(2)
	ss.Add(3)
	ds.Put("uno", *s)
	ds.Put("dos", *ss)
	ds.Put("tres", *ss)
	fmt.Println("Clave: valor en el diccionario (String, Set[Int])")
	var aux set.Set[int]
	for _, key := range ds.GetKeys() {
		aux = ds.Get(key)
		fmt.Printf("%v:\t%v\n",key, aux.String())
	}
	fmt.Println("--------------------")
	dl := dictionary.NewDictionary[string, linkedlist.LinkedList[string]]()
	l:= linkedlist.NewLinkedList[string]()
	l.InsertAt("A",0)
	l.InsertAt("B",1)
	l.InsertAt("C",2)
	dl.Put("Lista 1",*l)
	ll:= linkedlist.NewLinkedList[string]()
	ll.InsertAt("D",0)
	ll.InsertAt("E",1)
	ll.InsertAt("F",2)
	dl.Put("Lista 2", *ll)
	fmt.Println("Clave: valor en el diccionario (String, LinkedList[String])")
	var auxl linkedlist.LinkedList[string]
	for _, key := range dl.GetKeys() {
		auxl = dl.Get(key)
		fmt.Printf("%v:\t%v\n",key, auxl.String())
	}
}