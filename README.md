Présentation
------------
Ce projet fait parti du cours d'INFO601 de l'Université Savoie Mont Blanc. </br>
Le but est de résoudre le problème du site [Advent of Code du 22/12/2021](https://adventofcode.com/2021/day/22). </br>
Tuteurs : Pierre Hyvernat et Valentin Gledel

Prérequis
------------
* Installer un éditeur de go
* Télécharger tous les fichiers de ce dépot dans un dossier 

Utilisation
------------
1. Modifier l'input
Pour lancer le code qui vous donnera la réponse au problème, il vous faut tout d'abord modifier le nom du fichier de votre input. </br>
Pour cela, ouvrez le fichier 'adventCode.go' et modifiez la variable 'filename' à la ligne 11 pour renseigner le nom et le chemin de votre fichier d'input. </br>

2. Lancer le programme 
Pour lancer le programme, ouvrer un terminal, allez dans le dossier contenant tous les fichiers de ce dépot et lancez la commande 
```
go run .
```
Vous trouverez comme résultat dans le terminal un résultat tel que celui ci : 
```
partie 1:
570815
partie 2:
1268312839438137
```
Vous trouverez donc le résultat pour la première partie et la deuxième partie du problème. </br>

3. Lancer le benchmark 
Pour lancer le benchmark lié au programme, lancez la commande 
```
go test -bench=Part
```
Vous trouverez en résultat le nombre de fois que les tests ont été effctué, le nombre de nanosecondes écoulés par opérations, et le temps total de l'exécution du programme. 

Explication du code
------------
Dans cette partie voous trouverez plus d'explications à propos du fonctionnement de l'algorithme en lui même. </br>
Le code est ici expliqué en langage naturel. </br>
Dans la première partie, le principe est d'utiliser une matrice à trois dimensions pour représenter le réactor de -50 à 50 en x, en y et en z. </br>
Dans la deuxième partie, le principe est d'ajouter des pavés et notamment des intersections de pavés dans un tableau pour pouvoir par la suite ajouter ou enlever les valeurs des pavés 'vues' pour compter combien d'interrupteurs sont à 'on' à la fin. </br>

1. Partie 1 
```
Ouvrir le fichier de l'input 
Création d'une matrice à 3 dimensions de booléen 
Initialisation de cette matrice : 
    Pour chaque case de cette matrice, lui associer la valeur faux
Pour chaque ligne de cet input faire : 
    récupérer la ligne 
    récupérer à partir de cette ligne le mot (on/off), le x minimum et maximum, le y minimum et maximum, le z minimum et maximum
    Si le pavé récupéré est dans la zone d'initialisation (les valeurs minimum et maximum de ce pavé sont compris dans la zone d'initialisation) faire : 
        limiter les valeurs de x, y et z du pavé lut pour ne prendre que l'intersection avec le pavé d'initialisation, additionné à 50 pour avoir l'indice dans la matrice de booléen
        Pour chaque valeur x du pavé faire : 
            Pour chaque valeur y du pavé faire : 
                Pour chaque valeur z du pavé faire : 
                    Si le mot est 'off' alors faire : 
                        associer la valeur de la matrice en (x,y,z) à faux
                    Sinon faire :
                        associer la valeur de la matrice en (x,y,z) à vrai
récupérer le nombre de cases dasn la matrice qui sont a vrai : 
    Pour chaque case de cette matrice, augmenter un compteur si la valeur de cette case est vrai 
Retourner ce nombre </br>
```

2. Partie 2
```
Ouvrir le fichier de l'input 
Création d'un tableau de Pave (structure avec les coordonnées minimum et maximum en x, y et z, et un booléen qui permet de savoir si le pavé met ses valeurs à 'on' (vrai) ou à 'off' (faux))
Pour chaque ligne de cet input faire : 
    récupérer la ligne 
    récupérer à partir de cette ligne le mot (on/off), le x minimum et maximum, le y minimum et maximum, le z minimum et maximum
    Si le mot est 'on' faire : 
        créer un nouveau pavé avec les informations du pavé récupérés et en mettant le booléen à vrai
    Sinon faire :
        créer un nouveau pavé avec les informations du pavé récupérés et en mettant le booléen à faux
    Pour chaque indice dans le tableau de pavé faire :
        récupérer le pavé à cet indice 
        rechercher l'intersection du pavé lut avec le pavé du tableau de pavé : 
            créer un nouveau pavé vide
            Si le pavé lut est dans la zone du pavé du tableau de pavé alors faire : 
                récupérer la plus grande valeur des minimums en x, y et z de ces deux pavés et associer ces valeurs au nouveau pavé
                récupérer la plus petite valeur des maximums en x, y et z de ces deux pavés et associer ces valeurs au nouveau pavé
                associer la valeur de ce nouveau pavé à l'inverse de la valeur du pavé du tableau de pavé 
            Sinon faire : 
                créer un pavé impossible pour avoir un pavé dit 'null'    
        Si le nouveau pavé n'est pas le pavé impossible alors faire : 
            ajouter au tableau de pavé le nouveau pavé 
    Si le mot est 'on' alors faire : 
        ajouter au tableau de pavé le pavé lut 
Compter le nombre de valeurs dans les pavés qui sont à vrai : 
    Initialiser le compteur à 0
    Pour chaque pavé dans la tableau de pavé faire : 
        Si la valeur du pavé est à vrai alors faire : 
            ajouter au compteur le volume du pavé (longuer * largeur * hauteur)
        Sinon faire : 
            enlever au compteur le volume du pavé (longuer * largeur * hauteur)
Retourner ce nombre
```
