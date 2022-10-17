package votos

import "fmt"

type partidoImplementacion struct {
	nombre      string
	candidatos  [CANT_VOTACION]string
	numeroVotos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosBlancos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearPartidoEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.numeroVotos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.numeroVotos[tipo] == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido.nombre, partido.candidatos[tipo], partido.numeroVotos[tipo])
	}
	return fmt.Sprintf("%s - %s: %d votos", partido.nombre, partido.candidatos[tipo], partido.numeroVotos[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosBlancos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votosBlancos[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto", blanco.votosBlancos[tipo])
	}
	return fmt.Sprintf("Votos en Blanco: %d votos", blanco.votosBlancos[tipo])
}

func VotosImpugnados() string {
	if votosImpugnados == 1 {
		return fmt.Sprintf("Votos Impugnados: %d voto", votosImpugnados)
	}
	return fmt.Sprintf("Votos Impugnados: %d votos", votosImpugnados)
}
