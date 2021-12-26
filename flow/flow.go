package flow

import "github.com/fogleman/gg"

// A5 paper dimensions
const (
	width  = 420
	height = 595
)

// RenderAndSave renders a context and save it as png in disk
func RenderAndSave() {
	ctx := gg.NewContext(width, height)
	// Set white background
	ctx.DrawRectangle(0, 0, width, height)
	ctx.SetRGB(1, 1, 1)
	ctx.Fill()
	ctx.SavePNG("flowchart.png")
}
