package forms

// can have multiple errors
type errors map[string][]string

// Add adds and error msg for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	// first index of error string
	return es[0]
}
