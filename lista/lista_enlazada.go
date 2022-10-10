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
	lista            *listaEnlazada[T]
	posicionActual   *nodo[T]
	posicionAnterior *nodo[T]
}

// Metodos de listaEnlazada

func (l *listaEnlazada[T]) panicSiVacia() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) crearNodo(nuevoDato T) *nodo[T] {
	return &nodo[T]{dato: nuevoDato}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	prox := l.primero
	l.primero = nuevoNodo
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	} else {
		l.primero.proximo = prox
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	ant := l.ultimo
	l.ultimo = nuevoNodo
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		ant.proximo = l.ultimo
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	l.panicSiVacia()
	const _LARGO_MIN = 1
	dato := l.primero.dato

	if l.largo == _LARGO_MIN {
		l.primero = nil
		l.ultimo = nil
	} else {
		l.primero = l.primero.proximo
	}
	l.largo--

	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	l.panicSiVacia()
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	l.panicSiVacia()
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

// Iterar - iterador Interno
func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for l.primero != nil && visitar(l.primero.dato) {
		l.primero = l.primero.proximo
	}
}

// Iterador - Crea un iterador externo
func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorListaEnlazada[T])
	iter.lista = l
	iter.posicionActual = l.primero
	iter.posicionAnterior = l.primero
	return iter
}

//Metodos de iteradorListaEnlazada (EXTERNO)

func (i *iteradorListaEnlazada[T]) panicSinSiguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (i *iteradorListaEnlazada[T]) VerActual() T {
	i.panicSinSiguiente()
	return i.posicionActual.dato
}

func (i *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return i.posicionActual != nil
}

func (i *iteradorListaEnlazada[T]) Siguiente() T {
	i.panicSinSiguiente()
	elemento := i.posicionActual.dato
	if i.posicionActual != i.lista.primero {
		i.posicionAnterior = i.posicionAnterior.proximo
	}
	i.posicionActual = i.posicionActual.proximo
	return elemento
}

func (i *iteradorListaEnlazada[T]) Insertar(nuevoDato T) {

	if i.lista.EstaVacia() || i.posicionActual == i.lista.primero {
		// Vacia o Principio
		i.lista.InsertarPrimero(nuevoDato)
		i.posicionActual = i.lista.primero
		i.posicionAnterior = i.lista.primero
	} else if i.posicionActual == nil {
		// Final
		i.lista.InsertarUltimo(nuevoDato)
		i.posicionActual = i.lista.ultimo
	} else {
		// Medio
		nuevoNodo := i.lista.crearNodo(nuevoDato)

		prox := i.posicionActual
		i.posicionActual = nuevoNodo
		i.posicionAnterior.proximo = i.posicionActual
		i.posicionActual.proximo = prox
		i.lista.largo++
	}
}

func (i *iteradorListaEnlazada[T]) Borrar() T {
	i.panicSinSiguiente()
	i.lista.panicSiVacia()
	dato := i.posicionActual.dato

	if i.posicionActual == i.lista.primero {
		// Principio
		i.lista.BorrarPrimero()
		i.Siguiente()
	} else if i.posicionActual == i.lista.ultimo {
		// Final
		i.posicionAnterior.proximo, i.posicionActual = nil, nil
		i.lista.ultimo = i.posicionAnterior
		i.lista.largo--
	} else {
		// Medio
		i.posicionAnterior.proximo = i.posicionActual.proximo
		i.posicionActual = i.posicionAnterior.proximo
		i.lista.largo--
	}
	return dato
}

// CrearListaEnlazada - Funcion creadora de lista
func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}
