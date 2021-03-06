package gdka

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gdk"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gdki"
)

type pixbuf struct {
	*gdk.Pixbuf
}

func wrapPixbufSimple(v *gdk.Pixbuf) *pixbuf {
	if v == nil {
		return nil
	}
	return &pixbuf{v}
}

func wrapPixbuf(v *gdk.Pixbuf, e error) (*pixbuf, error) {
	return wrapPixbufSimple(v), e
}

func UnwrapPixbuf(v gdki.Pixbuf) *gdk.Pixbuf {
	if v == nil {
		return nil
	}
	return v.(*pixbuf).Pixbuf
}
