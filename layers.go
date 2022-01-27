package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

// The elements that dont require initialization outside
//gameloop are placed acording to their drawing priority

func toplayyer() {
	squarepos := rl.GetMousePosition()

	var posx int32 = int32(squarepos.X)
	var posy int32 = int32(squarepos.Y)
	rl.DrawCircle(posx, posy, 20, rl.Red)
}

func midlayer() {

}

func bottomlayer() {
	rl.DrawRectangle(600, 59, 650, 650, rl.Red)
}
