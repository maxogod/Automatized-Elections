package main

import (
	"fmt"
	"os"
	PA "rerepolez/procesarDatos"
)

var (
	ARGS             = os.Args[1:]
	partidos, padron = PA.ProcesarArchivos(ARGS)
)

// go build ~ ./rerepolez <archivo partidos> <archivo padron>

func main() {
	fmt.Println(partidos, padron)
}
