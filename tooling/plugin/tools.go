package plugin

import "slices"

// # Tools
//
// PluginManager tools for simple plugins using
type Tools struct {
	Pm *PluginManager
}

func (t *Tools) HasFlag(f string) bool {
	return slices.Contains(t.Pm.flags, f)
}

func (t *Tools) SetFlag(f string) {
	t.Pm.flags = append(t.Pm.flags, f)
}

func (t *Tools) IsPluginInstaled(p string) bool {
	_, ok := t.Pm.Plugins[p]
	return ok
}
