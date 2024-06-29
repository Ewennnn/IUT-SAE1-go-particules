package particles

import (
	"math/rand"
	"project-particles/config"
	"testing"
	"time"
)

// Ce fichier à pour but de tester les fonctions crée pour réaliser la première partie du projet.
// Test de la fonction RandomNumer.
func TestRandomNumber(t *testing.T) {
	var v float64 = 100
	for i := 0; i < 1000; i++ {
		r := RandomNumber(-v, v)
		if r < -v || r > v {
			t.Fail()
		}
	}
}

// Tests de la fonction CreateParticle faisant appel à la fonction RandomNumber.
func TestCreateParticle0(t *testing.T) {
	P := CreateParticle(0, 0, 0, 1, 0, false)
	if P.PositionX != 0 || P.PositionY != 0 ||
		P.VitesseX != 0 || P.VitesseY != 0 {
		t.Fail()
	}
}
func TestCreateParticle1(t *testing.T) {
	P := CreateParticle(100, 100, 5, 1, 0, true)
	if P.PositionX != 100 || P.PositionY != 100 ||
		P.VitesseX < -5 || P.VitesseX > 5 ||
		P.VitesseY < -5 || P.VitesseY > 5 ||
		!P.IsExplosive {
		t.Fail()
	}
}
func TestCreateParticleSpawn(t *testing.T) {
	config.Get("../config.json")

	P := CreateParticle(config.General.SpawnX, config.General.SpawnY, 1, 1, 0, false)
	if P.PositionX != config.General.SpawnX || P.PositionY != config.General.SpawnY ||
		P.VitesseX < -1 || P.VitesseX > 1 ||
		P.VitesseY < -1 || P.VitesseY > 1 {
		t.Fail()
	}
}
func TestCreateParticleRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	config.Get("../config.json")

	for i := 0; i < 100; i++ {
		var PosX float64 = float64(rand.Intn(config.General.WindowSizeX))
		var PosY float64 = float64(rand.Intn(config.General.WindowSizeY))

		P := CreateParticle(PosX, PosY, 10, 1, 0, false)

		if P.PositionX < 0 || P.PositionX > float64(config.General.WindowSizeX) ||
			P.PositionY < 0 || P.PositionY > float64(config.General.WindowSizeY) ||
			P.VitesseX < -10 || P.VitesseX > 10 ||
			P.VitesseY < -10 || P.VitesseY > 10 {
			t.Fail()
		}
	}
}

// Tests de la fonction NewSystem faisant appel à la fonction CreateParticle.
func TestNewSystem0(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 0
	S := NewSystem()

	if len(S.Content) != config.General.InitNumParticles {
		t.Fail()
	}
}

func TestNewSystem100(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 100
	S := NewSystem()

	if len(S.Content) != config.General.InitNumParticles {
		t.Fail()
	}
}

// Tests de la méthode add faisant appel à la fonction CreateParticle.
// La fonction add ajoute une particule à la fois, on ne peux pas la tester pour les valeurs décimales.
func TestAdd1(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 0
	S := NewSystem()
	S.add(true, 0)
	if len(S.Content) != 1 {
		t.Fail()
	}
}
func TestAdd10(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 10
	S := NewSystem()
	for i := 0; i < 10; i++ {
		S.add(true, 0)
	}
	if len(S.Content) != 20 {
		t.Fail()
	}
}

// Tests de la méthode Update faisant appel à la mathode add.
func TestZero(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 0
	S := NewSystem()
	for i := 0; i < 100; i++ {
		S.Update()
	}
	if len(S.Content) != 0 {
		t.Fail()
	}
}
func TestUpdatePositif1(t *testing.T) {
	config.Get("../config.json")
	config.General.SpawnType = "none"
	config.General.Marge = false
	config.General.ActiveLife = false
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 3
	config.General.Vitesse = 0
	S := NewSystem()
	for i := 0; i < 10; i++ {
		S.Update()
	}
	if len(S.Content) != 30 {
		t.Error(len(S.Content))
	}
}
func TestUpdatePositif2(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 10
	config.General.SpawnRate = 15
	config.General.Vitesse = 0
	config.General.SpawnType = "none"
	config.General.Marge = false
	config.General.ActiveLife = false
	S := NewSystem()
	for i := 0; i < 60; i++ {
		S.Update()
	}
	if len(S.Content) != 910 {
		t.Error(len(S.Content))
	}
}

func TestUpdateZero5(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 0.5
	config.General.Vitesse = 0
	config.General.SpawnType = "none"
	config.General.Marge = false
	config.General.ActiveLife = false
	S := NewSystem()
	for i := 0; i < 60; i++ {
		S.Update()
	}
	if len(S.Content) != 30 {
		t.Error(len(S.Content))
	}
}
func TestUpdateZerozero16(t *testing.T) {
	// Le cas de 1 particule par seconde est spécial, son SpawnRate et un nombre décimal ayant une infinitée de chiffres après la virgule.
	// On aura donc qu'une valeur approchée qui se rapproche fortement de la réalitée.
	// Dans sa réelle utilisation, la méthode Update est appelée ENVIRON 60 fois par secondes, parfois plus, parfois moins.
	// On ne se rend alors pas compte des minimes différences théoriques.
	config.Get("../config.json")
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 0.01666667
	config.General.Vitesse = 0
	config.General.SpawnType = "none"
	S := NewSystem()
	for i := 0; i < 120; i++ {
		S.Update()
	}
	if len(S.Content) < 1 || len(S.Content) > 2 {
		t.Error(len(S.Content))
	}
}

func TestUpdatePosition(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 1
	config.General.SpawnRate = 0
	config.General.Vitesse = 1

	S := NewSystem()
	originalPositionX := S.Content[0].PositionX
	originalPositionY := S.Content[0].PositionY

	for i := 0; i < 10; i++ {
		S.Update()
	}
	if S.Content[0].PositionX == originalPositionX || S.Content[0].PositionY == originalPositionY {
		t.Fail()
	}
}

// Les tests suivant sont ceux permettant de réaliser les extensions proposées en partie 5 du projet

// Test des méthodes MakeDeath() et MakeAlive()
func TestAliveAndDeath(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles = 10
	config.General.SpawnRate = 0

	S := NewSystem()
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Death == true {
			t.Fail()
		}
		S.Content[i].Kill()
	}
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Death == false {
			t.Fail()
		}
		S.Content[i].Revive()
	}
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Death == true {
			t.Fail()
		}
	}
}

// Test de l'extension 5.1 (Gravitée)
// On crée une particule et on vérifie si elle est correctement déviée.
func TestGravity(t *testing.T) {
	config.Get("../config.json")

	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnType = "none"
	config.General.SpawnY = 0
	config.General.Vitesse = 1

	config.General.Gravity = true
	config.General.GravityCoefficient = 1

	config.General.Marge = false
	config.General.ActiveLife = false

	S := NewSystem()
	S.Content[0].VitesseX = 0
	S.Content[0].VitesseY = 0

	if len(S.Content) != 1 {
		t.Fail()
	} else {
		for i := 0; i < 10; i++ {
			S.Update()
		}
		if S.Content[0].PositionY != 55 {
			t.Fail()
		}
	}
}

// Test de l'extension 5.2 (Marge à l'extérieur de l'écran)
// On admet par des vérifications visuelles que si la particule est à l'état de "morte" et qu'elle est à l'extérieur des marges, alors on ne la voit pas et qu'elle n'est plus mise à jour.
func TestMarge(t *testing.T) {
	config.Get("../config.json")

	config.General.WindowSizeX = 100
	config.General.WindowSizeY = 100

	config.General.Gravity = false
	config.General.SpawnType = "none"
	config.General.Marge = true
	config.General.MargeCoefficient = 10

	config.General.RandomSpawn = false
	config.General.SpawnX = 50
	config.General.SpawnY = 50

	config.General.InitNumParticles = 10
	config.General.SpawnRate = 0
	config.General.Vitesse = 5
	config.General.ActiveLife = false

	S := NewSystem()
	for i := 0; i < len(S.Content); i++ {
		S.Content[i].VitesseX = 1
		S.Content[i].VitesseY = 1
	}

	for i := 0; i < 1000; i++ {
		S.Update()
	}
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Death == false ||
			S.Content[i].PositionX > -10 && S.Content[i].PositionX < 110 ||
			S.Content[i].PositionY > -10 && S.Content[i].PositionY < 110 {
			t.Error(S.Content[i].PositionX, S.Content[i].PositionY, S.Content[i].Death)
		}
	}

}

// Test de l'extension 5.3 (Durée de vié)
// On définit la durée de vie d'une particule, si au bout des 100 updates une d'entre-elles n'est pas morte, c'est qu'il y a une erreur.
// On test également avec les marges, dans certains cas (qu'on ne voit pas) une particule peut devenir morte en dépassant cette marge.
func TestLife(t *testing.T) {
	config.Get("../config.json")
	config.General.Marge = true
	config.General.Gravity = false
	config.General.SpawnType = "none"
	config.General.ActiveLife = true
	config.General.RandomLife = false
	config.General.Life = 69

	config.General.InitNumParticles = 10
	config.General.SpawnRate = 0

	S := NewSystem()

	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Life != config.General.Life {
			t.Fail()
		}
	}
	for i := 0; i < 100; i++ {
		S.Update()
	}
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].Life != 0 && S.Content[i].Death != true ||
			S.Content[i].Opacity > float64(config.General.Life)-1/float64(config.General.Life) {
			t.Fail()
		}
	}

	config.General.RandomLife = true

	W := NewSystem()

	for i := 0; i < len(W.Content); i++ {
		if W.Content[i].Life < 0 || W.Content[1].Life > config.General.Life {
			t.Fail()
		}
	}
	for i := 0; i < 100; i++ {
		W.Update()
	}
	for i := 0; i < len(S.Content); i++ {
		if W.Content[i].Life != 0 && W.Content[i].Death != true {
			t.Fail()
		}
	}
}

// Test de l'extension 5.7 (Forme du générateur)
// Ce test teste uniquement l'explosion et non la trainée ou le spawn carré.
func TestExplose(t *testing.T) {
	// appel uniquement de la méthode Explose()
	// Compter le nombre de particules crée via cette méthode
	// Vérifier les durées de vie des particules (random & fixe)
	// Vérifier les coordonnées de spawn des particules

	config.Get("../config.json")
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = true
	config.General.SpawnRate = 0
	config.General.ActiveLife = false
	config.General.Marge = false

	config.General.SpawnType = "explosion"
	config.General.Trainee = false
	config.General.RandomExplosedTime = false
	config.General.ExplosedTime = 100
	config.General.NbExplose = 45
	config.General.VitesseExplose = 1.2

	S := NewSystem()
	S.Content[0].Explose(&S, 0)

	if len(S.Content) != 46 {
		t.Fail()
	}
	for i := 1; i < len(S.Content); i++ {
		if S.Content[i].PositionX != S.Content[0].PositionX ||
			S.Content[i].PositionY != S.Content[0].PositionY ||
			S.Content[i].Life != config.General.ExplosedTime {
			t.Error(S.Content[i].Life)
		}
	}

	config.General.RandomExplosedTime = true

	W := NewSystem()
	W.Content[0].Explose(&W, 0)

	for i := 1; i < len(W.Content); i++ {
		if W.Content[i].Life < 0 ||
			W.Content[i].Life > config.General.ExplosedTime {
			t.Error(W.Content[i].Life)
		}
	}
}

// Test de l'extension 5.7 (Forme du générateur)
// Ce test teste uniquement la trainée et non l'explosion
func TestTrainee(t *testing.T) {
	// len(S.Content) = nbparticules + nbparticules(nbexplose*nbupdate)
	// Seulement si ExplosedTime > nbUpdate
	// Car si une particule est morte, la méthode Uptade() la réutilise
	config.Get("../config.json")
	config.General.InitNumParticles = 9
	config.General.RandomSpawn = true
	config.General.SpawnRate = 0
	config.General.ActiveLife = false
	config.General.Marge = false

	config.General.SpawnType = "explosion"
	config.General.Trainee = true
	config.General.RandomExplosedTime = false
	config.General.ExplosedTime = 200
	config.General.NbExplose = 40
	config.General.VitesseExplose = 1.2

	S := NewSystem()

	for i := 0; i < 180; i++ {
		S.Update()
	}
	if len(S.Content) != 64809 {
		t.Error(len(S.Content))
	}
}

// Test de l'extension 5.7 (Forme du générateur)
// Ce test teste uniquement le spawn carré et non l'explosion
func TestCarre(t *testing.T) {
	// Vérifier si les particules apparaissent après les marges de RayonApparition
	// Vérifier cette apparition avec la variable IsExplosive (variable réutilisée pour éviter de recrée une variable)
	// Vérifier si la vie de la particule une fois qu'elle a franchi cette marge à la même vie que lorsqu'elle a été crée
	config.Get("../config.json")
	config.General.InitNumParticles = 1000
	config.General.SpawnRate = 0
	config.General.ActiveLife = false
	config.General.Marge = false

	config.General.SpawnType = "carre"
	config.General.RayonApparition = 100

	S := NewSystem()

	if len(S.Content) != 1000 {
		t.Fail()
	}
	for i := 0; i < len(S.Content); i++ {
		S.Content[i].VitesseX, S.Content[i].VitesseY = 1, 1
		if !S.Content[i].IsExplosive {
			t.Fail()
		}
	}
	for i := 0; i < 101; i++ {
		S.Update()
	}
	for i := 0; i < len(S.Content); i++ {
		if S.Content[i].IsExplosive || S.Content[i].Life != S.Content[i].OriginalLife {
			t.Fail()
		}
	}
}
