package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
	end   Point
}

func (square *Square) End() Point {
	square.end = Point{
		x: square.start.x + int(square.a),
		y: square.start.y + int(square.a),
	}
	return square.end
}

func (square *Square) Area() uint {
	return square.a * square.a
}

func (square *Square) Perimeter() uint {
	return square.a * 4
}
