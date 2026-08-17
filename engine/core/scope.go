package core

import "github.com/pt-main/lc/public/errors"

type ScopeType map[string]interface{}

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
