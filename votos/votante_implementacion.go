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
	}
	votante.voto = votante.pilaDeVotos.VerTope()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.estadoVoto {
		err := new(errores.ErrorVotanteFraudulento)
		return Voto{}, err

	} else {
		votante.estadoVoto = true
		return votante.voto, nil
	}

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
