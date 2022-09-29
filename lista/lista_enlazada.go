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

func (listaEnlazada[T]) crearNodo(nuevoDato T) *nodo[T] {
	return &nodo[T]{dato: nuevoDato}
}

func (l listaEnlazada[T]) errores() {
	if l.EstaVacia() {
		panic("che loco esta vacio")
	}
}

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
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
	l.errores()
	dato := l.primero.dato
	if l.largo == 1 {
		l.primero = nil
		l.ultimo = nil
	} else {
		l.primero = l.primero.proximo
		l.primero.anterior = l.primero.anterior.anterior
	}
	l.largo--

	return dato
}

func (l listaEnlazada[T]) VerPrimero() T {
	l.errores()
	return l.primero.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	l.errores()
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
