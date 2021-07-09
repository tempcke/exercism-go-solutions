package flatten

// Flatten nested list into a flat list excluding nils
func Flatten(in interface{}) []interface{} {
	out := []interface{}{}

	for _, v := range in.([]interface{}) {
		if v != nil {
			switch v.(type) {
			case []interface{}:
				out = append(out, Flatten(v)...)
			default:
				out = append(out, v)
			}
		}
	}

	return out
}
