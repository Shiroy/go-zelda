package main

import "fmt"

type Jeu struct {
	window Fenetre
}

func CreateJeu() Jeu {
	return Jeu{}
}

func (j Jeu) Initialiser() {
	fmt.Println("Initialisation")

	j.window = MakeFenetre(800, 600, "Zelda !")
	j.window.Creer()
}

func (j Jeu) Run() {
	fmt.Println("En cours...")

	for j.window.Closed() == false {
		j.window.ProcessEvents()
	}
}

func (j Jeu) Eteindre() {
	fmt.Println("Fermeture du jeu")

	j.window.w.Destroy()
}
