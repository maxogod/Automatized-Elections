package lista

type Lista[T any] interface {
	EstaVacia() bool

	InsertarPrimero(T)

	InsertarUltimo(T)

	BorrarPrimero() T

	VerPrimero() T

	VerUltimo() T

	Largo() int

	Iterar(visitar func(T) bool)

	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual Devuelve el dato en el que el iterador esta parado
	VerActual() T

	// HaySiguiente Devuelve un bool indicando si existe siguiente dato o no
	HaySiguiente() bool

	// Siguiente Avanza al siguiente dato y devuelve el dato que acaba de dejar atras
	Siguiente() T

	// Insertar Inserta un nuevo dato en la posision anterior a la que esta parado
	Insertar(T)

	// Borrar Borra el dato en el que esta parado y lo devuelve
	Borrar() T
}
