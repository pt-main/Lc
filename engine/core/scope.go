package core

import (
	"reflect"
	"sync"

	"github.com/pt-main/lc/public/errors"
)

// ScopeType stays a plain map so that existing code writing scope[k] = v
// keeps compiling. Synchronised access goes through the helpers below.
type ScopeType map[string]interface{}

// scopeGuards holds one mutex per scope map.
//
// The entry is keyed by the map pointer, but the pointer alone is unsafe: the
// runtime may hand the same address to a new map once the original is
// collected, which would silently give the new scope an unrelated scope's
// mutex. Holding the map itself in the entry keeps it alive, so its address
// cannot be reused while the guard is registered.
//
// The table is only read and written under scopeGuardsMu, which is never held
// while a scope's own lock is taken, so the two cannot deadlock.
var (
	scopeGuardsMu sync.Mutex
	scopeGuards   = make(map[uintptr]*scopeGuardEntry)
)

// nilScopeGuard backs a nil scope, which has nothing to protect.
var nilScopeGuard sync.RWMutex

type scopeGuardEntry struct {
	scope ScopeType
	mu    sync.RWMutex
}

func scopeGuard(st ScopeType) *sync.RWMutex {
	if st == nil {
		return &nilScopeGuard
	}
	key := reflect.ValueOf(st).Pointer()
	scopeGuardsMu.Lock()
	entry, ok := scopeGuards[key]
	if !ok || entry.scope == nil {
		entry = &scopeGuardEntry{scope: st}
		scopeGuards[key] = entry
	}
	scopeGuardsMu.Unlock()
	return &entry.mu
}

// ReleaseScopeGuard drops the cached guard of a scope that is about to be
// discarded, so the table does not grow for the lifetime of the process. The
// caller must guarantee that no other goroutine still uses the scope.
func ReleaseScopeGuard(st ScopeType) {
	if st == nil {
		return
	}
	key := reflect.ValueOf(st).Pointer()
	scopeGuardsMu.Lock()
	if entry, ok := scopeGuards[key]; ok && entry.scope != nil &&
		reflect.ValueOf(entry.scope).Pointer() == key {
		delete(scopeGuards, key)
	}
	scopeGuardsMu.Unlock()
}

// Err errors.ScopeGetError.
// With meta: EMK(0, "string") - key
func ScopeGet[T any](st ScopeType, what string) (T, ErrorInterface) {
	var nul T
	val, ok := st[what]
	if !ok {
		return nul, Err(errors.ScopeGetError, "Invalid key: %v", what).
			WithMeta(EMK(0, "string"), what)
	}
	res, ok := val.(T)
	if !ok {
		return nul, Err(errors.ScopeGetError, "Invalid type for key: %v", what).
			WithMeta(EMK(0, "string"), what)
	}
	return res, nil
}

// ScopeGetSynced is ScopeGet guarded by the scope lock, for scopes written
// from one goroutine while another reads them.
func ScopeGetSynced[T any](st ScopeType, what string) (T, ErrorInterface) {
	scopeGuard(st).RLock()
	defer scopeGuard(st).RUnlock()
	return ScopeGet[T](st, what)
}

// ScopeSetSynced stores a value under the scope lock.
func ScopeSetSynced(st ScopeType, what string, value interface{}) {
	scopeGuard(st).Lock()
	defer scopeGuard(st).Unlock()
	st[what] = value
}

// ScopeDeleteSynced removes a key under the scope lock.
func ScopeDeleteSynced(st ScopeType, what string) {
	scopeGuard(st).Lock()
	defer scopeGuard(st).Unlock()
	delete(st, what)
}
