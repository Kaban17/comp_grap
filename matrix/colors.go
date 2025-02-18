package matrix

import (
	"fmt"
)

type Color interface {
	Get() string
}
type GrayColor struct {
	Value uint8
}
type RGBColor struct {
	R, G, B uint8
}

func (g GrayColor) Get() string {
	return fmt.Sprintf("Gray(%d)", g.Value)
}
func (c RGBColor) Get() string {
	return fmt.Sprintf("RGB(%d, %d, %d)", c.R, c.G, c.B)
}
