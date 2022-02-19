package main

import (
	//	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)



type Initdata struct{

	pixrec rl.Rectangle
	pixarr [16][16]rl.Rectangle

	colorarray  [16][16]rl.Color
	colorbuffer [16][16]rl.Color
	undobuffer  [16][16]rl.Color

}



func init16(i *Initdata , side float32,xx int,yy int) {

	posx := float32(xx)
	posy := float32(yy)

	


	i.pixrec = rl.Rectangle{

		X : posx,
		Y : posy,
		Width : side,
		Height : side,
	}

	x := 0
	y := 0

	for x < 15 {

		
		if y == 16 {
			
			y = 0
			x++
			i.pixrec.X = posx
			i.pixrec.Y = i.pixrec.Y + side			
			}
		
		for y < 16 {
			

			i.pixarr[y][x] = i.pixrec
			
			y++

			i.pixrec.X = i.pixrec.X + side
		}
	}

}

func initcolorarray(i *Initdata) {

	x := 0 
	y := 0

	for x < 15 {

		if y == 16 {
			
			y = 0
			x++
			
			}
		
		for y < 16 {
			
			i.colorarray  [y][x] = rl.White
			i.colorbuffer [y][x] = rl.White
			i.undobuffer  [y][x] = rl.White

			y++

		}
	}
}


func draw16(i *Initdata){

	x := 0
	y := 0

	for x < 15 {

		if y == 16 {
				y = 0
				x++
			}
		
		for y < 16 {
			

			rl.DrawRectangleRec(i.pixarr[y][x],i.colorarray[y][x])

			y++
		}
	}
}


func main() {
	rl.InitWindow(1280, 720, "hyello")

	rl.SetWindowState(0x00000004)
	rl.SetTargetFPS(120)

	
	iposx :=  rl.GetScreenWidth()/4
	iposy :=  rl.GetScreenHeight()/8
	
	var i Initdata

	var side float32 = 32

	initcolorarray(&i) 
	
	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonDown(0){
			if side < 64 {
				side = side + 1
			}
			
		}

		if rl.IsMouseButtonDown(1){
			if side > 8 {
				side = side - 1
			}
		}

		if rl.IsKeyDown(87){ //W
			iposy -= 2
		}
		
		if rl.IsKeyDown(83){ //S
			iposy += 2
		}
		
		if rl.IsKeyDown(65){//A
			iposx -= 2
		}
		
		if rl.IsKeyDown(68){
			iposx += 2
		}
		if rl.IsKeyDown(82){
			
			iposx =  rl.GetScreenWidth()/4
			iposy =  rl.GetScreenHeight()/8
	
		}
		
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		

		init16(&i,side,iposx,iposy)

		
		
		draw16(&i)

		//		fmt.Println(rl.GetFPS())
		
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

