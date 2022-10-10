package procesarDatos

import (
	"bufio"
	"os"
	"rerepolez/errores"
	"rerepolez/votos"
	"strconv"
	"strings"
)

func procesarPartidos(archivoNombre string) ([]votos.Partido, bool) {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		return nil, true
	}
	defer archivo.Close()

	const (
		NOMBRE_PARTIDO_POS = 0
		SEPARADOR          = ","
	)
	partidos := []votos.Partido{votos.CrearVotosEnBlanco()} // Inicializado con el partido en blanco en pos 0
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := strings.Split(strings.TrimSuffix(s.Text(), "\n"), SEPARADOR)
		nombrePartido := linea[NOMBRE_PARTIDO_POS]
		candidatos := [votos.CANT_VOTACION]string{linea[votos.PRESIDENTE+1], linea[votos.GOBERNADOR+1], linea[votos.INTENDENTE+1]}
		partidos = append(partidos, votos.CrearPartido(nombrePartido, candidatos))
	}
	return partidos, false
}

func procesarPadron(archivoNombre string) ([]votos.Votante, bool) {
	archivo, err := os.Open(archivoNombre)
	if err != nil {
		return nil, true
	}
	defer archivo.Close()

	votantes := make([]votos.Votante, 0)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		padron, _ := strconv.Atoi(s.Text())
		nuevoVotante := votos.CrearVotante(padron)
		votantes = append(votantes, nuevoVotante)
	}
	return ordenarVotantes(votantes), false
}

func ProcesarArchivos(args []string) ([]votos.Partido, []votos.Votante, error) {
	if len(args) != 2 {
		return nil, nil, new(errores.ErrorParametros)
	}
	const (
		PARTIDOS_POS = iota
		PADRON_POS
	)
	partidos, errPartidos := procesarPartidos(args[PARTIDOS_POS])
	padron, errPadron := procesarPadron(args[PADRON_POS])
	if errPadron || errPartidos {
		return nil, nil, new(errores.ErrorLeerArchivo)
	}

	return partidos, padron, nil
}
