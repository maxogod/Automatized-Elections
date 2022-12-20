# Automatizacion de votaciones
<a href="./English-README.md">English DOC</a>

El programa Elecciones automatiza el proceso de votar.
Para esto en la inicializacion del programa se le dan los dnis que pueden votar en esta
maquina/ejecucion y los partidos validos por los cuales se puede votar.

## Para correr:
~~~
Ejecutar el archivo 'elecciones' o de ser necesario recompilar el package 'elecciones' 
para luego ejecutarlo.
~~~

## Comandos:
~~~
ingresar <NumeroDNI>
~~~
Permite que un votante se loguee con su dni

~~~
votar <TIPO-VOTO> <NumeroLista>
~~~
Permite que el votante agrege una parte de la boleta a su voto.
TIPO-VOTO = Presidente/Gobernador/Intendente
NumeroLista = ID del partido politico

~~~
deshacer
~~~
Permite deshacer la ultima accion

~~~
fin-votar
~~~
Cierra y envia el voto al repertorio de votos.

En caso de ser usado indebidamente el programa dara los errores correspondientes y seran visibles en
pantalla, y si ese error conlleva otra accion se guardaran datos del voto y votante, 
como en caso que haya un votante fraudulento, que no solo no se tendria en 
cuenta su voto sino que se guardaria quien 
es para que alguna decicion de que hacer con esa persona sea tomada.
