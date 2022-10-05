package main

import (
	"fmt"
	"os"
	PA "rerepolez/procesarArchivos"
)

var ARGS = os.Args[1:]

// go build ~ ./rerepolez <archivo partidos> <archivo padron>

func main() {
	arr1, arr2 := PA.ProcesarArchivos(ARGS)
	fmt.Println(arr1, arr2)
}
