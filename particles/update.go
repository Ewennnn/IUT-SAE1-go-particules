package particles

import (
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {

	// Suite d'instructions correspondant à la partie 4
	// et sous condition d'activation dans le fichier json, les extension de la partie 5 expliquées dans le fichier readme.md
	for i := 0; i <= s.LastInLife; i++ {

		s.Content[i].PositionUpdate()

		if config.General.SpawnType == "explosion" {
			// L'extension 5.7, a besoin de l'extension 5.3 pour fonctionner, on appel donc automatiquement la méthode liée à cette extension
			s.Content[i].Explose(s, i)
			s.Content[i].LifeCount()
		}
		if config.General.SpawnType == "carre" {
			s.Content[i].SquareSpawn()
		}

		if config.General.Marge {
			s.Content[i].IsInScreen()
		}
		if config.General.ActiveLife && config.General.SpawnType != "explosion" {
			s.Content[i].LifeCount()
		}

		if s.Content[i].Death {
			s.Content[i], s.Content[s.LastInLife] = s.Content[s.LastInLife], s.Content[i]
			s.LastInLife--
			i--
		}
	}

	// Crée des particules en fonction de la valeur de SpawnRate.
	// La boucle crée le nombre de particule voulue par appel de la fonction Update et prend en compte les valeurs en dessous de 1.
	s.Generate += config.General.SpawnRate
	for s.Generate >= 1 {
		s.add(true, 0)
		s.Generate -= 1
	}
}
