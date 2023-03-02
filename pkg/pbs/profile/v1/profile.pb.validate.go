// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/pbs/profile/v1/profile.proto

package v1

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

// Validate checks the field values on ReadProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadProfileRequestMultiError, or nil if none found.
func (m *ReadProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 36 {
		err := ReadProfileRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return ReadProfileRequestMultiError(errors)
	}

	return nil
}

// ReadProfileRequestMultiError is an error wrapping multiple validation errors
// returned by ReadProfileRequest.ValidateAll() if the designated constraints
// aren't met.
type ReadProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadProfileRequestMultiError) AllErrors() []error { return m }

// ReadProfileRequestValidationError is the validation error returned by
// ReadProfileRequest.Validate if the designated constraints aren't met.
type ReadProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadProfileRequestValidationError) ErrorName() string {
	return "ReadProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadProfileRequestValidationError{}

// Validate checks the field values on ReadProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadProfileResponseMultiError, or nil if none found.
func (m *ReadProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for LegalName

	// no validation rules for UserName

	// no validation rules for Address

	// no validation rules for DateOfBirth

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadProfileResponseValidationError{
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
				errors = append(errors, ReadProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadProfileResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReadProfileResponseMultiError(errors)
	}

	return nil
}

// ReadProfileResponseMultiError is an error wrapping multiple validation
// errors returned by ReadProfileResponse.ValidateAll() if the designated
// constraints aren't met.
type ReadProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadProfileResponseMultiError) AllErrors() []error { return m }

// ReadProfileResponseValidationError is the validation error returned by
// ReadProfileResponse.Validate if the designated constraints aren't met.
type ReadProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadProfileResponseValidationError) ErrorName() string {
	return "ReadProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadProfileResponseValidationError{}

// Validate checks the field values on UpdateProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProfileRequestMultiError, or nil if none found.
func (m *UpdateProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 36 {
		err := UpdateProfileRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = UpdateProfileRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 50 {
		err := UpdateProfileRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetLegalName()); l < 3 || l > 100 {
		err := UpdateProfileRequestValidationError{
			field:  "LegalName",
			reason: "value length must be between 3 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUserName()); l < 3 || l > 20 {
		err := UpdateProfileRequestValidationError{
			field:  "UserName",
			reason: "value length must be between 3 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetAddress()); l < 3 || l > 250 {
		err := UpdateProfileRequestValidationError{
			field:  "Address",
			reason: "value length must be between 3 and 250 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDateOfBirth()); l < 9 || l > 150 {
		err := UpdateProfileRequestValidationError{
			field:  "DateOfBirth",
			reason: "value length must be between 9 and 150 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateProfileRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateProfileRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateProfileRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UpdateProfileRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProfileRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProfileRequestMultiError) AllErrors() []error { return m }

// UpdateProfileRequestValidationError is the validation error returned by
// UpdateProfileRequest.Validate if the designated constraints aren't met.
type UpdateProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProfileRequestValidationError) ErrorName() string {
	return "UpdateProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProfileRequestValidationError{}

// Validate checks the field values on UpdateProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProfileResponseMultiError, or nil if none found.
func (m *UpdateProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for LegalName

	// no validation rules for UserName

	// no validation rules for Address

	// no validation rules for DateOfBirth

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProfileResponseValidationError{
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
				errors = append(errors, UpdateProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProfileResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateProfileResponseMultiError(errors)
	}

	return nil
}

// UpdateProfileResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateProfileResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProfileResponseMultiError) AllErrors() []error { return m }

// UpdateProfileResponseValidationError is the validation error returned by
// UpdateProfileResponse.Validate if the designated constraints aren't met.
type UpdateProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProfileResponseValidationError) ErrorName() string {
	return "UpdateProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProfileResponseValidationError{}

// Validate checks the field values on CreateProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProfileRequestMultiError, or nil if none found.
func (m *CreateProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 36 {
		err := CreateProfileRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = CreateProfileRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 50 {
		err := CreateProfileRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetLegalName()); l < 3 || l > 100 {
		err := CreateProfileRequestValidationError{
			field:  "LegalName",
			reason: "value length must be between 3 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUserName()); l < 3 || l > 20 {
		err := CreateProfileRequestValidationError{
			field:  "UserName",
			reason: "value length must be between 3 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetAddress()); l < 3 || l > 250 {
		err := CreateProfileRequestValidationError{
			field:  "Address",
			reason: "value length must be between 3 and 250 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDateOfBirth()); l < 9 || l > 150 {
		err := CreateProfileRequestValidationError{
			field:  "DateOfBirth",
			reason: "value length must be between 9 and 150 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateProfileRequestMultiError(errors)
	}

	return nil
}

func (m *CreateProfileRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateProfileRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateProfileRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProfileRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProfileRequestMultiError) AllErrors() []error { return m }

// CreateProfileRequestValidationError is the validation error returned by
// CreateProfileRequest.Validate if the designated constraints aren't met.
type CreateProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProfileRequestValidationError) ErrorName() string {
	return "CreateProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProfileRequestValidationError{}

// Validate checks the field values on CreateProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProfileResponseMultiError, or nil if none found.
func (m *CreateProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for LegalName

	// no validation rules for UserName

	// no validation rules for Address

	// no validation rules for DateOfBirth

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProfileResponseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProfileResponseValidationError{
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
				errors = append(errors, CreateProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProfileResponseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProfileResponseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateProfileResponseMultiError(errors)
	}

	return nil
}

// CreateProfileResponseMultiError is an error wrapping multiple validation
// errors returned by CreateProfileResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProfileResponseMultiError) AllErrors() []error { return m }

// CreateProfileResponseValidationError is the validation error returned by
// CreateProfileResponse.Validate if the designated constraints aren't met.
type CreateProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProfileResponseValidationError) ErrorName() string {
	return "CreateProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProfileResponseValidationError{}
