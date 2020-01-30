package draw

// Logic is the struct which is saved to disk
type Logic struct {
	Pages []*Page

	CurrentPage int

	Light Lighter `json:"-"`
}

// Lighter is able to light stuff up
type Lighter interface {
	Light(int, int, int, int) error
	Clear() error
}

// NewLogic returns a new game logic
func NewLogic(l Lighter) *Logic {
	p := make([]*Page, 16)
	for i := range p {
		p[i] = NewPage()
	}

	return &Logic{Pages: p, Light: l}
}

// Press accepts keypresses
func (l *Logic) KeyDown(x, y int) {
	// if x or y is at 8, we are selecting a page
	if y == 8 {
		l.CurrentPage = x + 8
		l.Render()
		return
	}
	if x == 8 {
		l.CurrentPage = y
		l.Render()
		return
	}

	// pixel presses should be handled without having to render stuff from scratch
	l.Pages[l.CurrentPage].Pixels[x][y].Toggle(func(r, g int) {
		l.Light.Light(x, y, r, g)
	})

}

func (l *Logic) Render() {
	err := l.Light.Clear()
	if err != nil {
		panic(err)
	}

	if l.CurrentPage < 8 {
		l.Light.Light(8, l.CurrentPage, 3, 0)
	} else {
		l.Light.Light(l.CurrentPage-8, 8, 3, 0)
	}
	l.Pages[l.CurrentPage].RenderTo(l.Light.Light)

}
