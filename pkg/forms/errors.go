package forms

type errors map[string][]string

// Add an error to our map of errors.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get fetches the first error message for the given field.
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}
