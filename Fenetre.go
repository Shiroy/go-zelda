package main

import "github.com/veandco/go-sdl2/sdl"

type Fenetre struct {
	largeur int
	hauteur int
	titre   string
	w       *sdl.Window

	closed bool
}

func MakeFenetre(largeur, hauteur int, titre string) Fenetre {
	return Fenetre{largeur, hauteur, titre, nil, false}
}

func (window *Fenetre) Creer() {
	window.w = sdl.CreateWindow(window.titre, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, window.largeur, window.hauteur, sdl.WINDOW_SHOWN)
	//window.w.
}

func (window *Fenetre) ProcessEvents() {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			window.closed = true
			break
		}
	}
}

func (window *Fenetre) Closed() bool {
	return window.closed
}
