package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.

var WherePrint int

func (g *game) Draw(screen *ebiten.Image) {

	// Donne la position Y du menu en fonction de si le mode debug est activé ou pas
	if config.General.Debug {
		WherePrint = 70
	} else {
		WherePrint = 10
	}

	for _, p := range g.system.Content {
		if p.Death {
			continue
		}
		options := ebiten.DrawImageOptions{}
		options.GeoM.Rotate(p.Rotation)
		options.GeoM.Scale(p.ScaleX, p.ScaleY)
		options.GeoM.Translate(p.PositionX, p.PositionY)
		options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
		screen.DrawImage(assets.ParticleImage, &options)
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombre de particules dans le tableau : ", len(g.system.Content)), 0, 20)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombre de particules affichées :       ", len(g.system.Content)-len(g.system.Content)+g.system.LastInLife+1), 0, 30)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombre de particules mortes :          ", len(g.system.Content)-1-g.system.LastInLife), 0, 40)
	}

	// Affiche la variable qui est sélectionnée dans le menu avec sa valeur (en direct)
	if config.General.Configurable && config.Show {
		ebitenutil.DebugPrintAt(screen, fmt.Sprint(config.ToPrint.Index, " - ", config.ToPrint.Extense, " (", config.ToPrint.Type, ")"), 0, WherePrint-10)
		if Configurable[k].Type == "bool" {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint(config.ToPrint.ValueBool), 0, WherePrint)
		}
		if Configurable[k].Type == "int" {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint(config.ToPrint.ValueInt), 0, WherePrint)
		}
		if Configurable[k].Type == "float64" {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint(float64(config.ToPrint.ValueFloat)), 0, WherePrint)
		}
		if Configurable[k].Type == "string" {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint(config.ToPrint.ValueString), 0, WherePrint)
		}
	}
}
