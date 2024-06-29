package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	// Initialisation du seed du random et déclaration des variables
	rand.Seed(time.Now().UnixNano())
	s := System{Content: []Particle{}, Generate: 0, LastInLife: -1}

	if config.General.SpawnType == "carre" || config.General.MoveGenerator {
		config.General.RandomSpawn = false
	}

	// Boucle créant un tableau de taille définie par InitNumParticles dans le fichier config.json
	for i := 0; i < config.General.InitNumParticles; i++ {
		s.add(true, 0)
	}

	return s
}
