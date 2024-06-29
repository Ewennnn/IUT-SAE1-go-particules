package particles

// System définit un système de particules.
// Pour le moment il ne contient qu'un tableau de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content    []Particle
	Generate   float64
	LastInLife int
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
type Particle struct {
	PositionX, PositionY            float64
	VitesseX, VitesseY              float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	Death                           bool // Par défaut une particule n'est pas morte
	OriginalLife                    int  // Sert à calculer l'opacitée de la particule en fonction de sa vie
	Life                            int  // Variable permettant de calculer la vie d'une particule
	IsExplosive                     bool // Variable permettant de savoir si une particule est "explosive" ou non
}
