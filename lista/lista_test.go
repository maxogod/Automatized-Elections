package lista_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"lista/lista"
	"testing"
)

func TestListaEnlazada(t *testing.T) {
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

func TestIteradoresExternos(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	//iter := lista.Iterador()
	//for n := 0; n < 3; n++ {
	//	iter.Insertar(n)
	//	iter.Siguiente()
	//}
	count := 0
	ptrCount := &count
	lista.Iterar(func(v int) bool {
		*ptrCount++
		fmt.Println(v)
		return *ptrCount < 5
	})
}
