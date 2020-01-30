package draw

// Pixel represents a Pixel
type Pixel struct {
	R bool
	G bool
}

type Btn func(int, int)

// Toggle sets state and toggles a Btn
func (p *Pixel) Toggle(b Btn) {
	if !p.R && !p.G {
		p.R = true
		b(3, 0)
		return
	}
	if p.R && p.G {
		p.R = false
		p.G = false
		b(0, 0)
		return
	}
	if p.R {
		p.R = false
		p.G = true
		b(0, 3)
		return
	}
	if p.G {
		p.R = true
		p.G = true
		b(3, 3)
		return
	}
}

func (p *Pixel) RenderTo(b Btn) {
	if !p.R && !p.G {
		b(0, 0)
		return
	}
	if p.R && p.G {
		b(3, 3)
		return
	}
	if p.R {
		b(3, 0)
		return
	}
	if p.G {
		b(0, 3)
		return
	}
}
