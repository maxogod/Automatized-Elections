package lista

type Lista[T any] interface {

	// EstaVacia Devuelve true si la lista esta vacia, false en caso contrario
	EstaVacia() bool

	// InsertarPrimero Inserta un elemento al principio de la lista
	InsertarPrimero(T)

	// InsertarUltimo Inserta un elemento al final de la lista
	InsertarUltimo(T)

	// BorrarPrimero Borra y devuelve el primer elemento de la lista
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista
	VerPrimero() T

	// VerUltimo devuelve el ultimo elemento de la lista
	VerUltimo() T

	// Largo devuelve el largo de la lista
	Largo() int

	// Iterar recorre todos los datos de la lista hasta el final o hasta que el func visitar devuelve false
	Iterar(visitar func(T) bool)

	// Iterador es el iterador externa de la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual Devuelve el dato en el que el iterador esta parado
	VerActual() T

	// HaySiguiente Devuelve un bool indicando si existe siguiente dato o no
	HaySiguiente() bool

	// Siguiente Avanza al siguiente dato y devuelve el dato que acaba de dejar atras
	Siguiente() T

	// Insertar Inserta un nuevo dato en la posision anterior a la que estaba parado y se queda parado en esa posicion
	Insertar(T)

	// Borrar Borra el dato en el que esta parado y lo devuelve y se queda parado en esa misma posicion
	// (que ahora tendria el dato que era el siguiente al que borre)
	Borrar() T
}
