package gtka

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gtki"
)

type treePath struct {
	*gtk.TreePath
}

func wrapTreePathSimple(v *gtk.TreePath) *treePath {
	if v == nil {
		return nil
	}
	return &treePath{v}
}

func wrapTreePath(v *gtk.TreePath, e error) (*treePath, error) {
	return wrapTreePathSimple(v), e
}

func unwrapTreePath(v gtki.TreePath) *gtk.TreePath {
	if v == nil {
		return nil
	}
	return v.(*treePath).TreePath
}
