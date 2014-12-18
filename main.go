package main

import "fmt"

func main() {
	fmt.Println("Hello world !")

	jeu := Jeu{}
	jeu.Initialiser()
	jeu.Run()
	jeu.Eteindre()
}
