package plugin

import "slices"

type Tools struct {
	Pm PluginManager
}

func (t *Tools) HasFlag(f string) bool {
	return slices.Contains(t.Pm.Flags, f)
}

func (t *Tools) SetFlag(f string) {
	t.Pm.Flags = append(t.Pm.Flags, f)
}

func (t *Tools) IsPluginInstaled(p string) bool {
	_, ok := t.Pm.Plugins[p]
	return ok
}
