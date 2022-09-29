package lista

type nodo[T any] struct {
	dato     T
	proximo  *nodo[T]
	anterior *nodo[T]
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
	return l.largo == 0
}

func (listaEnlazada[T]) crearNodo(nuevoDato T) *nodo[T] {
	return &nodo[T]{dato: nuevoDato}
}

func (l *listaEnlazada[T]) InsertarPrimero(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	if l.primero == nil {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		prox := l.primero
		l.primero = nuevoNodo
		l.primero.proximo = prox
		l.primero.proximo.anterior = l.primero
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	if l.ultimo == nil {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		ant := l.ultimo
		l.ultimo = nuevoNodo
		l.ultimo.anterior = ant
		l.ultimo.anterior.proximo = l.ultimo
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	dato := l.primero.dato
	l.primero = l.primero.proximo
	l.ultimo = l.ultimo.proximo
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l listaEnlazada[T]) VerPrimero() T {
	return l.primero.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	return l.ultimo.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
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
	return new(listaEnlazada[T])
}
