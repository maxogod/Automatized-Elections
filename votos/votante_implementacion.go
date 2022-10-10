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
	if votante.voto.Impugnado {
		votosImpugnados++
	} else if votante.pilaDeVotos.EstaVacia() {
		guardarVoto(votante.voto.VotoPorTipo, partido)
	} else {
		guardarVoto(votante.pilaDeVotos.VerTope().VotoPorTipo, partido)
	}
	votante.estadoVoto = true
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
		return TipoVoto(0), new(errores.ErrorTipoVoto)
	}
}

func guardarVoto(votos [CANT_VOTACION]int, partidos *[]Partido) {
	for tipo, alternativa := range votos {
		partidoElegido := (*partidos)[alternativa]
		partidoElegido.VotadoPara(TipoVoto(tipo))
	}
}

func CheckearDniValido(dni int, padron []Votante) (indiceEnPadron int, err error) {
	if dni <= 0 {
		return -1, new(errores.DNIError)
	}
	indice := buscarVotanteEnPadron(dni, 0, len(padron), padron)
	if indice != -1 {
		return indice, nil
	}
	return indice, new(errores.DNIFueraPadron)
}

func buscarVotanteEnPadron(dni, ini, fin int, votantes []Votante) int {
	if ini > fin {
		// No esta :(
		return -1
	}
	medio := (ini + fin) / 2
	if votantes[medio].LeerDNI() > dni {
		// Miro izq
		return buscarVotanteEnPadron(dni, ini, medio-1, votantes)
	} else if votantes[medio].LeerDNI() < dni {
		// Miro derecha
		return buscarVotanteEnPadron(dni, medio+1, fin, votantes)
	} else {
		// Encontrado!
		return medio
	}
}
