package particles

import (
	"project-particles/config"
)

func (p *Particle) PositionUpdate() {
	// Méthode mettant à jour la position des particules en fonction de leur vitesse et de la gravitée.
	// Cette gravitée peut être négative.
	if p.Death {

	}
	if config.General.Gravity {
		p.VitesseY += config.General.GravityCoefficient

		p.PositionX += p.VitesseX
		p.PositionY += p.VitesseY

	} else {
		p.PositionX += p.VitesseX
		p.PositionY += p.VitesseY
	}
}

func (p *Particle) IsInScreen() {
	// Tue la particule si elle est en dehors de l'écran + la marge donnée
	if !p.Death && p.PositionX <= -config.General.MargeCoefficient ||
		!p.Death && p.PositionY <= -config.General.MargeCoefficient ||
		!p.Death && p.PositionX >= float64(config.General.WindowSizeX)+config.General.MargeCoefficient ||
		!p.Death && p.PositionY >= float64(config.General.WindowSizeY)+config.General.MargeCoefficient {
		p.Kill()
	}
}

func (p *Particle) LifeCount() {
	// Calcul la vie d'une particule et la tue si sa vie est égale 0 et qu'elle n'est pas déjà morte
	if p.Life == 0 && !p.Death {
		p.Kill()
	} else if p.Life >= 1 && !p.Death {
		p.Opacity -= 1 / float64(p.OriginalLife)
		p.ColorBlue -= 1 / float64(p.OriginalLife)
		p.ColorGreen -= 1 / float64(p.OriginalLife)
		p.Life -= 1
	}
}

func (p *Particle) Explose(s *System, i int) {
	// La méthode Explose() vérifie si une particule est "explosive" ou non et crée en conséquences
	// le nombre de particules définit dans le fichier config.json
	if p.IsExplosive {
		for j := 0; j < config.General.NbExplose; j++ {
			s.add(false, i)
		}
		if !config.General.Trainee {
			p.Kill()
		}
	}
}

func (p *Particle) SquareSpawn() {
	// Fait spawn les particules sur les bords d'un carré de demi-diagonale RayonApparition.
	// On réutilise la variable IsExplosible de chaque particule pour savoir si elle est spawnable ou non.
	// Par défaut une particule est spawnable (IsExplosive = true) et donc dès que sa position permet son apparition,
	// elle n'est plus spawnable (IsExplosive = false). Cela permet d'éviter de la faire réapparaitre en boucle quand elle replis les conditions d'apparitions.

	if p.IsExplosive && (p.PositionX < config.General.SpawnX-config.General.RayonApparition ||
		p.PositionX > config.General.SpawnX+config.General.RayonApparition ||
		p.PositionY < config.General.SpawnY-config.General.RayonApparition ||
		p.PositionY > config.General.SpawnY+config.General.RayonApparition) {

		p.IsExplosive = false
		p.Opacity = 1
		if config.General.ActiveLife {
			p.Life = p.OriginalLife
		}
	}
}

func (P *Particle) Kill() {
	// Tue une particule
	P.Death = true
}
func (P *Particle) Revive() {
	// Récussite une particule
	P.Death = false
}
