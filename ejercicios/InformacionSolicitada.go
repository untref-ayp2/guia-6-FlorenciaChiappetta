package ejercicios

import (
	"guia6/dictionary"
)

// 4. Tenemos que ayudar a los docentes a preparar la información solicitada por el
// Departamento de Alumnos.
// Cuando toman asistencia anotan los presentes por
// día de clase, pero ahora desde Alumnos se les pide que envíen la información
// de asistencia por alumno.  Por ejemplo, si la
// lista que se recibe es:
// {“Mie 10”: [“Ana", "Pedro"], “Vie 12”:, [“Ana", "Luz”], “Mie 17”:, [“Luz”, "Pedro"]}
// el resultado debe ser
// {"Ana":["Mie 10", "Vie 12"], "Luz": ["Vie 12", "Mie 17"], "Pedro": ["Mie 10", "Mie 17"]}
func Ejercicio2(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {
	dict := dictionary.NewDictionary[string, []string]()
	for _, dia := range entrada.GetKeys() {
		for _, nombre := range entrada.Get(dia) {
			if dict.Contains(nombre) {
				value := dict.Get(nombre)
				value = append(value, dia)
				dict.Put(nombre, value)
			} else {
				value := []string{dia}
				dict.Put(nombre, value)
			}
		}
	}
	return &dict
}

func Ejercicio2Invertido(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {

	dict := dictionary.NewDictionary[string, []string]()
	for _, nombre := range entrada.GetKeys() {
		for _, dia := range entrada.Get(nombre) {
			if dict.Contains(dia) {
				values := entrada.Get(dia)
				values = append(values, nombre)
				dict.Put(dia, values)
			} else {
				values := []string{nombre}
				dict.Put(dia, values)
			}
		}

	}
	return &dict
}

// func Ejercicio2 (tareas[]Tarea)
// // ordenar las tareas por tiempo usando burbujeo

// 	for i:=0; i < len(tareas)-1;i++{
// 		for j:=0; j < len(tareas)-1;j++{

// 	if tareas[j].Tiempo> tareas[j+1].Tiempo{

// 	tareas[j],tareas[j+1] = tareas[j+1],tareas[j]
// 		}
// 	}
// }
// func Materias (materias dictionary.Dictionary[string,[]string]) dictionary.Dictionary.NewDictionary[string,[]string]()

// for _, apellido := range materias.GetKeys(){
// for _, alumno := range materias.Get(apellido){

// arreglo := []string{}
// if alumnos.Contains(alumno){
// arreglo = alumnos.Get(alumno)
// }
// arreglo = append(arreglo,apellido)
// alumnos.Put(alumno,apellido)
// }
// }
// return alumnos
// }
