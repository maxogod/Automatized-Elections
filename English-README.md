# Votation Automatization

The program 'elecciones' automatizes the voting process.
For this purpose the program takes two mandatory parameters 
* a file with all the political Parties
* a file with all the valid IDs that are allowed to vote in this machine/run

## To Run:

Execute the 'elecciones' file or recompile the 'elecciones' package with the command `go build elecciones`


## Commands:
~~~
ingresar <NumeroDNI>
~~~
  Allows for a voter to log in with their ID

~~~
votar <TIPO-VOTO> <NumeroLista>
~~~
Allows for a voter to fill their vote.

~~~
deshacer
~~~
Allows a voter to undo their last action

~~~
fin-votar
~~~
Closes and submits the vote to the temporary database

In case the program is used inappropriately it will display the corresponding errors and save the data of the vote and voter if necessary (e.g. if an illegal vote is submited)
