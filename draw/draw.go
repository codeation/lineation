package draw

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/lineation/draw/box"
	"github.com/codeation/lineation/draw/modifiedstatus"
	"github.com/codeation/lineation/mindmap"
)

type syncer interface {
	Sync()
	Chan() <-chan event.Eventer
}

type View struct {
	w              *impress.Window
	windowSize     image.Point
	offset         image.Point
	rootBox        *box.Box
	activeBox      *box.Box
	modifiedStatus *modifiedstatus.ModifiedStatus
	markRedraw     bool
}

func NewView(w *impress.Window, box *box.Box, modifiedStatus *modifiedstatus.ModifiedStatus) *View {
	return &View{
		w:              w,
		windowSize:     image.Pt(1, 1),
		rootBox:        box,
		activeBox:      box,
		modifiedStatus: modifiedStatus,
	}
}

func (v *View) GetNodes() *mindmap.Node {
	return v.rootBox.GetNodes()
}

func (v *View) Modified(ok bool) {
	v.modifiedStatus.Modified(ok)
}

/*
func (v *View) animeOffset(nextOffset image.Point, s syncer) {
	const steps = 5
	const animeDuration = 100 * time.Millisecond
	for i := 1; i < steps; i++ {
		tempOffset := image.Pt((nextOffset.X*i+v.offset.X*(steps-i))/steps,
			(nextOffset.Y*i+v.offset.Y*(steps-i))/steps)
		v.rootBox.Draw(tempOffset)
		v.rootBox.DrawGrid(v.w, tempOffset, nil)
		v.w.Show()
		since := time.Now()
		s.Sync()
		if len(s.Chan()) > 0 {
			break
		}
		remains := animeDuration/steps - time.Since(since)
		if remains > 0 {
			time.Sleep(remains)
		}
		v.w.Clear()
	}
}
*/

func (v *View) QueueDraw() {
	v.markRedraw = true
}

func (v *View) ReDraw(s syncer) {
	if !v.markRedraw {
		return
	}
	v.markRedraw = false
	shifted := v.rootBox.SplitLeftRight()
	aligned := v.rootBox.Align(image.Pt(v.windowSize.X/2, 0))
	nextOffset := v.activeBox.GetOffset(v.windowSize, v.offset)
	nextOffset = v.rootBox.Fit(v.windowSize, nextOffset)
	shifted = shifted || aligned || v.offset != nextOffset
	if shifted {
		v.w.Clear()
		/*
			if nextOffset != v.offset {
				v.animeOffset(nextOffset, s)
			}
		*/
		v.offset = nextOffset
	}
	v.rootBox.Draw(v.offset)
	if shifted {
		v.rootBox.DrawGrid(v.w, v.offset, nil)
		v.w.Show()
	}
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
	next := v.activeBox.Down()
	if next == v.activeBox || next == nil {
		return
	}
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) KeyUp() {
	next := v.activeBox.Up()
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
	ok := v.activeBox.Left()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) KeyRight() {
	ok := v.activeBox.Right()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) RemoveLastChar() {
	ok := v.activeBox.Backspace()
	if !ok {
		return
	}
	v.QueueDraw()
}

func (v *View) InsertChar(alpha rune) {
	v.activeBox.Insert(alpha)
	v.QueueDraw()
}

func (v *View) AddChildNode(app *impress.Application) {
	next := v.activeBox.AddChildNode(app)
	v.activeBox.SetActive(false)
	v.activeBox = next
	v.activeBox.SetActive(true)
	v.QueueDraw()
}

func (v *View) AddNextNode(app *impress.Application) {
	if v.activeBox == v.rootBox {
		return
	}
	next := v.activeBox.AddNextNode(app)
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
