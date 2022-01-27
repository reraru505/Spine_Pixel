package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Pallete struct {
	box1 rl.Rectangle
	box2 rl.Rectangle
	box3 rl.Rectangle

	col1 rl.Color
	col2 rl.Color
	col3 rl.Color
}

func pallateinit() Pallete {

	var pallete Pallete
	var redbox rl.Rectangle = rl.Rectangle{
		X:      40,
		Y:      40,
		Width:  40,
		Height: 40,
	}

	var greenbox rl.Rectangle = rl.Rectangle{
		X:      40,
		Y:      120,
		Width:  40,
		Height: 40,
	}

	var bluebox rl.Rectangle = rl.Rectangle{
		X:      40,
		Y:      200,
		Width:  40,
		Height: 40,
	}

	pallete.box1 = redbox
	pallete.box2 = greenbox
	pallete.box3 = bluebox

	pallete.col1 = rl.Red
	pallete.col2 = rl.Green
	pallete.col3 = rl.Blue

	return pallete
}

func drawpallete(pal *Pallete) {
	rl.DrawRectangleRec(pal.box1, pal.col1)
	rl.DrawRectangleRec(pal.box2, pal.col2)
	rl.DrawRectangleRec(pal.box3, pal.col3)
}

// func colorset(pal Pallete, mousecol *rl.Color) {

// 	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pal.box1) {
// 		*mousecol = rl.Red //&pal.col1
// 	}

// 	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pal.box2) {
// 		*mousecol = rl.Green //&pal.col2
// 	}

// 	if rl.CheckCollisionPointRec(rl.GetMousePosition(), pal.box3) {
// 		*mousecol = rl.Blue //&pal.col3
// 	}
// }

func brush(arr *[16][16]rl.Color, col rl.Color) {

	mousepos := rl.GetMousePosition()

	// var *somebuffer [16][16]rl.Color

	if mousepos.X < 1245 && mousepos.X > 605 && mousepos.Y > 64 && mousepos.Y < 704 {

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {

			// &somebuffer = arr

			arr[indexfinder(mousepos).xpos][indexfinder(mousepos).ypos] = col
		}
		if rl.IsMouseButtonDown(rl.MouseRightButton) {

			arr[indexfinder(mousepos).xpos][indexfinder(mousepos).ypos] = rl.White
		}
		// if rl.IsMouseButtonDown(rl.MouseRightButton) {

		// 	arr = &somebuffer
		// }

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
