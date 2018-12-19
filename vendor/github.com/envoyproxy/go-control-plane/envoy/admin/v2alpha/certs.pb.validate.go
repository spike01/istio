// Code generated by protoc-gen-validate
// source: envoy/admin/v2alpha/certs.proto
// DO NOT EDIT!!!

package envoy_admin_v2alpha

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Certificates with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Certificates) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCertificates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificatesValidationError{
					Field:  fmt.Sprintf("Certificates[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// CertificatesValidationError is the validation error returned by
// Certificates.Validate if the designated constraints aren't met.
type CertificatesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CertificatesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificates.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CertificatesValidationError{}

// Validate checks the field values on Certificate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Certificate) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCaCert() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateValidationError{
					Field:  fmt.Sprintf("CaCert[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetCertChain() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateValidationError{
					Field:  fmt.Sprintf("CertChain[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// CertificateValidationError is the validation error returned by
// Certificate.Validate if the designated constraints aren't met.
type CertificateValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CertificateValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificate.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CertificateValidationError{}

// Validate checks the field values on CertificateDetails with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CertificateDetails) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Path

	// no validation rules for SerialNumber

	for idx, item := range m.GetSubjectAltNames() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateDetailsValidationError{
					Field:  fmt.Sprintf("SubjectAltNames[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for DaysUntilExpiration

	if v, ok := interface{}(m.GetValidFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetailsValidationError{
				Field:  "ValidFrom",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetExpirationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetailsValidationError{
				Field:  "ExpirationTime",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// CertificateDetailsValidationError is the validation error returned by
// CertificateDetails.Validate if the designated constraints aren't met.
type CertificateDetailsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CertificateDetailsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateDetails.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CertificateDetailsValidationError{}

// Validate checks the field values on SubjectAlternateName with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubjectAlternateName) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Name.(type) {

	case *SubjectAlternateName_Dns:
		// no validation rules for Dns

	case *SubjectAlternateName_Uri:
		// no validation rules for Uri

	}

	return nil
}

// SubjectAlternateNameValidationError is the validation error returned by
// SubjectAlternateName.Validate if the designated constraints aren't met.
type SubjectAlternateNameValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e SubjectAlternateNameValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectAlternateName.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = SubjectAlternateNameValidationError{}