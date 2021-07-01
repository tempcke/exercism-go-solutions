package erratum

import "errors"

// ErrFrobPanic is returned when the call to Resource.Frob causes a panic
var ErrFrobPanic = errors.New("meh")

// Use the resourceOpener to get a resource and Frob the input
func Use(o ResourceOpener, input string) error {
	var resource Resource
	var err error

	retryLimit := 10
	for i := 0; i < retryLimit; i++ {
		if resource, err = o(); err == nil {
			break
		}
		if !errors.As(err, &TransientError{}) {
			return err
		}
	}

	defer resource.Close()
	if ok := frob(resource, input); !ok {
		return ErrFrobPanic
	}
	return nil
}

func frob(resource Resource, input string) bool {
	defer frobPanicCheck(resource)
	resource.Frob(input)
	return true
}

func frobPanicCheck(resource Resource) {
	if r := recover(); r != nil {
		if err, ok := r.(FrobError); ok {
			resource.Defrob(err.defrobTag)
		}
	}
}
