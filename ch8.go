package main

import (
	"ch8/kochline"
	"ch8/sb"
	"github.com/gen2brain/raylib-go/raylib"
)

//DisplayKochLinePoint : ...
func DisplayKochLinePoint(kps *kochline.KochLines) {
	for _, kl := range kps.Klines {
		//		kl.Display()
		rl.DrawLine(int32(kl.Start.X), int32(kl.Start.Y),
			int32(kl.End.X), int32(kl.End.Y), rl.Black)
	}
}
func main() {
	kls := kochline.NewKochLines(sb.Vec2{0, 300}, sb.Vec2{800, 300})
	for i := 0; i < 5; i++ {
		kls.Generate()
	}
	screenWidth := int32(800)
	screenHeight := int32(600)
	// Init window is topmost func for raylib
	rl.InitWindow(screenWidth, screenHeight, "Ch - 8 KochLine")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DisplayKochLinePoint(kls)
		rl.EndDrawing()
	}

}
