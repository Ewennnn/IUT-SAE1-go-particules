package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s *System) add(isexplo bool, i int) {

	// La méthode add permet d'ajouter une particule au tableau ou de réutiliser une particule existante.
	// Elle permet également de générer une particule explosive ou non en fonction des paramètres reçu.
	var PosX, PosY, vit, opac float64
	var life int

	if isexplo {
		PosX, PosY, vit, opac, life = ReturnConfigValues()
	} else {
		PosX = s.Content[i].PositionX
		PosY = s.Content[i].PositionY
		vit = config.General.VitesseExplose
		opac = 1
		if config.General.RandomExplosedTime {
			life = rand.Intn(config.General.ExplosedTime)
		} else {
			life = config.General.ExplosedTime
		}
	}

	if s.LastInLife < len(s.Content)-1 {
		s.Content[s.LastInLife+1] = CreateParticle(PosX, PosY, vit, opac, life, isexplo)
		s.LastInLife++
	} else {
		s.Content = append(s.Content, CreateParticle(PosX, PosY, vit, opac, life, isexplo))
		s.LastInLife++
	}
}

func ReturnConfigValues() (x, y, vit, opac float64, life int) {
	// Fonction permettant de retourner toutes les valeurs nécéssaires à la création d'une particule
	var PosX, PosY float64
	var Life int
	var vitesse float64 = config.General.Vitesse

	if config.General.RandomSpawn {
		// Génère de coordonnées aléatoires
		PosX = float64(rand.Intn(config.General.WindowSizeX))
		PosY = float64(rand.Intn(config.General.WindowSizeY))
	} else {
		// Positionne la particule aux coordonnées de spawn
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
	if config.General.ActiveLife {
		Life = config.General.Life
		if config.General.RandomLife {
			Life = int(RandomNumber(0, float64(config.General.Life)))
		}
	} else {
		Life = -1
	}
	if config.General.SpawnType == "carre" {
		opac = 0
	} else {
		opac = 1
	}
	return PosX, PosY, vitesse, opac, Life
}

func CreateParticle(PosX, PosY, vit, opac float64, life int, isexplo bool) Particle {
	// Fonction permettant de crée une particule.
	return Particle{
		PositionX: PosX,
		PositionY: PosY,
		VitesseX:  RandomNumber(-vit, vit), VitesseY: RandomNumber(-vit, vit),
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		ScaleX: 1, ScaleY: 1,
		Opacity:      opac,
		OriginalLife: life, Life: life,
		IsExplosive: isexplo,
	}
}

func RandomNumber(min, max float64) float64 {
	// Fonction permettant de retoruner un nombre aléatoire dans un interval donnée.
	return min + rand.Float64()*(max-min)
}
