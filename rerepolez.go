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

var (
	ARGS       = os.Args[1:]
	CANDIDATOS = [3]string{"Presidente", "Gobierno", "Intendente"}
)

// go build ~ ./rerepolez <archivo partidos> <archivo padron>

// go build rerepolez
// ./rerepolez tests/01_partidos tests/02_padron

func main() {
	partidos, padron := PA.ProcesarArchivos(ARGS)
	colaVotantes := COLA.CrearColaEnlazada[V.Votante]()

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		entrada := strings.Split(scan.Text(), " ")

		commndo := entrada[0]

		switch commndo {
		case "ingresar":
			dni, _ := strconv.Atoi(entrada[1])
			errorDni := V.CheckearDniValido(dni, padron)

			if errorDni == nil {
				colaVotantes.Encolar(V.CrearVotante(dni))
			}
			mostrarError(errorDni, "OK")
			break
		case "votar":
			tipoVoto, errConversion := V.ConvertirTipoVoto(entrada[1])
			mostrarError(errConversion, "")
			nroLista, _ := strconv.Atoi(entrada[2])

			errorFilaVacia(colaVotantes)
			votanteActual := colaVotantes.VerPrimero()
			errAlternativa := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			mostrarError(errAlternativa, "OK")

			break
		case "deshacer":
			errSinAnterior := (colaVotantes.VerPrimero()).Deshacer()
			mostrarError(errSinAnterior, "OK")
			break
		case "fin-votar":
			errorFilaVacia(colaVotantes)
			errFraudulento := colaVotantes.Desencolar().FinVoto(&partidos)
			mostrarError(errFraudulento, "OK")
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
	fmt.Println("+++SALIDA+++")
	salida(partidos)
}

func salida(partidos []V.Partido) {
	for tipoVoto, candidato := range CANDIDATOS {
		fmt.Printf("%s : \n", candidato)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(V.TipoVoto(tipoVoto)))
		}
	}
	fmt.Println(V.VotosImpugnados())
}

func errorFilaVacia(cola COLA.Cola[V.Votante]) {
	if cola.EstaVacia() {
		fmt.Println(new(errores.FilaVacia))
	}
}

func mostrarError(err error, alternativa string) {
	if err != nil || alternativa == "" {
		fmt.Println(err)
	} else {
		fmt.Println(alternativa)
	}
}
