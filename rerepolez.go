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
	CANDIDATOS = [3]string{"Presidente", "Gobernador", "Intendente"}
)

// go build ~ ./rerepolez <archivo partidos> <archivo padron>

// go build rerepolez
// ./rerepolez tests/01_partidos tests/02_padron

func main() {
	partidos, padron, errLectura := PA.ProcesarArchivos(ARGS)
	if errLectura {
		fmt.Println(new(errores.ErrorLeerArchivo))
		return
	}
	colaVotantes := COLA.CrearColaEnlazada[V.Votante]()
	yaVotaron := make([]V.Votante, 0)

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		entrada := strings.Split(scan.Text(), " ")

		commndo := entrada[0]

		switch commndo {
		case "ingresar":
			dni, _ := strconv.Atoi(entrada[1])
			errorDni := V.CheckearDniValido(dni, padron)

			if errorDni == nil {
				if votanteFradulento(yaVotaron, dni) {
					// TODO sacarlo de padron
					break
				}
				votanteActual := V.CrearVotante(dni)
				colaVotantes.Encolar(votanteActual)
				yaVotaron = append(yaVotaron, votanteActual)
			}
			mostrarError(errorDni, "OK")
			break
		case "votar":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			tipoVoto, errConversion := V.ConvertirTipoVoto(entrada[1])
			mostrarError(errConversion, "")
			nroLista, _ := strconv.Atoi(entrada[2])

			votanteActual := colaVotantes.VerPrimero()
			errAlternativa := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			mostrarError(errAlternativa, "OK")

			break
		case "deshacer":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errSinAnterior := (colaVotantes.VerPrimero()).Deshacer()
			mostrarError(errSinAnterior, "OK")
			break
		case "fin-votar":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			errFraudulento := colaVotantes.Desencolar().FinVoto(&partidos)
			mostrarError(errFraudulento, "OK")
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
	salida(partidos)
}

func salida(partidos []V.Partido) {
	for tipoVoto, candidato := range CANDIDATOS {
		fmt.Printf("%s : \n", candidato)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(V.TipoVoto(tipoVoto)))
		}
		fmt.Println()
	}
	fmt.Println(V.VotosImpugnados())
}

func votanteFradulento(yaVotaron []V.Votante, dni int) bool {
	for _, votante := range yaVotaron {
		if dni == votante.LeerDNI() {
			fmt.Println(errores.ErrorVotanteFraudulento{Dni: dni})
			return true
		}
	}
	return false
}

func mostrarError(err error, alternativa string) {
	if err != nil || alternativa == "" {
		fmt.Println(err)
	} else {
		fmt.Println(alternativa)
	}
}
