package lista

type nodo[T any] struct {
	dato    T
	proximo *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	//TODO averiguar q onda esto
}

func (l listaEnlazada[T]) EstaVacia() bool {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) InsertarPrimero(t T) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) InsertarUltimo(t T) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) BorrarPrimero() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) VerPrimero() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) VerUltimo() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Largo() int {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Iterador() IteradorLista[T] {
	//TODO implement me
	panic("implement me")
}

func CrearListaEnlazada[T any]() Lista[T] {
	l := new(listaEnlazada[T])
	return l
}
