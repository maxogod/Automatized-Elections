package votos

import (
	"elecciones/errores"
	"elecciones/pila"
	"strings"
)

type votanteImplementacion struct {
	dni            int
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
	votoActual := *new(Voto)
	if !votante.historialVotos.EstaVacia() {
		votoActual = votante.historialVotos.VerTope()
	}

	if alternativa < 0 || alternativa >= lenPartidos {
		return new(errores.ErrorAlternativaInvalida), !FRAUDULENTO
	} else if alternativa == IMPUGNADO {
		votoActual.Impugnado = true
	} else {
		votoActual.VotoPorTipo[tipo] = alternativa
	}

	votante.historialVotos.Apilar(votoActual)
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

	return nil, !FRAUDULENTO
}

func (votante *votanteImplementacion) FinVoto(partido *[]Partido) error {
	if votante.yaVoto {
		err := errores.ErrorVotanteFraudulento{Dni: votante.dni}
		return err
	}

	if votante.historialVotos.EstaVacia() {
		// Guardar voto en blanco
		votoBlanco := new(Voto)
		guardarVoto(votoBlanco.VotoPorTipo, partido)
	} else if votante.historialVotos.VerTope().Impugnado {
		votosImpugnados++
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
