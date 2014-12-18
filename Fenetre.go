package main

import "github.com/veandco/go-sdl2/sdl"

type Fenetre struct {
	largeur int
	hauteur int
	titre   string
	w       *sdl.Window
}

func (window Fenetre) Creer() {
	window.w = sdl.CreateWindow(window.titre, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, window.largeur, window.hauteur, sdl.WINDOW_SHOWN)
}
