package gdka

import "github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gdk"

func wrapEventAsEventButton(v *event) *eventButton {
	wrapped, _ := wrapEventButton(&gdk.EventButton{v.Event}, nil)
	return wrapped
}
