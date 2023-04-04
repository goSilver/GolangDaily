package g_decorator

type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct {
}

func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 有颜色的正方形，对普通正方形对增强
type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{square: square, color: color}
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}
