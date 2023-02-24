package forms

type errors map[string][]string

// Add adds am error message to a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	error := e[field]

	if len(error) == 0 {
		return ""
	}

	return error[0]
}
