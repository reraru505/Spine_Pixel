package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type indices struct {
	xpos int32
	ypos int32
}

func boxarraypopulator() [16][16]rl.Rectangle {

	var boxarr [16][16]rl.Rectangle

	var box rl.Rectangle = rl.Rectangle{
		X:      605,
		Y:      64,
		Width:  40,
		Height: 40,
	}

	i1 := 0
	i2 := 0

	for i1 != 15 {
		if i2 < 16 {
			boxarr[i2][i1] = box

			box.X = box.X + 40

			i2++
		}
		if i2 > 15 {
			i2 = 0
			i1++
			box.X = 605
			box.Y = box.Y + 40
		}
	}

	return boxarr

}

func colorarraypopulator() [16][16]rl.Color {

	// var color rl.Color = rl.Color{
	// 	R: 255,
	// 	G: 255,
	// 	B: 255,
	// }

	color := rl.White

	var colorarr [16][16]rl.Color

	i1 := 0
	i2 := 0

	for i1 != 15 {

		if i2 < 16 {
			colorarr[i2][i1] = color
			i2++
		}
		if i2 > 15 {
			i2 = 0
			i1++
		}
	}

	return colorarr

}

func recarrdraw(arr [16][16]rl.Rectangle, crr [16][16]rl.Color) {

	i1 := 0
	i2 := 0

	for i2 != 15 {

		if i1 > 15 {
			i1 = 0
			i2++
		}

		if i1 < 16 {

			rl.DrawRectangleRec(arr[i1][i2], crr[i1][i2])

			i1++
		}

	}
}

func brush(arr *[16][16]rl.Color) {

	mousepos := rl.GetMousePosition()

	if mousepos.X < 1245 && mousepos.X > 605 && mousepos.Y > 64 && mousepos.Y < 704 {

		arr[indexfinder(mousepos).xpos][indexfinder(mousepos).ypos] = rl.Red
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

func main() {

	const width int32 = 1280
	const height int32 = 720

	rl.InitWindow(width, height, "test window")

	rl.SetWindowState(rl.FlagWindowResizable)

	rl.SetTargetFPS(60)

	bux := boxarraypopulator()
	cux := colorarraypopulator()

	for !rl.WindowShouldClose() {

		squarepos := rl.GetMousePosition()

		var posx int32 = int32(squarepos.X)
		var posy int32 = int32(squarepos.Y)

		brush(&cux)

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(605, 64, 640, 640, rl.Red)

		recarrdraw(bux, cux)

		rl.DrawCircle(posx, posy, 20, rl.Red)

		rl.DrawRectangle(64, 64, 64, 64, rl.Red)
		rl.DrawRectangle(64, 192, 64, 64, rl.Green)
		rl.DrawRectangle(64, 320, 64, 64, rl.Blue)

		rl.EndDrawing()
	}

	rl.CloseWindow()

}
