package votos

import "rerepolez/errores"

type votanteImplementacion struct {
	dni        int
	votoActual Voto
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = checkearDni(dni)
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}

func checkearDni(dni int) int {
	// TODO Revisar si tirar error
	if dni < 0 {
		err := new(errores.DNIError)
		panic(err.Error())
	} else if false { // TODO si no esta en padron
		err := new(errores.DNIFueraPadron)
		panic(err.Error())
	}
	return dni
}
