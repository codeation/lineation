package text

func (t *Text) Insert(alpha rune) {
	t.content.Insert(alpha)
}

func (t *Text) Backspace() bool {
	return t.content.Backspace()
}

func (t *Text) End() {
	t.content.End()
}

func (t *Text) String() string {
	return t.content.String()
}

func (t *Text) Left() bool {
	return t.content.Left()
}

func (t *Text) Right() bool {
	return t.content.Right()
}

func (t *Text) Lines() int {
	return t.content.Lines()
}

func (t *Text) Line(no int) string {
	return t.content.Line(no)
}

func (t *Text) Cursor() (int, int) {
	return t.content.Cursor()
}
