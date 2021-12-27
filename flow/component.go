package flow

import "github.com/fogleman/gg"

// Component represents an object that can be append
// into a flowchart.
type Component interface {
	Add(ctx *gg.Context)
}
