package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errores"
	PA "rerepolez/procesarDatos"
)

var (
	ARGS = os.Args[1:]
)

// go build ~ ./rerepolez <archivo partidos> <archivo padron>

func main() {
	partidos, padron := PA.ProcesarArchivos(ARGS)
	fmt.Println(partidos, padron)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		switch scan.Text() {
		case "ingresar":
			fmt.Println("hallo")
			break
		case "votar":
			fmt.Println("a")
			break
		case "deshacer":
			fmt.Println("b")
			break
		case "fin-votar":
			fmt.Println("bye")
			break
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
	fmt.Println("resultados")
}
