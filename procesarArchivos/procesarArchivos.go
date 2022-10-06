package procesarArchivos

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
)

func procesarPartidos(archivoNombre string) []votos.Partido {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		error_ := new(errores.ErrorLeerArchivo)
		panic(error_.Error())
	}
	defer archivo.Close()

	const NOMBRE_PARTIDO_POS = 0
	partidos := []votos.Partido{votos.CrearVotosEnBlanco()} // Inicializado con el partido en blanco en pos 0
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := strings.Split(strings.TrimSuffix(s.Text(), "\n"), ",")
		nombrePartido := linea[NOMBRE_PARTIDO_POS]
		candidatos := [votos.CANT_VOTACION]string{linea[votos.PRESIDENTE+1], linea[votos.GOBERNADOR+1], linea[votos.INTENDENTE+1]}
		partidos = append(partidos, votos.CrearPartido(nombrePartido, candidatos))
	}
	return partidos
}

func procesarPadron(archivoNombre string) []int {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		error_ := new(errores.ErrorLeerArchivo)
		panic(error_.Error())
	}
	defer archivo.Close()

	dnis := make([]int, 0)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		padron, _ := strconv.Atoi(s.Text())
		dnis = append(dnis, padron)
	}
	return dnis
}

func ProcesarArchivos(args []string) ([]votos.Partido, []int) {
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
