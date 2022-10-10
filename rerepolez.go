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

func main() {
	partidos, padron, errLectura := PA.ProcesarArchivos(ARGS)
	if mostrarError(errLectura, "") {
		return
	}
	colaVotantes := COLA.CrearColaEnlazada[V.Votante]()

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		entrada := strings.Split(scan.Text(), " ")

		commndo := entrada[0]

		switch commndo {
		case "ingresar":
			dni, _ := strconv.Atoi(entrada[1]) // O(n)
			indiceEnPadron, errorDni := V.CheckearDniValido(dni, padron)

			if errorDni == nil {
				colaVotantes.Encolar(padron[indiceEnPadron])
			}
			mostrarError(errorDni, "OK")
			break
		case "votar":
			if colaVotantes.EstaVacia() {
				fmt.Println(new(errores.FilaVacia))
				break
			}
			tipoVoto, errConversion := V.ConvertirTipoVoto(entrada[1])
			if mostrarError(errConversion, "") {
				break
			}
			if !isNumeric(entrada[2]) {
				break
			}
			nroLista, _ := strconv.Atoi(entrada[2])

			votanteActual := colaVotantes.VerPrimero()
			errAlternativa, fraudulento := votanteActual.Votar(tipoVoto, nroLista, len(partidos))
			if fraudulento {
				colaVotantes.Desencolar()
			}
			mostrarError(errAlternativa, "OK")

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
			mostrarError(errDeshacer, "OK")
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

func votanteFradulento(yaVotaron []V.Votante, dni int) bool {
	for _, votante := range yaVotaron {
		if dni == votante.LeerDNI() {
			fmt.Println(errores.ErrorVotanteFraudulento{Dni: dni})
			return true
		}
	}
	return false
}

func mostrarError(err error, alternativa string) bool {
	if err != nil {
		fmt.Println(err)
		return true
	} else if alternativa != "" {
		fmt.Println(alternativa)
	}
	return false
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(new(errores.ErrorAlternativaInvalida))
	}
	return err == nil
}
