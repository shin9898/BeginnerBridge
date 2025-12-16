package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"horseracing/game"
)

func main() {
	// フォントを初期化
	if err := game.InitFont(); err != nil {
		log.Printf("フォント初期化エラー: %v", err)
	}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Derby Dream - Horse Racing Tycoon")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
