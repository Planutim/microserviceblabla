package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

// MyError is an error implementation that includes a time and message.
type ComplexError struct {
	Message string
	Code    int
	frame   xerrors.Frame
}

func (ce ComplexError) FormatError(p xerrors.Printer) error {
	p.Printf("%d %s", ce.Code, ce.Message)
	ce.frame.Format(p)
	return nil
}
func (ce ComplexError) Format(f fmt.State, c rune) {
	xerrors.FormatError(ce, f, c)
}

func (ce ComplexError) Error() string {
	return fmt.Sprint(ce)
}

func someComplexErrorHappens() error {
	complexErr := ComplexError{
		Code:    1234,
		Message: "this is my error message",
		frame:   xerrors.Caller(1), // skip the first frame
	}
	return xerrors.Errorf(
		"uh oh! something terribly complex happened: %w", complexErr)
}

func main() {
	cerr := someComplexErrorHappens()
	var originalErr ComplexError
	if xerrors.As(cerr, &originalErr) {
		fmt.Println("yes")
	}
}
