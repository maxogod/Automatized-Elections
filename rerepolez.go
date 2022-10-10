package main

import (
	"bufio"
	"fmt"
	"os"
	COLA "rerepolez/cola"
	"rerepolez/errores"
	PA "rerepolez/procesarDatos"
	V "rerepolez/votos"
	"strconv"
	"strings"
)

const (
	COMANDO_INGRESADO = iota
	PRIMER_PARAMETRO
	SEGUNDO_PARAMETRO
	OK        = "OK"
	VACIO     = ""
	SEPARADOR = " "
)

var (
	ARGS       = os.Args[PRIMER_PARAMETRO:]
	CANDIDATOS = [V.CANT_VOTACION]string{"Presidente", "Gobernador", "Intendente"}
)

func main() {
	partidos, padron, errLectura := PA.ProcesarArchivos(ARGS)
	if mostrarError(errLectura, VACIO) {
		return
	}
	colaVotantes := COLA.CrearColaEnlazada[V.Votante]()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		entrada := strings.Split(scan.Text(), SEPARADOR)
		comando := entrada[COMANDO_INGRESADO]

		switch comando {
		case "ingresar":
			dni, _ := strconv.Atoi(entrada[PRIMER_PARAMETRO])
			indiceEnPadron, errorDni := V.CheckearDniValido(dni, padron)
			if errorDni == nil {
				colaVotantes.Encolar(padron[indiceEnPadron])
			}
			mostrarError(errorDni, OK)
			break

		case "votar":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			tipoVoto, errConversion := V.ConvertirTipoVoto(entrada[PRIMER_PARAMETRO])
			if mostrarError(errConversion, VACIO) || !esNumerico(entrada[SEGUNDO_PARAMETRO]) {
				break
			}
			nroLista, _ := strconv.Atoi(entrada[SEGUNDO_PARAMETRO])

			votanteActual := colaVotantes.VerPrimero()
			errAlternativa, fraudulento := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			if fraudulento {
				colaVotantes.Desencolar()
			}
			mostrarError(errAlternativa, OK)
			break

		case "deshacer":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errDeshacer, fraudulento := (colaVotantes.VerPrimero()).Deshacer()
			if fraudulento {
				colaVotantes.Desencolar()
			}
			mostrarError(errDeshacer, OK)
			break

		case "fin-votar":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errFraudulento := colaVotantes.Desencolar().FinVoto(&partidos)
			mostrarError(errFraudulento, OK)

		case "seAprueba?":
			fmt.Println("Obvio papa promedio 10 :) xd")

		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
	if !colaVotantes.EstaVacia() {
		fmt.Println(new(errores.ErrorCiudadanosSinVotar))
	}
	salida(partidos)
}

func salida(partidos []V.Partido) {
	for tipoVoto, candidato := range CANDIDATOS {
		fmt.Printf("%s:\n", candidato)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(V.TipoVoto(tipoVoto)))
		}
		fmt.Println()
	}
	fmt.Println(V.VotosImpugnados())
}

func mostrarError(err error, alternativa string) bool {
	if err != nil {
		fmt.Println(err)
		return true
	} else if alternativa != VACIO {
		fmt.Println(alternativa)
	}
	return false
}

func esNumerico(cadena string) bool {
	const TAMANIO_BIT = 64
	_, err := strconv.ParseFloat(cadena, TAMANIO_BIT)
	if err != nil {
		fmt.Println(new(errores.ErrorAlternativaInvalida))
	}
	return err == nil
}
