package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) crearNodo(nuevoDato T) *nodo[T] {
	return &nodo[T]{dato: nuevoDato}
}

func (c *colaEnlazada[T]) Encolar(nuevoDato T) {
	nuevoNodo := c.crearNodo(nuevoDato)

	if c.primero == nil {
		c.primero = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	elem := c.primero.dato
	c.primero = c.primero.siguiente
	return elem
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}
