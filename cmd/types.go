package cmd

var (
	Types = map[string]struct{}{
		"artists":  {},
		"labels":   {},
		"masters":  {},
		"releases": {},
	}
)

func getTypes(types []string) []string {
	if len(types) == 0 {
		return getKeys(Types)
	}

	res := map[string]struct{}{}
	for _, t := range types {
		if t[len(t)-1:] != "s" {
			t += "s"
		}
		if containsKey(Types, t) {
			res[t] = struct{}{}
		}
	}
	return getKeys(res)
}

func getKeys[T comparable, V any](in map[T]V) []T {
	v := make([]T, 0)
	for k := range in {
		v = append(v, k)
	}
	return v
}

func containsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
