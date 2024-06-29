package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           float64
	SpawnRate                float64
	Vitesse                  float64 // Vitesse maximale d'une particule
	Gravity                  bool    // Activer l'extension gravitée
	GravityCoefficient       float64 // Coefficient de graitée (plus il est élevé, plus les particules sont attirés vers le bas ou vers le haut rapidement)
	Marge                    bool    // Activer l'extension extérieur de l'écran
	MargeCoefficient         float64 // Marge après laquelle les particules sont tuées si elle est dépassée
	ActiveLife               bool    // Activer l'extension Durée de vie
	RandomLife               bool    // Les particules peuvent avoir une vie aléatoire entre 0 et Life
	Life                     int     // Vie d'une particule
	SpawnType                string  // Activer l'extension Forme du Générateur en fonction de la valeur rentrée (valeurs possibles disponibles dans le fichier readme.md)
	Trainee                  bool    // Activer une trainée, implique que spawtype soit explosif
	RandomExplosedTime       bool    // Les particules générées par l'explosion d'une autre ont une durée de vie aléatoire entre 0 et ExplosedTime
	ExplosedTime             int     // Durée de vie d'une particule générée à la suite d'une explosion
	NbExplose                int     // Nombre de particules générées à la suite de l'explosion d'une autre
	VitesseExplose           float64 // Vitesse des particules générées à la suite de l'explosion d'une autre
	RayonApparition          float64 // Distance après laquelle une particule apparait
	MoveGenerator            bool    // Permet de déplacer le générateur avec les flèches
	MouseMouvement           bool    // Permet de déplacer le point de spawn des particules avec la souris
	WheelSpeed               bool    // Permet de modifier la vitesse des particules avec la molette de la souris
	Configurable             bool    // Permet de modifier les valeurs du config.json directement dans le "jeu"
}

var General Config

// Structure et variable permettant la communication entre le package main et le package particles
// pour afficher le menu et leur valeur.
type ConfPrint struct {
	Index       int
	Extense     string
	Type        string
	ValueBool   bool
	ValueInt    int
	ValueFloat  float64
	ValueString string
}

var Show bool         // booléen permettant de savoir si le menu de configuration est affiché ou non
var ToPrint ConfPrint // Variable de structure permettant d'échanger les informations
