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
			errorDni := V.CheckearDniValido(dni, padron)

			if errorDni == nil {
				colaVotantes.Encolar(V.CrearVotante(dni))
				fmt.Println("OK")
			} else {
				fmt.Println(errorDni)
			}
			break
		case "votar":
			//tipoVoto := entrada[1]
			//nroLista, _ := strconv.Atoi(entrada[2])

			break
		case "deshacer":
			errorDeshacer := (colaVotantes.VerPrimero()).Deshacer()
			if errorDeshacer != nil {
				println(errorDeshacer)
				// TODO no c pq printea un pos de memoria en vez del error (esta hecha igual a ~ingresar~)
			}
			break
		case "fin-votar":
			voto, err := colaVotantes.Desencolar().FinVoto()
			if err != nil {
				fmt.Println(err)
			}
			voto.Impugnado = true // Borrar linea, es solo para usar voto y compilar.
			// TODO hacer algo con voto
			break
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
}
