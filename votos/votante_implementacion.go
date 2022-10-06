package votos

type votanteImplementacion struct {
	dni        int
	voto       *Voto
	estadoVoto bool
}

func CrearVotante(dni int) Votante {
	return nil
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
