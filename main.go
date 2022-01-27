package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {

	const width int32 = 1280
	const height int32 = 720

	rl.InitWindow(width, height, "test window")

	//rl.SetWindowState(rl.FlagWindowResizable)

	rl.SetTargetFPS(60)

	bux := boxarraypopulator()
	cux := colorarraypopulator()
	pal := pallateinit()
	var mousecol rl.Color

	// The main game loop of the code resides here

	for !rl.WindowShouldClose() {

		var point rl.Vector2 = rl.GetMousePosition()

		//mousecol = rl.Green
		if rl.CheckCollisionPointRec(point, pal.box1) && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousecol = rl.Red //&pal.col1
		}

		if rl.CheckCollisionPointRec(point, pal.box2) && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousecol = rl.Green //&pal.col2
		}

		if rl.CheckCollisionPointRec(point, pal.box3) && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			mousecol = rl.Blue //&pal.col3
		}

		// colorset(pal, &mousecol)

		brush(&cux, mousecol)

		rl.BeginDrawing()

		// Drawing starts here

		rl.ClearBackground(rl.Black)

		// bottomlayer()

		midlayer()

		drawpallete(&pal)

		recarrdraw(bux, cux)

		toplayyer()

		// Drawing ends here

		rl.EndDrawing()
	}

	rl.CloseWindow()

}
