package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/errores"
	pa "rerepolez/procesarDatos"
	v "rerepolez/votos"
	"strconv"
	"strings"
)

const (
	INGRESAR  = "ingresar"
	VOTAR     = "votar"
	DESHACER  = "deshacer"
	FIN_VOTAR = "fin-votar"
	SEPARADOR = " "
)

var (
	ARGS = os.Args[1:]
)

func main() {
	partidos, padron, errLectura := pa.ProcesarArchivos(ARGS)
	if error_Y_Bool(errLectura) {
		return
	}
	votantes := cola.CrearColaEnlazada[v.Votante]()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		entrada := strings.Split(scan.Text(), SEPARADOR)
		comando := entrada[0]

		switch comando {
		case INGRESAR:
			dni, _ := strconv.Atoi(entrada[1])
			// Atoi devuelve 0 si no casteable a int lo cual es un dni invalido, por esta razon no se maneja el error
			// (en este caso)
			indiceEnPadron, errorDni := pa.CheckearDniValido(dni, padron)
			if errorDni == nil {
				votantes.Encolar(padron[indiceEnPadron])
			}
			error_O_Ok(errorDni) // mostrar al usuario si tod0 ok o tod0 mal

		case VOTAR:
			if votantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			tipoVoto, errConversion := v.ConvertirTipoVoto(entrada[1])
			nroLista, errConverAlternativa := strconv.Atoi(entrada[2])

			if errConverAlternativa != nil {
				errConverAlternativa = new(errores.ErrorAlternativaInvalida)
			}
			if error_Y_Bool(errConversion) || error_Y_Bool(errConverAlternativa) {
				break
			}

			votanteActual := votantes.VerPrimero()
			errAlVotar, fraudulento := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			if fraudulento {
				votantes.Desencolar()
			}
			error_O_Ok(errAlVotar)

		case DESHACER:
			if votantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errDeshacer, fraudulento := (votantes.VerPrimero()).Deshacer()
			if fraudulento {
				votantes.Desencolar()
			}
			error_O_Ok(errDeshacer)

		case FIN_VOTAR:
			if votantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errFraudulento := votantes.Desencolar().FinVoto(&partidos)
			error_O_Ok(errFraudulento)

		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
	if !votantes.EstaVacia() {
		fmt.Println(new(errores.ErrorCiudadanosSinVotar))
	}
	salida(partidos)
}

// salida Printea los resultados de la votacion
func salida(partidos []v.Partido) {
	tipoDeCandidatos := []string{"Presidente", "Gobernador", "Intendente"}
	for tipoVoto, candidato := range tipoDeCandidatos {
		fmt.Printf("%s:\n", candidato)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(v.TipoVoto(tipoVoto)))
		}
		fmt.Println()
	}
	fmt.Println(v.VotosImpugnados())
}

// error_O_Ok printea el error pasado en caso != nil sino printea OK.
func error_O_Ok(err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}

// error_Y_Bool printea el error pasado en caso != nil, retorna bool si hubo o no error.
func error_Y_Bool(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
