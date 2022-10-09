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
	fmt.Println(partidos)
	fmt.Println(padron)

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
				OK()
			} else {
				panic(errorDni)
			}
			break
		case "votar":
			tipoVoto := V.ConvertirTipoVoto(entrada[1])
			nroLista, _ := strconv.Atoi(entrada[2])
			votanteActual := colaVotantes.VerPrimero()
			err := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			if err != nil {
				fmt.Println(err)
			} else {
				OK()
			}

			break
		case "deshacer":
			err := (colaVotantes.VerPrimero()).Deshacer()
			if err != nil {
				fmt.Println(err)
			} else {
				OK()
			}
			break
		case "fin-votar":
			err := colaVotantes.Desencolar().FinVoto(&partidos)
			if err != nil {
				fmt.Println(err)
			} else {
				OK()
			}
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
		fmt.Println("+++SALIDA+++")
		salida(partidos)

	}
}

func salida(partidos []V.Partido) {
	for candidato := range CANDIDATOS {
		fmt.Printf("%s : \n", candidato)

	}

}

func OK() {
	fmt.Println("OK")
}
