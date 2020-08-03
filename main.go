package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/sizedru/otus_game_mk/game"
)

func main() {
	ebiten.SetWindowSize(600, 600)

	ebiten.RunGame(game.New())
}
