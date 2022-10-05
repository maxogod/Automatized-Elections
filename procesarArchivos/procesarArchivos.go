package procesarArchivos

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"strconv"
	"strings"
)

func procesarPartidos(archivoNombre string) [][]string {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		error_ := new(errores.ErrorLeerArchivo)
		panic(error_.Error())
	}
	defer archivo.Close()

	arr := make([][]string, 0)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		arr = append(arr, strings.Split(strings.TrimSuffix(s.Text(), "\n"), ","))
	}
	return arr
}

func procesarPadron(archivoNombre string) []int {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		error_ := new(errores.ErrorLeerArchivo)
		panic(error_.Error())
	}
	defer archivo.Close()

	arr := make([]int, 0)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		padron, _ := strconv.Atoi(s.Text())
		arr = append(arr, padron)
	}
	return arr
}

func ProcesarArchivos(args []string) ([][]string, []int) {
	if len(args) != 2 {
		err := new(errores.ErrorLeerArchivo)
		panic(err.Error())
	}
	const (
		PARTIDOS_POS = 0
		PADRON_POS   = 1
	)
	partidos := procesarPartidos(args[PARTIDOS_POS])
	padron := procesarPadron(args[PADRON_POS])
	return partidos, padron
}
