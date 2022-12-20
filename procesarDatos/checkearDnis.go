package procesarDatos

import (
	"elecciones/errores"
	v "elecciones/votos"
)

func CheckearDniValido(dni int, padron []v.Votante) (indiceEnPadron int, err error) {
	if dni <= 0 {
		return -1, new(errores.DNIError)
	}

	return buscarVotanteEnPadron(dni, 0, len(padron), padron)
}

func buscarVotanteEnPadron(dni, ini, fin int, votantes []v.Votante) (int, error) {
	if ini > fin {
		// No esta
		return -1, new(errores.DNIFueraPadron)
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
		return medio, nil
	}
}
