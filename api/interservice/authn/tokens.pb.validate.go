// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: interservice/authn/tokens.proto

package authn

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

// Validate checks the field values on CreateTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateTokenReqMultiError,
// or nil if none found.
func (m *CreateTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_CreateTokenReq_Id_Pattern.MatchString(m.GetId()) {
		err := CreateTokenReqValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[a-z0-9-_]{1,64}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Active

	_CreateTokenReq_Projects_Unique := make(map[string]struct{}, len(m.GetProjects()))

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if _, exists := _CreateTokenReq_Projects_Unique[item]; exists {
			err := CreateTokenReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_CreateTokenReq_Projects_Unique[item] = struct{}{}
		}

		if !_CreateTokenReq_Projects_Pattern.MatchString(item) {
			err := CreateTokenReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "value does not match regex pattern \"^[a-z0-9-_]{1,64}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CreateTokenReqMultiError(errors)
	}

	return nil
}

// CreateTokenReqMultiError is an error wrapping multiple validation errors
// returned by CreateTokenReq.ValidateAll() if the designated constraints
// aren't met.
type CreateTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenReqMultiError) AllErrors() []error { return m }

// CreateTokenReqValidationError is the validation error returned by
// CreateTokenReq.Validate if the designated constraints aren't met.
type CreateTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenReqValidationError) ErrorName() string { return "CreateTokenReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenReqValidationError{}

var _CreateTokenReq_Id_Pattern = regexp.MustCompile("^[a-z0-9-_]{1,64}$")

var _CreateTokenReq_Projects_Pattern = regexp.MustCompile("^[a-z0-9-_]{1,64}$")

// Validate checks the field values on CreateTokenWithValueReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTokenWithValueReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenWithValueReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTokenWithValueReqMultiError, or nil if none found.
func (m *CreateTokenWithValueReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenWithValueReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Active

	// no validation rules for Value

	_CreateTokenWithValueReq_Projects_Unique := make(map[string]struct{}, len(m.GetProjects()))

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if _, exists := _CreateTokenWithValueReq_Projects_Unique[item]; exists {
			err := CreateTokenWithValueReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_CreateTokenWithValueReq_Projects_Unique[item] = struct{}{}
		}

		if !_CreateTokenWithValueReq_Projects_Pattern.MatchString(item) {
			err := CreateTokenWithValueReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "value does not match regex pattern \"^[a-z0-9-_]{1,64}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CreateTokenWithValueReqMultiError(errors)
	}

	return nil
}

// CreateTokenWithValueReqMultiError is an error wrapping multiple validation
// errors returned by CreateTokenWithValueReq.ValidateAll() if the designated
// constraints aren't met.
type CreateTokenWithValueReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenWithValueReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenWithValueReqMultiError) AllErrors() []error { return m }

// CreateTokenWithValueReqValidationError is the validation error returned by
// CreateTokenWithValueReq.Validate if the designated constraints aren't met.
type CreateTokenWithValueReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenWithValueReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenWithValueReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenWithValueReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenWithValueReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenWithValueReqValidationError) ErrorName() string {
	return "CreateTokenWithValueReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTokenWithValueReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenWithValueReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenWithValueReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenWithValueReqValidationError{}

var _CreateTokenWithValueReq_Projects_Pattern = regexp.MustCompile("^[a-z0-9-_]{1,64}$")

// Validate checks the field values on UpdateTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTokenReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateTokenReqMultiError,
// or nil if none found.
func (m *UpdateTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Active

	// no validation rules for Name

	_UpdateTokenReq_Projects_Unique := make(map[string]struct{}, len(m.GetProjects()))

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if _, exists := _UpdateTokenReq_Projects_Unique[item]; exists {
			err := UpdateTokenReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_UpdateTokenReq_Projects_Unique[item] = struct{}{}
		}

		if !_UpdateTokenReq_Projects_Pattern.MatchString(item) {
			err := UpdateTokenReqValidationError{
				field:  fmt.Sprintf("Projects[%v]", idx),
				reason: "value does not match regex pattern \"^[a-z0-9-_]{1,64}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateTokenReqMultiError(errors)
	}

	return nil
}

// UpdateTokenReqMultiError is an error wrapping multiple validation errors
// returned by UpdateTokenReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTokenReqMultiError) AllErrors() []error { return m }

// UpdateTokenReqValidationError is the validation error returned by
// UpdateTokenReq.Validate if the designated constraints aren't met.
type UpdateTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTokenReqValidationError) ErrorName() string { return "UpdateTokenReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTokenReqValidationError{}

var _UpdateTokenReq_Projects_Pattern = regexp.MustCompile("^[a-z0-9-_]{1,64}$")

// Validate checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Token) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokenMultiError, or nil if none found.
func (m *Token) ValidateAll() error {
	return m.validate(true)
}

func (m *Token) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Value

	// no validation rules for Active

	// no validation rules for Created

	// no validation rules for Updated

	if len(errors) > 0 {
		return TokenMultiError(errors)
	}

	return nil
}

// TokenMultiError is an error wrapping multiple validation errors returned by
// Token.ValidateAll() if the designated constraints aren't met.
type TokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenMultiError) AllErrors() []error { return m }

// TokenValidationError is the validation error returned by Token.Validate if
// the designated constraints aren't met.
type TokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenValidationError) ErrorName() string { return "TokenValidationError" }

// Error satisfies the builtin error interface
func (e TokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenValidationError{}

// Validate checks the field values on Tokens with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tokens) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tokens with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokensMultiError, or nil if none found.
func (m *Tokens) ValidateAll() error {
	return m.validate(true)
}

func (m *Tokens) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTokens() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TokensValidationError{
						field:  fmt.Sprintf("Tokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TokensValidationError{
						field:  fmt.Sprintf("Tokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TokensValidationError{
					field:  fmt.Sprintf("Tokens[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TokensMultiError(errors)
	}

	return nil
}

// TokensMultiError is an error wrapping multiple validation errors returned by
// Tokens.ValidateAll() if the designated constraints aren't met.
type TokensMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokensMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokensMultiError) AllErrors() []error { return m }

// TokensValidationError is the validation error returned by Tokens.Validate if
// the designated constraints aren't met.
type TokensValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokensValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokensValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokensValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokensValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokensValidationError) ErrorName() string { return "TokensValidationError" }

// Error satisfies the builtin error interface
func (e TokensValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokens.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokensValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokensValidationError{}

// Validate checks the field values on Value with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Value with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ValueMultiError, or nil if none found.
func (m *Value) ValidateAll() error {
	return m.validate(true)
}

func (m *Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return ValueMultiError(errors)
	}

	return nil
}

// ValueMultiError is an error wrapping multiple validation errors returned by
// Value.ValidateAll() if the designated constraints aren't met.
type ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValueMultiError) AllErrors() []error { return m }

// ValueValidationError is the validation error returned by Value.Validate if
// the designated constraints aren't met.
type ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValueValidationError) ErrorName() string { return "ValueValidationError" }

// Error satisfies the builtin error interface
func (e ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValueValidationError{}

// Validate checks the field values on GetTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTokenReqMultiError, or
// nil if none found.
func (m *GetTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetTokenReqMultiError(errors)
	}

	return nil
}

// GetTokenReqMultiError is an error wrapping multiple validation errors
// returned by GetTokenReq.ValidateAll() if the designated constraints aren't met.
type GetTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTokenReqMultiError) AllErrors() []error { return m }

// GetTokenReqValidationError is the validation error returned by
// GetTokenReq.Validate if the designated constraints aren't met.
type GetTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTokenReqValidationError) ErrorName() string { return "GetTokenReqValidationError" }

// Error satisfies the builtin error interface
func (e GetTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTokenReqValidationError{}

// Validate checks the field values on GetTokensReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTokensReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTokensReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTokensReqMultiError, or
// nil if none found.
func (m *GetTokensReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTokensReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTokensReqMultiError(errors)
	}

	return nil
}

// GetTokensReqMultiError is an error wrapping multiple validation errors
// returned by GetTokensReq.ValidateAll() if the designated constraints aren't met.
type GetTokensReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTokensReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTokensReqMultiError) AllErrors() []error { return m }

// GetTokensReqValidationError is the validation error returned by
// GetTokensReq.Validate if the designated constraints aren't met.
type GetTokensReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTokensReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTokensReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTokensReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTokensReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTokensReqValidationError) ErrorName() string { return "GetTokensReqValidationError" }

// Error satisfies the builtin error interface
func (e GetTokensReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTokensReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTokensReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTokensReqValidationError{}

// Validate checks the field values on DeleteTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTokenReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteTokenReqMultiError,
// or nil if none found.
func (m *DeleteTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTokenReqMultiError(errors)
	}

	return nil
}

// DeleteTokenReqMultiError is an error wrapping multiple validation errors
// returned by DeleteTokenReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTokenReqMultiError) AllErrors() []error { return m }

// DeleteTokenReqValidationError is the validation error returned by
// DeleteTokenReq.Validate if the designated constraints aren't met.
type DeleteTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTokenReqValidationError) ErrorName() string { return "DeleteTokenReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTokenReqValidationError{}

// Validate checks the field values on DeleteTokenResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTokenResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTokenResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTokenRespMultiError, or nil if none found.
func (m *DeleteTokenResp) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTokenResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTokenRespMultiError(errors)
	}

	return nil
}

// DeleteTokenRespMultiError is an error wrapping multiple validation errors
// returned by DeleteTokenResp.ValidateAll() if the designated constraints
// aren't met.
type DeleteTokenRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTokenRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTokenRespMultiError) AllErrors() []error { return m }

// DeleteTokenRespValidationError is the validation error returned by
// DeleteTokenResp.Validate if the designated constraints aren't met.
type DeleteTokenRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTokenRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTokenRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTokenRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTokenRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTokenRespValidationError) ErrorName() string { return "DeleteTokenRespValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTokenRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTokenResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTokenRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTokenRespValidationError{}
