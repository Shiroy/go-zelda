package main

import "fmt"

func main() {
	fmt.Println("Hello world !")

	var jeu = CreateJeu()

	jeu.Initialiser()
	defer jeu.Eteindre()

	jeu.Run()
}
