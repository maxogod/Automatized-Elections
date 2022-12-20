package procesarDatos

import V "elecciones/votos"

func ordenarVotantes(votantes []V.Votante) []V.Votante {

	if len(votantes) < 2 {
		return votantes
	}
	mid := (len(votantes)) / 2
	return merge(ordenarVotantes(votantes[:mid]), ordenarVotantes(votantes[mid:]))
}

func merge(izq, der []V.Votante) []V.Votante {
	salida := make([]V.Votante, len(izq)+len(der))
	i, j, k := 0, 0, 0
	for i < len(izq) && j < len(der) {
		if izq[i].LeerDNI() <= der[j].LeerDNI() {
			salida[k] = izq[i]
			i++
		} else {
			salida[k] = der[j]
			j++
		}
		k++
	}
	for i < len(izq) {
		salida[k] = izq[i]
		i++
		k++
	}
	for j < len(der) {
		salida[k] = der[j]
		j++
		k++
	}
	return salida
}
