package main

import "net/url"

type errors map[string][]string

func (e errors) Get(field string) string {
	errorsSlice := e[field]

	if len(errorsSlice) == 0 {
		return ""
	}

	return errorsSlice[0]
}

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

type Form struct {
	Data   url.Values
	Errors errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data:   data,
		Errors: map[string][]string{},
	}
}
