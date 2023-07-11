package draw

import (
	"image"

	"github.com/codeation/lineation/draw/box"
)

type dragState struct {
	catchPoint           image.Point
	oldParent, newParent *box.Box
	beforeBox            *box.Box
	IsGridCleared        bool
}

func (state *dragState) IsDroppable() bool {
	return state.newParent != nil
}

func (v *View) Catch(eventPoint image.Point) (*dragState, bool) {
	if v.activeBox == nil || v.activeBox == v.rootBox {
		return nil, false
	}
	pt := eventPoint.Sub(v.offset)
	if !v.activeBox.In(pt) {
		return nil, false
	}

	return &dragState{
		catchPoint: eventPoint,
	}, true
}

func (v *View) Drag(state *dragState, eventPoint image.Point) {
	plus := eventPoint.Sub(state.catchPoint.Sub(v.offset))
	state.newParent, state.beforeBox = v.rootBox.FindOther(
		image.Pt(eventPoint.X, v.activeBox.Point().Y+plus.Y).Sub(v.offset), v.activeBox)
	v.activeBox.Drag(plus)
}

func (v *View) DrawDrag(state *dragState) {
	if !state.IsGridCleared {
		v.activeBox.Raise()
		v.w.Clear()
		v.rootBox.DrawGrid(v.w, v.offset, v.activeBox)
		v.w.Show()
		state.IsGridCleared = true
	}
	if state.newParent != state.oldParent {
		if state.oldParent != nil {
			state.oldParent.DeEmphasize()
		}
		if state.newParent != nil {
			state.newParent.Emphasize()
		}
		state.oldParent = state.newParent
	}
}

func (v *View) DrawRemain(state *dragState) {
	v.w.Clear()
	v.rootBox.DrawGrid(v.w, v.offset, nil)
	v.w.Show()
	if state.newParent != nil {
		state.newParent.DeEmphasize()
	}
	v.activeBox.Drag(image.Pt(0, 0).Add(v.offset))
}

func (v *View) Drop(state *dragState, eventPoint image.Point) {
	state.newParent.DeEmphasize()
	state.newParent.Adopt(v.activeBox, state.beforeBox)
}
