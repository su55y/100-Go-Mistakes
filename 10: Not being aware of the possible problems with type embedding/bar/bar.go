package bar

type Bar struct {
	baz int
}

func (b *Bar) Baz() int {
	return b.baz
}

func (b *Bar) SetBaz(baz int) {
	b.baz = baz
}
