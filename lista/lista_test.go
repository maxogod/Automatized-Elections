package lista_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"lista/lista"
	"testing"
)

func Test(t *testing.T) {
	lista := lista.CrearListaEnlazada()
	fmt.Print(lista)
	require.True(t, true)
}
