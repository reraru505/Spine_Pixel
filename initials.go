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

	for i1 != 16 {
		if i2 < 16 {
			boxarr[i2][i1] = box

			box.X = box.X + 40

			i2++
		}
		if i2 == 16 {
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

	// for i2 != 15 {

	// 	if i1 > 15 {
	// 		i1 = 0
	// 		i2++
	// 	}

	// 	if i1 < 16 {

	// 		rl.DrawRectangleRec(arr[i1][i2], crr[i1][i2])

	// 		i1++
	// 	}

	// }

	for i1 != 15 {

		if i2 == 16 {

			i2 = 0

			i1++
		}

		rl.DrawRectangleRec(arr[i2][i1], crr[i2][i1])

		i2++
	}
}
