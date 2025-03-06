package triangle

type Point2D struct {
	X, Y float64
}
type TriangleVertices struct {
	X0, Y0 float64
	X1, Y1 float64
	X2, Y2 float64
}
type BarycentricCoordinates struct {
	X, Y, Z float64
}

func (p *Point2D) Bar_coord(t *TriangleVertices) BarycentricCoordinates {
	lambda0 := ((p.X-t.X0)*(t.Y1-t.Y2) - (t.X1-t.X2)*(p.Y-t.Y2)) / ((t.X0-t.X2)*(t.Y1-t.Y2) - (t.X1-t.X2)*(t.Y0-t.Y2))
	lambda1 := ((t.X0-t.X2)*(p.Y-t.Y2) - (p.X-t.X2)*(t.Y0-t.Y2)) / ((t.X0-t.X2)*(t.Y1-t.Y2) - (t.X1-t.X2)*(t.Y0-t.Y2))
	lambda2 := 1.0 - lambda0 - lambda1
	return BarycentricCoordinates{X: lambda0, Y: lambda1, Z: lambda2}
}
