package cola_test

import (
	ColaTDA "cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	t.Log("Pruebas con cola inicialmente vacia")
	cola := ColaTDA.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestFifo(t *testing.T) {
	t.Log("Pruebas de FIFO")
	cola := ColaTDA.CrearColaEnlazada[string]()

	cola.Encolar("1ro")
	require.EqualValues(t, "1ro", cola.VerPrimero())
	cola.Encolar("2do")
	require.EqualValues(t, "1ro", cola.Desencolar())
	require.EqualValues(t, "2do", cola.VerPrimero())
	cola.Encolar("3ro")
	require.EqualValues(t, "2do", cola.Desencolar())
	require.EqualValues(t, "3ro", cola.VerPrimero())
	require.EqualValues(t, "3ro", cola.Desencolar())
}

func TestVolumen(t *testing.T) {
	t.Log("Pruebas de volumen")
	const _VOLUMEN = 10000
	cola := ColaTDA.CrearColaEnlazada[int]()

	for i := 0; i <= _VOLUMEN; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())

	for i := 0; i <= _VOLUMEN; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestCondicionesBorde(t *testing.T) {
	t.Log("Pruebas de borde")
	cola := ColaTDA.CrearColaEnlazada[float64]()

	// Encolar, desencolar y volver a encolar sin estar vacia
	cola.Encolar(24.546)
	cola.Encolar(0.1)
	require.EqualValues(t, 24.546, cola.Desencolar())
	cola.Encolar(25)
	require.EqualValues(t, 0.1, cola.Desencolar())
	require.EqualValues(t, 25, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}
