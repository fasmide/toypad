package draw

// Page represents a "drawing"
type Page struct {
	// Pixels: [x][y]Pixel
	Pixels [][]Pixel
}

// NewPage with all buttons turned off
func NewPage() *Page {
	a := make([][]Pixel, 8)
	for i := range a {
		a[i] = make([]Pixel, 8)
	}

	return &Page{Pixels: a}
}

func (p *Page) RenderTo(f func(x, y, r, g int) error) {
	for x, line := range p.Pixels {
		for y, pixel := range line {
			pixel.RenderTo(func(r, g int) {
				f(x, y, r, g)
			})
		}
	}
}
