package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {

	g.system.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyL) && !config.General.RandomSpawn && config.General.SpawnType != "explosion" {
		if config.General.ActiveLife {
			config.General.ActiveLife = false
		} else {
			config.General.ActiveLife = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for i := 0; i < len(g.system.Content); i++ {
			g.system.Content[i].Kill()
		}
	}
	if config.General.Configurable {
		ShowConfigMenu()
		ShowConfigByIndex()
	}

	if config.General.MoveGenerator {
		Mouvement()
	}

	return nil
}

func WheelSpeed() {
	_, dy := ebiten.Wheel()
	if config.General.Vitesse >= 0 || dy > 0 {
		config.General.Vitesse += dy * 0.1
	}
}

func Mouvement() {
	if config.General.MouseMouvement {
		// Déplacement à la souris
		mx, my := ebiten.CursorPosition()

		config.General.SpawnX, config.General.SpawnY = float64(mx), float64(my)

	} else {
		// Déplacement au clavier
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			config.General.SpawnY--
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			config.General.SpawnY++
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			config.General.SpawnX--
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			config.General.SpawnX++
		}
	}
}

// Variables globales permettant la gestion, l'affichage et la modification dans le menu
var k int
var h int

// Variables "poubelles" permettant de remplir la structure
var emptyb *bool
var emptyf *float64
var emptyi *int
var emptys *string

func ShowConfigMenu() {
	// Afficher ou cache le menu lorsque la touche G est active
	// Gère le menu complet, lorsque la touche + ou - (numpad) est active, le menu affiche la configuration suivante ou précédente du tableau Configurable.
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		if config.Show {
			config.Show = false

			config.ToPrint.Index = -1
			config.ToPrint.Extense = ""
			config.ToPrint.Type = ""
		} else {
			config.Show = true
			config.ToPrint.Index = k + 1
			config.ToPrint.Extense = Configurable[k].Extension
			config.ToPrint.Type = Configurable[k].Type
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) && config.Show {
		k++
		if k > len(Configurable)-1 {
			k = 0
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) && config.Show {
		k--
		if k < 0 {
			k = len(Configurable) - 1
		}
	}
	config.ToPrint.Index = k + 1
	config.ToPrint.Extense = Configurable[k].Extension
	config.ToPrint.Type = Configurable[k].Type
}

func ShowConfigByIndex() {
	// Cette fonction vérifie le type de variable à modifier et change les valeurs au roulement de la molette de la souris.
	// Les valeurs booléennes sont simplement échangés.
	// Le valeurs int et float sont augmentés et diminuer selon leur incrément
	// La variable string est modifiée selon son propre tableau

	_, dy := ebiten.Wheel()

	if Configurable[k].Type == "bool" {
		if (dy > 0 || dy < 0) && *Configurable[k].Parametrebool {
			*Configurable[k].Parametrebool = false
		} else if (dy > 0 || dy < 0) && !*Configurable[k].Parametrebool {
			*Configurable[k].Parametrebool = true
		}
		config.ToPrint.ValueBool = *Configurable[k].Parametrebool
	}
	if Configurable[k].Type == "int" {
		if dy > 0 {
			*Configurable[k].Parametreint++
		} else if dy < 0 {
			if !Configurable[k].CanBeNegative && *Configurable[k].Parametreint <= 0 {
				*Configurable[k].Parametreint = 0
			} else {
				*Configurable[k].Parametreint--
			}
		}
		config.ToPrint.ValueInt = *Configurable[k].Parametreint
	}
	if Configurable[k].Type == "float64" {
		if dy > 0 {
			*Configurable[k].Parametrefloat += Configurable[k].Increment
		} else if dy < 0 {
			if !Configurable[k].CanBeNegative && *Configurable[k].Parametrefloat <= 0 {
				*Configurable[k].Parametrefloat = 0
			} else {
				*Configurable[k].Parametrefloat -= Configurable[k].Increment
			}
		}
		config.ToPrint.ValueFloat = *Configurable[k].Parametrefloat
	}
	if Configurable[k].Type == "string" {
		if dy > 0 && h <= len(SpawnTabList)-1 {
			h++
			if h > len(SpawnTabList)-1 {
				h = 0
			}
			*Configurable[k].Parametrestr = SpawnTabList[h]
		} else if dy < 0 && h >= 0 {
			h--
			if h < 0 {
				h = len(SpawnTabList) - 1
			}
			*Configurable[k].Parametrestr = SpawnTabList[h]
		}
		config.ToPrint.ValueString = *Configurable[k].Parametrestr
	}
}

// Structure définissant chaque variable de la configuration
// Les types pointeurs sont utilisés pour pointer vers la variable config.General.[...] sans quoi on ne peux pas modifier directement la variable du fichier config.json .
type ModifConfig struct {
	Extension      string
	Type           string
	Parametrebool  *bool
	Parametreint   *int
	Parametrefloat *float64
	Increment      float64
	CanBeNegative  bool
	Parametrestr   *string
}

// Tableau des types de spawn
var SpawnTabList []string = []string{
	"none",
	"explosion",
	"carre",
}

// Configurable Tableau contenant toutes les variables du fichier config.json intéréssantes à modifier.
// Dans chaque index est repertorié :
// - Le nom de la variable
// - le type de variable (permettant une détection rapide et simple par une variable string)
// - L'adresse de la variable à modifier (config.Genral.[...])
// - L'incrémentation (valeur à ajouter ou à soustraire pour les valeurs int et float64)
// - Si la valeur peut être négative (int et float64)
var Configurable = []ModifConfig{
	{"RandomSpawn", "bool", &config.General.RandomSpawn, emptyi, emptyf, 0, false, emptys},
	{"SpawnRate", "float64", emptyb, emptyi, &config.General.SpawnRate, 0.01666667, false, emptys},
	{"Vitesse", "float64", emptyb, emptyi, &config.General.Vitesse, 0.1, false, emptys},

	{"Gravitée", "bool", &config.General.Gravity, emptyi, emptyf, 0, false, emptys},
	{"Coefficient de gravitée", "float64", emptyb, emptyi, &config.General.GravityCoefficient, 0.05, true, emptys},

	{"Marge", "bool", &config.General.Marge, emptyi, emptyf, 0, false, emptys},
	{"Taille de la marge", "float64", emptyb, emptyi, &config.General.MargeCoefficient, 1, true, emptys},

	{"Activer la vie", "bool", &config.General.ActiveLife, emptyi, emptyf, 0, false, emptys},
	{"Vie Aléatoire", "bool", &config.General.RandomLife, emptyi, emptyf, 0, false, emptys},
	{"Vie", "int", emptyb, &config.General.Life, emptyf, 1, false, emptys},

	{"Type de spawn", "string", emptyb, emptyi, emptyf, 0, false, &config.General.SpawnType},
	{"Trainée", "bool", &config.General.Trainee, emptyi, emptyf, 0, false, emptys},
	{"Durée de vie des particules d'explosion aléatoire", "bool", &config.General.RandomExplosedTime, emptyi, emptyf, 0, false, emptys},
	{"Durée de vie des particules d'explosion", "int", emptyb, &config.General.ExplosedTime, emptyf, 1, false, emptys},
	{"Nombre de particules à l'explosion", "int", emptyb, &config.General.NbExplose, emptyf, 1, false, emptys},
	{"Vitesse des particules d'explosion", "float64", emptyb, emptyi, &config.General.VitesseExplose, 0.1, false, emptys},
	{"Rayon d'apparition", "float64", emptyb, emptyi, &config.General.RayonApparition, 1, false, emptys},

	{"Déplacement du générateur", "bool", &config.General.MoveGenerator, emptyi, emptyf, 0, false, emptys},
	{"Déplacement du générateur avec la souris", "bool", &config.General.MouseMouvement, emptyi, emptyf, 0, false, emptys},

	{"Debug Mode", "bool", &config.General.Debug, emptyi, emptyf, 0, false, emptys},
}
