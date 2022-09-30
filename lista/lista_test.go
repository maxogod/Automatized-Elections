package lista_test

import (
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
