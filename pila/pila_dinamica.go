package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	const _FACTOR_AGRANDADOR = 2
	if p.cantidad == len(p.datos) {
		p.redimencionar(len(p.datos) * _FACTOR_AGRANDADOR)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	const (
		_LARGO_MINIMO     = 10
		_CUADRIPLICADOR   = 4
		_FACTOR_ACHICADOR = 2
	)
	if p.cantidad*_CUADRIPLICADOR <= len(p.datos) && p.cantidad*_CUADRIPLICADOR >= _LARGO_MINIMO {
		p.redimencionar(len(p.datos) / _FACTOR_ACHICADOR)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func (p *pilaDinamica[T]) redimencionar(nuevoLargo int) {
	newArr := make([]T, nuevoLargo)
	copy(newArr, p.datos)
	p.datos = newArr
}

func CrearPilaDinamica[T any]() Pila[T] {
	const _LARGO_INICIAL = 10
	p := new(pilaDinamica[T])
	p.datos = make([]T, _LARGO_INICIAL)
	return p
}
