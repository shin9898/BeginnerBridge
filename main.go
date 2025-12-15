package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"horseracing/game"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("競走馬育成シミュレーション - Horse Racing Tycoon")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
