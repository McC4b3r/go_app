// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: go_app/v1/requests.proto

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

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Limit

	// no validation rules for Offset

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteRequestMultiError, or
// nil if none found.
func (m *DeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteRequestMultiError(errors)
	}

	return nil
}

// DeleteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRequestMultiError) AllErrors() []error { return m }

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

// Validate checks the field values on DeleteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteResponseMultiError,
// or nil if none found.
func (m *DeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteResponseMultiError(errors)
	}

	return nil
}

// DeleteResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResponseMultiError) AllErrors() []error { return m }

// DeleteResponseValidationError is the validation error returned by
// DeleteResponse.Validate if the designated constraints aren't met.
type DeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResponseValidationError) ErrorName() string { return "DeleteResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate checks the field values on Hellos with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Hellos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Hellos with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HellosMultiError, or nil if none found.
func (m *Hellos) ValidateAll() error {
	return m.validate(true)
}

func (m *Hellos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetHellos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HellosValidationError{
						field:  fmt.Sprintf("Hellos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HellosValidationError{
						field:  fmt.Sprintf("Hellos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HellosValidationError{
					field:  fmt.Sprintf("Hellos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return HellosMultiError(errors)
	}

	return nil
}

// HellosMultiError is an error wrapping multiple validation errors returned by
// Hellos.ValidateAll() if the designated constraints aren't met.
type HellosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HellosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HellosMultiError) AllErrors() []error { return m }

// HellosValidationError is the validation error returned by Hellos.Validate if
// the designated constraints aren't met.
type HellosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HellosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HellosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HellosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HellosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HellosValidationError) ErrorName() string { return "HellosValidationError" }

// Error satisfies the builtin error interface
func (e HellosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHellos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HellosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HellosValidationError{}

// Validate checks the field values on Todos with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Todos) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Todos with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TodosMultiError, or nil if none found.
func (m *Todos) ValidateAll() error {
	return m.validate(true)
}

func (m *Todos) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTodos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TodosValidationError{
						field:  fmt.Sprintf("Todos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TodosValidationError{
						field:  fmt.Sprintf("Todos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TodosValidationError{
					field:  fmt.Sprintf("Todos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TodosMultiError(errors)
	}

	return nil
}

// TodosMultiError is an error wrapping multiple validation errors returned by
// Todos.ValidateAll() if the designated constraints aren't met.
type TodosMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TodosMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TodosMultiError) AllErrors() []error { return m }

// TodosValidationError is the validation error returned by Todos.Validate if
// the designated constraints aren't met.
type TodosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodosValidationError) ErrorName() string { return "TodosValidationError" }

// Error satisfies the builtin error interface
func (e TodosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodosValidationError{}

// Validate checks the field values on UpsertHellosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpsertHellosRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertHellosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpsertHellosRequestMultiError, or nil if none found.
func (m *UpsertHellosRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertHellosRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetHellos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpsertHellosRequestValidationError{
						field:  fmt.Sprintf("Hellos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpsertHellosRequestValidationError{
						field:  fmt.Sprintf("Hellos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpsertHellosRequestValidationError{
					field:  fmt.Sprintf("Hellos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpsertHellosRequestMultiError(errors)
	}

	return nil
}

// UpsertHellosRequestMultiError is an error wrapping multiple validation
// errors returned by UpsertHellosRequest.ValidateAll() if the designated
// constraints aren't met.
type UpsertHellosRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertHellosRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertHellosRequestMultiError) AllErrors() []error { return m }

// UpsertHellosRequestValidationError is the validation error returned by
// UpsertHellosRequest.Validate if the designated constraints aren't met.
type UpsertHellosRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertHellosRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertHellosRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertHellosRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertHellosRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertHellosRequestValidationError) ErrorName() string {
	return "UpsertHellosRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertHellosRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertHellosRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertHellosRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertHellosRequestValidationError{}

// Validate checks the field values on UpsertTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpsertTodosRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpsertTodosRequestMultiError, or nil if none found.
func (m *UpsertTodosRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertTodosRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTodos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpsertTodosRequestValidationError{
						field:  fmt.Sprintf("Todos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpsertTodosRequestValidationError{
						field:  fmt.Sprintf("Todos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpsertTodosRequestValidationError{
					field:  fmt.Sprintf("Todos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpsertTodosRequestMultiError(errors)
	}

	return nil
}

// UpsertTodosRequestMultiError is an error wrapping multiple validation errors
// returned by UpsertTodosRequest.ValidateAll() if the designated constraints
// aren't met.
type UpsertTodosRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertTodosRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertTodosRequestMultiError) AllErrors() []error { return m }

// UpsertTodosRequestValidationError is the validation error returned by
// UpsertTodosRequest.Validate if the designated constraints aren't met.
type UpsertTodosRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertTodosRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertTodosRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertTodosRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertTodosRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertTodosRequestValidationError) ErrorName() string {
	return "UpsertTodosRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertTodosRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertTodosRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertTodosRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertTodosRequestValidationError{}

// Validate checks the field values on DeleteTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTodosRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTodosRequestMultiError, or nil if none found.
func (m *DeleteTodosRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTodosRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTodosRequestMultiError(errors)
	}

	return nil
}

// DeleteTodosRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTodosRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTodosRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTodosRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTodosRequestMultiError) AllErrors() []error { return m }

// DeleteTodosRequestValidationError is the validation error returned by
// DeleteTodosRequest.Validate if the designated constraints aren't met.
type DeleteTodosRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTodosRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTodosRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTodosRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTodosRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTodosRequestValidationError) ErrorName() string {
	return "DeleteTodosRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTodosRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTodosRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTodosRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTodosRequestValidationError{}

// Validate checks the field values on DeleteTodosResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTodosResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTodosResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTodosResponseMultiError, or nil if none found.
func (m *DeleteTodosResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTodosResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTodosResponseMultiError(errors)
	}

	return nil
}

// DeleteTodosResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTodosResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTodosResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTodosResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTodosResponseMultiError) AllErrors() []error { return m }

// DeleteTodosResponseValidationError is the validation error returned by
// DeleteTodosResponse.Validate if the designated constraints aren't met.
type DeleteTodosResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTodosResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTodosResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTodosResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTodosResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTodosResponseValidationError) ErrorName() string {
	return "DeleteTodosResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTodosResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTodosResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTodosResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTodosResponseValidationError{}

// Validate checks the field values on ListTodosRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTodosRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTodosRequestMultiError, or nil if none found.
func (m *ListTodosRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTodosRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Limit

	// no validation rules for Offset

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return ListTodosRequestMultiError(errors)
	}

	return nil
}

// ListTodosRequestMultiError is an error wrapping multiple validation errors
// returned by ListTodosRequest.ValidateAll() if the designated constraints
// aren't met.
type ListTodosRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTodosRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTodosRequestMultiError) AllErrors() []error { return m }

// ListTodosRequestValidationError is the validation error returned by
// ListTodosRequest.Validate if the designated constraints aren't met.
type ListTodosRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTodosRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTodosRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTodosRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTodosRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTodosRequestValidationError) ErrorName() string { return "ListTodosRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListTodosRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTodosRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTodosRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTodosRequestValidationError{}

// Validate checks the field values on GetTodosRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTodosRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTodosRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTodosRequestMultiError, or nil if none found.
func (m *GetTodosRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTodosRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTodosRequestMultiError(errors)
	}

	return nil
}

// GetTodosRequestMultiError is an error wrapping multiple validation errors
// returned by GetTodosRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTodosRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTodosRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTodosRequestMultiError) AllErrors() []error { return m }

// GetTodosRequestValidationError is the validation error returned by
// GetTodosRequest.Validate if the designated constraints aren't met.
type GetTodosRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTodosRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTodosRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTodosRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTodosRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTodosRequestValidationError) ErrorName() string { return "GetTodosRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTodosRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTodosRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTodosRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTodosRequestValidationError{}
