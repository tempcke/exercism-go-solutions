package erratum

import (
	"errors"
	"fmt"
)

// RetryLimit is used so we are confident the loop has an exit
// by making it exported it allows the client to change the limit if needed
// I don't like globals for this purpose but the challenge conditions didn't allow for anything else really
var RetryLimit = 10

// Use the resourceOpener to get a resource and Frob the input
func Use(o ResourceOpener, input string) (err error) {
	resource, err := openResource(o)
	if err != nil {
		return err
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			if e, ok := r.(FrobError); ok {
				resource.Defrob(e.defrobTag)
			}
		}
	}()

	resource.Frob(input)

	return nil
}

func openResource(o ResourceOpener) (r Resource, err error) {
	for i := 0; ; i++ {
		if r, err = o(); err == nil {
			break
		}
		if !errors.As(err, &TransientError{}) {
			return r, err
		}
		if i >= RetryLimit {
			return r, errors.New("retry limit exceeded")
		}
	}
	return r, nil
}
