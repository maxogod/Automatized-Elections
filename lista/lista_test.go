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

// Pruebas de Iteradores Externos

func TestListaVaciaConIterador(t *testing.T) {
	t.Log("aseguramos de que nusetros iteradores lanzan exceptiones")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestInsertarAlprincipio(t *testing.T) {
	t.Log("Insertamos estando en el principio con iterador y checkeamos que haya sido efectivamente al principio")

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)

	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(0)

	i := 0
	lista.Iterar(func(valor int) bool {
		require.Equal(t, i, valor)
		i++
		return true
	})

}

func TestInsertarAlFinal(t *testing.T) {
	t.Log("Insertamos estando en el final con iterador y checkeamos que haya sido efectivamente al final")

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)

	iter := lista.Iterador()
	iter.Siguiente() // Estamos parados en nil (osea el final)
	iter.Insertar(1)
	iter.Siguiente()
	iter.Insertar(0)

	i := 2
	lista.Iterar(func(valor int) bool {
		require.Equal(t, i, valor)
		i--
		return true
	})

}

func TestInsertarEnMedio(t *testing.T) {
	t.Log("Insertamos valores en el medio de una lista con un iterador externo")
	const (
		_DATO_NUEVO_INSERTADO = 3
	)
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(_DATO_NUEVO_INSERTADO)
	require.False(t, lista.EstaVacia())

	i := 1
	lista.Iterar(func(valor int) bool {
		require.Equal(t, i, valor)
		i++
		return true
	})

}

func TestBorrarPrimero(t *testing.T) {
	t.Log("Borramos cuando se crea el iterador (osea al principio)")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)

	iter := lista.Iterador()
	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, lista.VerPrimero(), lista.VerUltimo())
}

func TestBorrarUltimo(t *testing.T) {
	t.Log("Borramos el ultimo y nos fijamos que efectivamente se haya borrado")
	const (
		UlIMO       = 6
		ANTE_ULTIMO = 5
	)
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	iter := lista.Iterador()
	for i := 1; i < UlIMO; i++ {
		require.Equal(t, i, iter.Siguiente())
	}

	require.Equal(t, UlIMO, iter.VerActual())
	require.Equal(t, UlIMO, iter.Borrar())
	require.Equal(t, ANTE_ULTIMO, lista.VerUltimo())
}

func TestBorrarEnMedio(t *testing.T) {
	t.Log("Borramos algunos datos en el medio de la lista utilizando un iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.Siguiente())
	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 3, iter.VerActual())

	// Veo que efectivamente la lista sea [1 3]
	i := 1
	lista.Iterar(func(valor int) bool {
		require.EqualValues(t, i, valor)
		i += 2
		return true
	})
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
