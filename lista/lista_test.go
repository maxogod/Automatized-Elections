package lista_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"lista/lista"
	"testing"
)

func Test(t *testing.T) {
	lista := lista.CrearListaEnlazada[int]()
	//lista.InsertarPrimero(1)
	//lista.InsertarPrimero(2)
	//lista.InsertarPrimero(3)

	lista.InsertarUltimo(123)
	lista.InsertarUltimo(23)
	lista.InsertarUltimo(0)

	fmt.Print(lista)

	require.True(t, true)
}
