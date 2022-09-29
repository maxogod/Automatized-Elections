package lista_test

import (
	"github.com/stretchr/testify/require"
	"lista/lista"
	"testing"
)

func Test(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "che loco esta vacio", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "che loco esta vacio", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}
