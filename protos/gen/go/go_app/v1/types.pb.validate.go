// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: go_app/v1/types.proto

package go_appv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Hello with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Hello) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Hello with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HelloMultiError, or nil if none found.
func (m *Hello) ValidateAll() error {
	return m.validate(true)
}

func (m *Hello) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HelloValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HelloValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HelloValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HelloValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HelloValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HelloType

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.PersonName != nil {
		// no validation rules for PersonName
	}

	if len(errors) > 0 {
		return HelloMultiError(errors)
	}

	return nil
}

// HelloMultiError is an error wrapping multiple validation errors returned by
// Hello.ValidateAll() if the designated constraints aren't met.
type HelloMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloMultiError) AllErrors() []error { return m }

// HelloValidationError is the validation error returned by Hello.Validate if
// the designated constraints aren't met.
type HelloValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloValidationError) ErrorName() string { return "HelloValidationError" }

// Error satisfies the builtin error interface
func (e HelloValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHello.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloValidationError{}

// Validate checks the field values on Todo with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Todo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Todo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TodoMultiError, or nil if none found.
func (m *Todo) ValidateAll() error {
	return m.validate(true)
}

func (m *Todo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TodoType

	// no validation rules for TodoName

	// no validation rules for Priority

	// no validation rules for Completed

	if m.Id != nil {
		// no validation rules for Id
	}

	if len(errors) > 0 {
		return TodoMultiError(errors)
	}

	return nil
}

// TodoMultiError is an error wrapping multiple validation errors returned by
// Todo.ValidateAll() if the designated constraints aren't met.
type TodoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TodoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TodoMultiError) AllErrors() []error { return m }

// TodoValidationError is the validation error returned by Todo.Validate if the
// designated constraints aren't met.
type TodoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodoValidationError) ErrorName() string { return "TodoValidationError" }

// Error satisfies the builtin error interface
func (e TodoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodoValidationError{}
