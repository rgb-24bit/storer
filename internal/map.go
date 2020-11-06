package internal

func MergeMap(from, to map[string]interface{}) map[string]interface{} {
	for fk, fv := range from {
		tv := to[fk]
		if isMapType(fv) && isMapType(tv) {
			to[fk] = mergeMap(toMapStringInterface(fv), toMapStringInterface(tv))
		} else {
			to[fk] = fv
		}
	}
	return to
}

func IsMapType(i interface{}) bool {
	_, i2ip := i.(map[interface{}]interface{})
	_, s2ip := i.(map[string]interface{})
	return i2ip || s2ip
}

func ToMapStringInterface(i interface{}) map[string]interface{} {
	if m, ok := i.(map[string]interface{}); ok {
		return m
	}
	m := make(map[string]interface{})
	for k, v := range i.(map[interface{}]interface{}) {
		if isMapType(v) {
			m[k.(string)] = toMapStringInterface(v)
		} else {
			m[k.(string)] = v
		}
	}
	return m
}

