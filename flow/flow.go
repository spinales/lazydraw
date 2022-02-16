package flow

import "github.com/fogleman/gg"

// A5 paper dimensions
const (
	width  = 420
	height = 595
)

type Flowchart struct {
	ctx *gg.Context
}

// Creates an empty flowchart
func NewFlowchart() *Flowchart {
	ctx := gg.NewContext(width, height)
	// Set white brackground
	ctx.DrawRectangle(0, 0, width, height)
	ctx.SetRGB(1, 1, 1)
	ctx.Fill()

	return &Flowchart{
		ctx: ctx,
	}
}

// Add adds a flowchar component into the f
// TODO: Change component type
func (f *Flowchart) Add(component Component) {
	component.Add(f.ctx)
}

// RenderAndSave renders a flowchart and save it as png in disk
func (f *Flowchart) RenderAndSave() {
	f.ctx.SavePNG("flowchart.png")
}
