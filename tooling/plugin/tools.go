package plugin

import "slices"

// Tools gives a plugin access to the manager internals, so plugins can
// coordinate without exporting the state itself.
type Tools struct {
	Pm *PluginManager
}

func (t *Tools) HasFlag(f string) bool {
	return slices.Contains(t.Pm.flags, f)
}

func (t *Tools) SetFlag(f string) {
	t.Pm.flags = append(t.Pm.flags, f)
}

func (t *Tools) IsPluginInstalled(name string) bool {
	_, ok := t.Pm.Plugins[name]
	return ok
}
