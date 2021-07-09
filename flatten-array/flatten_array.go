package flatten

// Flatten nested list into a flat list excluding nils
func Flatten(in interface{}) []interface{} {
	out := []interface{}{}

	for _, val := range in.([]interface{}) {
		switch v := val.(type) {
		case []interface{}:
			if v != nil {
				out = append(out, Flatten(v)...)
			}
		default:
			if v != nil {
				out = append(out, v)
			}
		}
	}

	return out
}
