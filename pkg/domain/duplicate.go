package domain

import "errors"

type IDuplicate interface {
	IsDuplicate(val string) bool
	TrySaveNewValue(val string) error
}

type Duplicate struct {
	vals map[string]bool
}

func New() IDuplicate {
	return Duplicate{vals: make(map[string]bool)}
}

//IsDuplicate checks if a value given by parameter was already saved previously.
func (d Duplicate) IsDuplicate(value string) bool {
	return d.vals[value]
}

//TrySaveNewValue tries to save a new value
func (d Duplicate) TrySaveNewValue(value string) error {
	if d.IsDuplicate(value) {
		return errors.New("value already registered")
	}
	d.vals[value] = true
	return nil
}