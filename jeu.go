package main

import "fmt"

type Jeu struct {
	window Fenetre
}

func (j Jeu) Initialiser() {
	fmt.Println("Initialisation")

	window := Fenetre{800, 600, "Zeelda !", nil}
	window.Creer()
}

func (j Jeu) Run() {
	fmt.Println("En cours...")

	for true {

	}
}

func (j Jeu) Eteindre() {
	fmt.Println("Fermeture du jeu")
}
