# Partie 5

Pour compiler le projet, il vous faudra utiliser la commande `go build` depuis le dossier racine du projet pour obtenir un éxecutable adapté à votre système d'exploitation.

## Exemple de configuration

Pour voir ce que donnent les configurations suivantes, il suffit de modifier le chemin de la ligne de code suivante "config.Get("config.json"), du fichier main.go du projet par le nom des fichiers des fichiers json.

1. La configuration du fichier config.json montre le système de particules avec un spawn en "explosion".
2. La configuration du fichier config2.json montre le système de particules avec un spawnrate conséquent afin de tester les "limites" du programme.
   - La taille de la fenetre de cette configuration à été volontairement modifiée pour afficher plus de particules. 
3. La configuration du fichier config3.json est le fichier utilisant l'extension 5.10. Grace à celle-ci il est facile et rapide de tester une infinité de configurations différentes.
   - Par défaut cette configuration génère des particules qui explosent avec une trainée en fonction de la position de la souris. Le mode debug est activé.


## Variables du fichier config.json en fonction des extensions

> ### Extension 5.1 - Gravitée
> - Activable depuis la variable gravity (bool) du fichier config.json .
> - Force de gravitée modifiable avec la variable GravityCoefficient
>
> Les extensions fonctionnant avec celle-ci :
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (peu importe le type de spawn)
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.2 - Extérieur de l'écran
> - Activable depuis la variable Marge (bool) du fichier config.json .
> - Marge modifiable avec la variable MargeCoefficient
>
> Les extensions fonctionnant avec celle-ci :
> - 5.1 Gravitée
> - 5.3 Durée de vie
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (peu importe le type de spawn)
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.3 - Durée de vie
> - Activable depuis la variable ActiveLife du ficheir config.json .
> - Les particules ayant leur vie maximale sont rouges et s'estompent petit à petit jusqu'a ce qu'elles soient mortes.
> - Variables :
>   - RandomLife (Vie d'une paritcule aléatoire entre 0 et Life)
>   - Life (Vie d'une particule)
>
> Les extensions fonctionnant avec celle-ci :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur 
>   - Peu importe le type de spawn
>   - Est automatiquement activée avec le spawn "explosion" (avec et sans trainée)
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.5 - Optimisation de la mémoire
> Fonctionne automatiquement avec toutes les extensions

> ### Extension 5.7 - Forme du générateur
> Activable en changeant la variable SpawnType (string) du fichier config.json .
>
> Les valeurs prises en compte par cette variable sont les suivantes :
> - "none" : spawn normal, les variables à modifier sont :
>   - InitNumParticles
>   - RandomSpawn
>   - SpawnX, SpawnY
>   - SpawnRate
>   - Vitesse
>
> - "explosion" : spawn en formant une explosion, les variables à modifier sont :
>   - Celles possibles avec "none"
>   - Trainee (transforme l'explosion en trainée de particules depuis la particule explosive)
>   - RandomExplosedTime (Si la durée de vie des particules formant l'explosion est aléatoire)
>   - ExplosedTime (Durée de vie des particules formant l'explosion)
>   - NbExplose (Nombre de particules crées lors d'une explosion)
>   - VitesseExplose (Vitesse des particules formant l'explosion)
>
> - "carre" : spawn en autour d'un carré, les variables modifiables sont :
>   - Celles possibles avec "none"
>   - RayonApparition (distance à partir duquel les particules apparaissent)
>   - A savoir : L'activation de ce type de spawn entraine nécéssairement la désactivation du spawn random des particules
>
> Les extensions fonctionnant avec celle-ci sont les suivantes :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.5 Optimisation de la mémoire
> - 5.9 (sauf explosion  & MoveGenerator)
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.9 - Déplacement du générateur
>
> Activable depuis la variable MoveGenerator (bool) du fichier config.json .
>
> L'activation de cette extension entraine nécessairement la désactivation du spawn random des particules.
>
> Les différentes options de cette extension sont :
>
> - Les touches directionnelles pour changer le lieu de spawn des particules, les variables à modifier sont :
>   - MoveGenerator
> - Le curseur pour changer la position de spawn des particules, les variables à modifier sont :
>   - MouseMouvement (sous conditions que MoveGenerator soit activée)
> - La touche R tue toutes les particules
> - La touche L active et désactive la vie des particules
> - La vitesse des particules est modifiable avec la molette de la souris, les variables à modifier sont :
>   - WheelSpeed
>
> Les extensions fonctionnant avec celle-ci sont les suivantes :
>
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (sauf explosion)
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.10 - Modification dynamique du système de particules
> Extension activable depuis la variable Configurable du fichier config.json
> Les différentes options de cette extension :
> - La touche G du clavier ouvre ou ferme le menu visible en haut à gauche de la fenetre
> - les touches + et - du clavier numérique permettent de naviguer dans le menu
> - La molette de la souris permet le changement des valeurs. Vers le haut augmente la valeur, vers le bas diminue la valeur.
>
> Les extensions fonctionnant avec celle-ci sont :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur
> - 5.9 Déplacement du générateur

## Extensions réalisées & fonctionnement

- [x] 5.1 - Gravité
  - Ajouter dans le Json un booleen pour activer la gravitée
  - Ajouter une valeur à la vitesse Y de chaque particule
  - [x] Tests
  
- [x] 5.2 - Extérieur de l’écran
  - Vérifier si la position des particules est toujours situés dans l'écran (avec une marge)
  - Marge modifiable dans le fichier json
  - Lorsque la particule dépasse ces marges, la considérer comme morte (variable bool dans la structure d'une particule)
  - Ne plus l'afficher (dans le fichier draw.go)
  - [x] Tests

- [x] 5.3 - Durée de vie
  - durée de vie modifiable dans le fichier json
  - 1 compteur d'appel par particule (variable dans la structure d'une particule à ajouter dans la fonctions CreateParticle)
  - Une particule qui n'a plus de durée de vie est considérée comme morte
  - L'opacitée d'une particule baisse automatiquement et en conséquences en fonction de sa durée de vie
  - [x] Tests

- [x] 5.5 - Optimisation de la mémoire
  - La méthode add() prend en compte :
    - Si il y a des particules mortes, on les réutilises
    - Sinon on en ajoute au tableau
  - Variable LastInLife dans laquelle est stocké l'index de la dernière particule vivante
  - [x] Tests
  
- [x] 5.7 - Forme du générateur
  - [x] Carré
    - N'afficher les particule qu'après une marge définie
    - Nécéssite que les particules apparaissent au point de spawn défini (Forcé)
  - [x] Explosion
    - Définir si une particule est "explosive" ou non
    - D'autres particules non explosives serront crées à partir de la position de la particule explosive
  - [x] Trainée
    - Fonctionne avec le spawn en explosion
    - A la place de tuer la particule générant l'explosion, on la laisse en vie et continue selon l'algorithme
  - [x] Tests
  
- [x] 5.9 - Déplacement du générateur
  - Vérifier la pression des touches du clavier ou la position de la souris et changer les champs de spawnX et spawnY du fichier json
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.

- [x] 5.10 - Modification dynamique du système de particules
  - Lorsque la touche G est pressée, si le menu n'est pas affiché, il est affiché sinon il est caché (et ne permet plus la modification de valeurs)
  - La navigation du menu se fait avec les touches + et -. Retour au premier paramètre du menu si on est au dernier, et au dernier si on est au premier.
  - La modification des valeurs se fait par la molette, un paramètre permet de savoir si une valeur peut être négative et se bloque à 0 si ce n'est pas le cas.
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.


## Tests


Résultat du `go test -cover` :
```
PASS
coverage: 98.9% of statements
ok      project-particles/particles     0.251s
```