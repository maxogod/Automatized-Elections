package votos

import (
	"rerepolez/errores"
	"rerepolez/pila"
	"strings"
)

type votanteImplementacion struct {
	dni            int
	voto           Voto
	yaVoto         bool
	historialVotos pila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.historialVotos = pila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa, lenPartidos int) (error, bool) {
	if votante.yaVoto {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}, FRAUDULENTO
	}

	if alternativa < 0 || alternativa >= lenPartidos {
		return new(errores.ErrorAlternativaInvalida), !FRAUDULENTO
	} else if alternativa == IMPUGNADO {
		votante.voto.Impugnado = true
	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}

	votante.historialVotos.Apilar(votante.voto)
	return nil, !FRAUDULENTO
}

func (votante *votanteImplementacion) Deshacer() (error, bool) {
	if votante.yaVoto {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}, FRAUDULENTO
	}

	if votante.historialVotos.EstaVacia() {
		return new(errores.ErrorNoHayVotosAnteriores), !FRAUDULENTO
	}
	votante.historialVotos.Desapilar()

	if votante.historialVotos.EstaVacia() {
		const VALOR_BASE = 0
		votante.voto.Impugnado = false
		votante.voto.VotoPorTipo = [CANT_VOTACION]int{VALOR_BASE, VALOR_BASE, VALOR_BASE}
	} else {
		votante.voto = votante.historialVotos.VerTope()
	}

	return nil, !FRAUDULENTO
}

func (votante *votanteImplementacion) FinVoto(partido *[]Partido) error {
	if votante.yaVoto {
		err := errores.ErrorVotanteFraudulento{Dni: votante.dni}
		return err
	}

	if votante.voto.Impugnado {
		votosImpugnados++
	} else if votante.historialVotos.EstaVacia() {
		// Guardar voto en blanco
		guardarVoto(votante.voto.VotoPorTipo, partido)
	} else {
		guardarVoto(votante.historialVotos.VerTope().VotoPorTipo, partido)
	}

	votante.yaVoto = true
	return nil
}

func ConvertirTipoVoto(candidato string) (TipoVoto, error) {
	switch strings.ToUpper(candidato) {
	case "PRESIDENTE":
		return PRESIDENTE, nil
	case "GOBERNADOR":
		return GOBERNADOR, nil
	case "INTENDENTE":
		return INTENDENTE, nil
	default:
		return NINGUNO, new(errores.ErrorTipoVoto)
	}
}

func guardarVoto(votos [CANT_VOTACION]int, partidos *[]Partido) {
	for tipo, alternativa := range votos {
		partidoElegido := (*partidos)[alternativa]
		partidoElegido.VotadoPara(TipoVoto(tipo))
	}
}
