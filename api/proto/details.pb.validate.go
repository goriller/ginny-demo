// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: details.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetDetailRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDetailRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 999 {
		return GetDetailRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 999",
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 3 {
		return GetDetailRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// GetDetailRequestValidationError is the validation error returned by
// GetDetailRequest.Validate if the designated constraints aren't met.
type GetDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDetailRequestValidationError) ErrorName() string { return "GetDetailRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDetailRequestValidationError{}

// Validate checks the field values on Detail with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Detail) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Price

	if v, ok := interface{}(m.GetCreatedTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DetailValidationError{
				field:  "CreatedTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DetailValidationError is the validation error returned by Detail.Validate if
// the designated constraints aren't met.
type DetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailValidationError) ErrorName() string { return "DetailValidationError" }

// Error satisfies the builtin error interface
func (e DetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailValidationError{}