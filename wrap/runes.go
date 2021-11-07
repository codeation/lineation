package wrap

import (
	"unicode/utf8"
)

type Runes struct {
	source string
	cursor int
}

func NewRunes(s string) *Runes {
	return &Runes{
		source: s,
	}
}

func (r *Runes) String() string {
	return r.source
}

func (r *Runes) Home() {
	r.cursor = 0
}

func (r *Runes) End() {
	r.cursor = len(r.source)
}

func (r *Runes) Left() bool {
	if r.cursor <= 0 {
		return false
	}
	_, size := utf8.DecodeLastRuneInString(r.source[:r.cursor])
	r.cursor -= size
	return true
}

func (r *Runes) Right() bool {
	if r.cursor >= len(r.source) {
		return false
	}
	_, size := utf8.DecodeRuneInString(r.source[r.cursor:])
	r.cursor += size
	return true
}

func (r *Runes) Insert(alpha rune) {
	r.source = r.source[:r.cursor] + string(alpha) + r.source[r.cursor:]
	_, size := utf8.DecodeRuneInString(r.source[r.cursor:])
	r.cursor += size
}

func (r *Runes) Backspace() bool {
	if r.cursor <= 0 {
		return false
	}
	_, size := utf8.DecodeLastRuneInString(r.source[:r.cursor])
	r.source = r.source[:r.cursor-size] + r.source[r.cursor:]
	r.cursor -= size
	return true
}
