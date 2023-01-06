package wrap

type splitter interface {
	Split(text string, edge int) []string
}

type edger interface {
	Edge() int
}

type Wrap struct {
	*Runes
	splitter  splitter
	edger     edger
	texts     []string
	needSplit bool
	lastEdge  int
}

func NewWrap(text string, splitter splitter, edger edger) *Wrap {
	return &Wrap{
		Runes:     NewRunes(text),
		splitter:  splitter,
		edger:     edger,
		needSplit: true,
	}
}

func (w *Wrap) enqueueSplit() {
	w.needSplit = true
}

func (w *Wrap) ensureSplit() {
	if !w.needSplit && w.lastEdge == w.edger.Edge() {
		return
	}
	w.lastEdge = w.edger.Edge()
	w.texts = w.splitter.Split(w.Runes.String(), w.lastEdge)
	w.needSplit = false
}

func (w *Wrap) Lines() int {
	w.ensureSplit()
	if len(w.texts) == 0 {
		return 1
	}
	return len(w.texts)
}

func (w *Wrap) Line(row int) string {
	w.ensureSplit()
	if row >= len(w.texts) {
		return ""
	}
	return w.texts[row]
}

func (w *Wrap) Cursor() (int, int) {
	w.ensureSplit()
	if len(w.texts) == 0 {
		return 0, 0
	}
	offset := 0
	for row := 0; row < len(w.texts); row++ {
		if w.cursor >= offset && w.cursor < offset+len(w.texts[row]) {
			return row, w.cursor - offset
		}
		offset += len(w.texts[row])
	}
	return len(w.texts) - 1, len(w.texts[len(w.texts)-1])
}

func (w *Wrap) Insert(alpha rune) {
	w.Runes.Insert(alpha)
	w.enqueueSplit()
}

func (w *Wrap) Backspace() bool {
	if !w.Runes.Backspace() {
		return false
	}
	w.enqueueSplit()
	return true
}
