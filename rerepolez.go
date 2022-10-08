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
	ARGS = os.Args[1:]
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
			colaVotantes.Encolar(V.CrearVotante(dni))
			break
		case "votar":
			//tipoVoto := entrada[1]
			//nroLista, _ := strconv.Atoi(entrada[2])

			break
		case "deshacer":

			break
		case "fin-votar":

			break
		default:
			panic(new(errores.ErrorParametros).Error())
		}
	}
}
