package votos

import (
	"rerepolez/errores"
	"rerepolez/pila"
	"strings"
)

type votanteImplementacion struct {
	dni         int
	voto        Voto
	estadoVoto  bool
	pilaDeVotos pila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.pilaDeVotos = pila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa, lenPartidos int) error {
	if alternativa < 0 || alternativa > lenPartidos {
		return new(errores.ErrorAlternativaInvalida)
	} else if alternativa == 0 {
		votante.voto.Impugnado = true

	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}
	votante.pilaDeVotos.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.pilaDeVotos.EstaVacia() {
		return new(errores.ErrorNoHayVotosAnteriores)
	}
	votante.pilaDeVotos.Desapilar()
	if votante.pilaDeVotos.EstaVacia() {
		const VALOR_BASE = 0
		votante.voto.Impugnado = false
		votante.voto.VotoPorTipo = [CANT_VOTACION]int{VALOR_BASE, VALOR_BASE, VALOR_BASE}
	} else {
		votante.voto = votante.pilaDeVotos.VerTope()
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto(partido *[]Partido) error {
	if votante.estadoVoto {
		err := errores.ErrorVotanteFraudulento{Dni: votante.dni}
		return err
	}
	votante.estadoVoto = true
	guardarVoto(votante.pilaDeVotos.VerTope().VotoPorTipo, partido)
	return nil
}

func ConvertirTipoVoto(candidato string) TipoVoto {
	switch strings.ToUpper(candidato) {
	case "PRESIDENTE":
		return PRESIDENTE
	case "GOBERNADOR":
		return GOBERNADOR
	case "INTENDENTE":
		return INTENDENTE
	default:
		panic(new(errores.ErrorTipoVoto).Error())
	}
	return NONE
}

func CheckearDniValido(dni int, padron []Votante) error {
	if dni < 0 || dni > 60000000 {
		return new(errores.DNIError)
	}
	for _, documento := range padron {
		if dni == documento.LeerDNI() {
			return nil
		}
	}
	return new(errores.DNIFueraPadron)
}

func guardarVoto(votos [CANT_VOTACION]int, partidos *[]Partido) {
	i := 0
	for voto := range votos {
		partidoElegido := (*partidos)[voto]
		partidoElegido.VotadoPara(TipoVoto(i))
		i++
	}

}
