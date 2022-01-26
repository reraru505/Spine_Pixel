package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func brush(arr *[16][16]rl.Color) {

	mousepos := rl.GetMousePosition()

	if mousepos.X < 1245 && mousepos.X > 605 && mousepos.Y > 64 && mousepos.Y < 704 {

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {

			arr[indexfinder(mousepos).xpos][indexfinder(mousepos).ypos] = rl.Red
		}
		if rl.IsMouseButtonDown(rl.MouseRightButton) {

			arr[indexfinder(mousepos).xpos][indexfinder(mousepos).ypos] = rl.White
		}

	}

}

func indexfinder(mouse rl.Vector2) indices {

	buffindex0 := mouse.X - 605
	buffindex1 := buffindex0 / 40

	var indexA int32 = int32(buffindex1)

	buffindex2 := mouse.Y - 64
	buffindex3 := buffindex2 / 40

	var indexB int32 = int32(buffindex3)

	var index indices = indices{
		xpos: indexA,
		ypos: indexB,
	}

	return index

}
