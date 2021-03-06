package main

import(
	"github.com/gen2brain/raylib-go/raylib"
)

// abstract type that defines 2 dimentional array of type color
type COLORSLICE_2d [][]rl.Color


// The structure that holds data for current workings
type DATASTRUCT struct{

	//This is an array of recs used to denote resolution
	recarr [][]rl.Rectangle
	// These are the arrays used for colors to be set on canvas 
	colorarr    COLORSLICE_2D 
	colorbuffer COLORSLICE_2D
	colorundo   COLORSLICE_2D
	// Acts as the single pixel
	pixrec rl.Rectangle

	//Arry that holds the states of the undobuffer
	states []COLORSLICE_2d

	//initial position for the array where dwawing takes place
	inposx float32
	imposy float32

	//side of the drawing space

	sideofbigbox int32

}

func (d DATASTRUCT) init16(){

	d.sideofbigbox = 16
	
	d.pixrec = rl.Rectangle{
		X : imposx,
		Y : imposy,
		Height : 32,
		Width : 32,
	}

	
	
}

//The method is use to fill the array of rectangles with the rectangles that represent one pixel

func(d DATASTRUCT) boxarraypopulator(){

	var x,y int32 = 0
	
	for x != d.sideofbigbox {
		if y != d.sideofbigbox {
			d.recarr[y][x] = d.pixrec
			y++
			d.pixrec.X = d.pixrec.X + 32
			
		}
		if y == d.sideofbigbox {
			x++
			d.pixrec.Y = d.pixrec.Y + 32
		}
		
	}

	
}

func main(){

	 const height int32 = 720
	const width int32 = 1280 
  	
	rl.InitWindow(width,height, "heloo windo")

	for !rl.WindowShouldClose(){

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
	
}




