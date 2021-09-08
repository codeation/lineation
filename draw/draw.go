package draw

import (
	"time"
	"unicode/utf8"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/mindmap"
)

type View struct {
	w          *impress.Window
	windowSize impress.Size
	offset     impress.Point
	rootBox    *Box
	activeBox  *Box
}

func NewView(w *impress.Window, box *Box) *View {
	return &View{
		w:          w,
		windowSize: impress.NewSize(1, 1),
		rootBox:    box,
		activeBox:  box,
	}
}

func (v *View) GetNodes() *mindmap.Node {
	return v.rootBox.GetNodes()
}

func (v *View) animeOffset(nextOffset impress.Point) {
	const steps = 5
	const animeDuration = 100 * time.Millisecond
	for i := 1; i < steps; i++ {
		tempOffset := impress.NewPoint((nextOffset.X*i+v.offset.X*(steps-i))/steps,
			(nextOffset.Y*i+v.offset.Y*(steps-i))/steps)
		v.rootBox.Draw(v.w, tempOffset)
		v.w.Show()
		time.Sleep(animeDuration / steps)
		v.w.Clear()
	}
}

func (v *View) Draw(withAnime bool) {
	nextOffset := v.activeBox.GetOffset(v.windowSize, v.offset)
	nextOffset = v.rootBox.Fit(v.windowSize, nextOffset)
	v.w.Clear()
	if nextOffset != v.offset && withAnime {
		v.animeOffset(nextOffset)
	}
	v.offset = nextOffset
	v.rootBox.Draw(v.w, v.offset)
	v.w.Show()
}

func (v *View) ConfigureSize(size impress.Size) {
	if v.windowSize == size {
		return
	}
	v.windowSize = size
	v.w.Size(impress.NewRect(0, 0, size.Width, size.Height))
	v.rootBox.Align(impress.NewPoint(size.Width/2, 20))
	v.Draw(false)
}

func (v *View) KeyDown() {
	next := v.activeBox.down()
	if next == v.activeBox || next == nil {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.Draw(true)
}

func (v *View) KeyUp() {
	next := v.activeBox.up()
	if next == v.activeBox || next == nil {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.Draw(true)
}

func (v *View) Click(point impress.Point) {
	next := v.rootBox.Find(point.Move(impress.NewPoint(-v.offset.X, -v.offset.Y)))
	if next == nil || next == v.activeBox {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.Draw(true)
}

func (v *View) RemoveLastChar() {
	text := v.activeBox.GetText()
	if text == "" {
		return
	}
	_, lastsize := utf8.DecodeLastRune([]byte(text))
	v.activeBox.SetText(text[:len(text)-lastsize])
	v.rootBox.Align(impress.NewPoint(v.windowSize.Width/2, 20))
	v.Draw(false)
}

func (v *View) InsertChar(alpha rune) {
	v.activeBox.SetText(v.activeBox.GetText() + string(alpha))
	v.rootBox.Align(impress.NewPoint(v.windowSize.Width/2, 20))
	v.Draw(false)
}

func (v *View) AddChildNode() {
	next := v.activeBox.AddChildNode()
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.rootBox.Align(impress.NewPoint(v.windowSize.Width/2, 20))
	v.Draw(false)
}

func (v *View) AddNextNode() {
	next := v.activeBox.AddNextNode()
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.rootBox.Align(impress.NewPoint(v.windowSize.Width/2, 20))
	v.Draw(false)
}

func (v *View) DeleteNode() {
	if v.activeBox == v.rootBox {
		return
	}
	next := v.activeBox.DeleteNode()
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.rootBox.Align(impress.NewPoint(v.windowSize.Width/2, 20))
	v.Draw(false)
}
