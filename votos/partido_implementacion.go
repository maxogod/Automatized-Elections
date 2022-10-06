package votos

type partidoImplementacion struct {
	nombre      string
	candidatos  [CANT_VOTACION]string
	numeroVotos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosBlancos    [CANT_VOTACION]int
	votosImpugnados int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.numeroVotos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return ""
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosBlancos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
