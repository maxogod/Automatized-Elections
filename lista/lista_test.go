package lista_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	TDALista "lista/lista"
	"testing"
)

func TestListaVacia(t *testing.T) {
	t.Log("Pruebas con lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestVerInserciones(t *testing.T) {
	t.Log("Pruebas con varias inserciones checkeando el valor correspondiente")
	lista := TDALista.CrearListaEnlazada[string]()

	// Inserto y veo al principio
	lista.InsertarPrimero("3er")
	require.EqualValues(t, "3er", lista.VerPrimero())
	lista.InsertarPrimero("2do")
	require.EqualValues(t, "2do", lista.VerPrimero())
	lista.InsertarPrimero("1ro")
	require.EqualValues(t, "1ro", lista.VerPrimero())
	require.EqualValues(t, "3er", lista.VerUltimo())

	// Inserto y veo al final
	lista.InsertarUltimo("4to")
	require.EqualValues(t, "4to", lista.VerUltimo())
	lista.InsertarUltimo("5to")
	require.EqualValues(t, "5to", lista.VerUltimo())
	lista.InsertarUltimo("6to")
	require.EqualValues(t, "6to", lista.VerUltimo())
	require.EqualValues(t, "1ro", lista.VerPrimero())
}

func TestVolumen(t *testing.T) {
	t.Log("Pruebas de volumen")
	const _VOLUMEN = 10000
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i <= _VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, _VOLUMEN+1, lista.Largo())

	for i := 0; i <= _VOLUMEN; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestCondicionBorde(t *testing.T) {
	t.Log("Pruebas condicion borde")
	lista := TDALista.CrearListaEnlazada[float64]()

	// Insertar al principio y final en combinacion borrar y volver a insertar
	lista.InsertarPrimero(1.0)
	lista.InsertarUltimo(2.0)
	require.EqualValues(t, 1.0, lista.BorrarPrimero())
	lista.InsertarUltimo(3.0)
	require.EqualValues(t, 2.0, lista.BorrarPrimero())
	require.EqualValues(t, 3.0, lista.BorrarPrimero())
	lista.InsertarUltimo(0.5)
	require.EqualValues(t, 0.5, lista.VerUltimo())
	require.EqualValues(t, 0.5, lista.VerPrimero())
	require.EqualValues(t, 0.5, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestIteradorInterno(t *testing.T) {
	t.Log("Pruebas con iterador interno")
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(0)

	// Hasta el final
	numeroActual := 0
	ptrNumeroActual := &numeroActual
	lista.Iterar(func(valor int) bool {
		require.EqualValues(t, *ptrNumeroActual, valor)
		*ptrNumeroActual++
		return true
	})

	// Hasta condicion
	*ptrNumeroActual = 0
	ultimoValorVisto := 0
	ptrUltimoValorVisto := &ultimoValorVisto
	lista.Iterar(func(valor int) bool {
		require.EqualValues(t, *ptrNumeroActual, valor)
		*ptrNumeroActual++
		*ptrUltimoValorVisto = valor
		return *ptrNumeroActual < 2
	})
	require.EqualValues(t, 1, *ptrUltimoValorVisto)
}

// Pruebas de Iteradores Externos

func TestListaVaciaConIterador(t *testing.T) {
	t.Log("aseguramos de que nusetros iteradores lanzan exceptiones")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestInsertarFinalConIterador(t *testing.T) {
	t.Log("Insertamos datos al final de una lista con un iterador externo")
	const (
		LISTA_LARGO_INICIAL = 3
		LISTA_LARGO_FINAL   = 2 * LISTA_LARGO_INICIAL
	)
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()

	for i := 1; i < LISTA_LARGO_INICIAL+1; i++ {
		require.Equal(t, i, iter.Siguiente())
	}

	for i := 4; i < LISTA_LARGO_FINAL+1; i++ {
		iter.Insertar(i)
		require.Equal(t, i, iter.Siguiente())
	}

	require.False(t, lista.EstaVacia())

	i := 1
	lista.Iterar(func(valor int) bool {
		require.Equal(t, i, valor)
		i++
		return true
	})

}

func TestDatosDeLista(t *testing.T) {
	t.Log("Revisamos todo los datos de la lista utilizando un iterador externo")
	const LISTA_LARGO = 6
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	iter := lista.Iterador()
	for i := 1; i < LISTA_LARGO+1; i++ {
		// Decimos que cuando estamos en el ultimo dato (6) hay siguiente
		// porque es nil, y seria el verdadero final (obvio nil no tiene siguiente)
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.Siguiente())
	}
	require.False(t, iter.HaySiguiente())
}

func TestBorrarOrdenadamente(t *testing.T) {
	t.Log("Borramos ordenadamente todo los datos de una lista utilizando un iterador externo ")
	const LISTA_LARGO = 6
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	iter := lista.Iterador()

	for i := 1; i < LISTA_LARGO+1; i++ {
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.Borrar())
	}
	require.True(t, lista.EstaVacia())
	require.False(t, iter.HaySiguiente())
}

func TestBorrarEnMedio(t *testing.T) {
	t.Log("Borramos algunos datos en el medio de la lista utiliznado un iterador externo")
	const (
		NODO_A_BORRAR        = 4
		SIGUIENTE_A_BORRAR   = NODO_A_BORRAR + 1
		SIGUIENTE_AL_BORRADO = SIGUIENTE_A_BORRAR + 1
	)

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	iter := lista.Iterador()
	for i := 1; i < NODO_A_BORRAR; i++ {
		require.Equal(t, i, iter.Siguiente())
	}

	require.Equal(t, NODO_A_BORRAR, iter.VerActual())
	require.Equal(t, NODO_A_BORRAR, iter.Borrar())
	require.Equal(t, SIGUIENTE_A_BORRAR, iter.Borrar())
	require.Equal(t, SIGUIENTE_AL_BORRADO, iter.VerActual())

}

func TestInsertarEnMedioConIterador(t *testing.T) {
	t.Log("Insertamos valores en el medio de una lista con un iterador externo")
	const (
		DATO_NUEVO_INSERTADO = 3
	)

	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(4)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(DATO_NUEVO_INSERTADO)

	require.False(t, lista.EstaVacia())

	i := 1
	lista.Iterar(func(valor int) bool {
		require.Equal(t, i, valor)
		i++
		return true
	})

}

func TestOne(t *testing.T) {
	// Pruebas durante development - NO INCLUIR -
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(8)
	lista.InsertarPrimero(0)

	for iter := lista.Iterador(); iter.HaySiguiente(); {
		if iter.VerActual()%2 == 0 {
			iter.Borrar()
		} else {
			iter.Siguiente()
		}
	}
	lista.Iterar(func(valor int) bool {
		fmt.Println(valor)
		return true
	})
}

func TestTwo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iter := lista.Iterador()

	iter.Insertar(2)

	lista.Iterar(func(valor int) bool {
		fmt.Println(valor)
		return true
	})
}
