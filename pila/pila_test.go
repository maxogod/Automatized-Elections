package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Pruebas con pila inicialmente vacia")
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestLifo(t *testing.T) {
	t.Log("Pruebas de LIFO")
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar("1st")
	require.EqualValues(t, "1st", pila.VerTope())
	pila.Apilar("2nd")
	require.EqualValues(t, "2nd", pila.VerTope())
	pila.Apilar("3rd")
	require.EqualValues(t, "3rd", pila.VerTope())

	require.EqualValues(t, "3rd", pila.Desapilar())
	require.EqualValues(t, "2nd", pila.Desapilar())
	require.EqualValues(t, "1st", pila.Desapilar())
}

func TestVolumen(t *testing.T) {
	t.Log("Pruebas de volumen")
	const _VOLUMEN = 10000
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= _VOLUMEN; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	require.False(t, pila.EstaVacia())

	for i := _VOLUMEN; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestCondicionesBorde(t *testing.T) {
	t.Log("Pruebas de borde")
	pila := TDAPila.CrearPilaDinamica[float64]()

	// Apila desapila y vuelve a apilar sin estar vacia
	pila.Apilar(24.546)
	pila.Apilar(0.1)
	require.EqualValues(t, 0.1, pila.Desapilar())
	pila.Apilar(9.5)
	require.EqualValues(t, 9.5, pila.VerTope())
	require.EqualValues(t, 9.5, pila.Desapilar())
	require.EqualValues(t, 24.546, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}
