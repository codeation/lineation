package draw

import (
	"image"
	"time"

	"github.com/codeation/impress"

	"github.com/codeation/lineation/mindmap"
	"github.com/codeation/lineation/palette"
)

type View struct {
	w          *impress.Window
	windowSize image.Point
	offset     image.Point
	rootBox    *Box
	activeBox  *Box
	markRedraw bool
	isModified bool
}

func NewView(w *impress.Window, box *Box) *View {
	return &View{
		w:          w,
		windowSize: image.Pt(1, 1),
		rootBox:    box,
		activeBox:  box,
	}
}

func (v *View) GetNodes() *mindmap.Node {
	return v.rootBox.GetNodes()
}

func (v *View) Modified(ok bool) {
	if v.isModified != ok {
		v.QueueDraw()
	}
	v.isModified = ok
}

func (v *View) animeOffset(nextOffset image.Point) {
	const steps = 5
	const animeDuration = 100 * time.Millisecond
	for i := 1; i < steps; i++ {
		tempOffset := image.Pt((nextOffset.X*i+v.offset.X*(steps-i))/steps,
			(nextOffset.Y*i+v.offset.Y*(steps-i))/steps)
		v.rootBox.Draw(v.w, tempOffset)
		v.w.Show()
		time.Sleep(animeDuration / steps)
		v.w.Clear()
	}
}

func (v *View) QueueDraw() {
	v.markRedraw = true
}

func (v *View) ReDraw() {
	if !v.markRedraw {
		return
	}
	v.markRedraw = false
	v.rootBox.SplitLeftRight()
	v.rootBox.Align(image.Pt(v.windowSize.X/2, 20))
	nextOffset := v.activeBox.GetOffset(v.windowSize, v.offset)
	nextOffset = v.rootBox.Fit(v.windowSize, nextOffset)
	v.w.Clear()
	if nextOffset != v.offset {
		v.animeOffset(nextOffset)
	}
	v.offset = nextOffset
	v.rootBox.Draw(v.w, v.offset)
	if v.isModified {
		v.w.Fill(image.Rect(2, 2, 8, 8), v.rootBox.pal.Color(palette.ActiveEdge))
	}
	v.w.Show()
	time.Sleep(10 * time.Millisecond)
}

func (v *View) ConfigureSize(size image.Point) {
	if v.windowSize == size {
		return
	}
	v.windowSize = size
	v.w.Size(image.Rect(0, 0, size.X, size.Y))
	v.QueueDraw()
}

func (v *View) KeyDown() {
	next := v.activeBox.down()
	if next == v.activeBox || next == nil {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) KeyUp() {
	next := v.activeBox.up()
	if next == v.activeBox || next == nil {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) Click(point image.Point) {
	next := v.rootBox.Find(point.Add(image.Pt(-v.offset.X, -v.offset.Y)))
	if next == nil || next == v.activeBox {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) KeyLeft() {
	ok := v.activeBox.textBox.Left()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) KeyRight() {
	ok := v.activeBox.textBox.Right()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) RemoveLastChar() {
	ok := v.activeBox.textBox.Backspace()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) InsertChar(alpha rune) {
	v.activeBox.textBox.Insert(alpha)
	v.QueueDraw()
}

func (v *View) AddChildNode() {
	next := v.activeBox.AddChildNode()
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) AddNextNode() {
	if v.activeBox == v.rootBox {
		return
	}
	next := v.activeBox.AddNextNode()
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) DeleteNode() {
	if v.activeBox == v.rootBox {
		return
	}
	next := v.activeBox.DeleteNode()
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}
