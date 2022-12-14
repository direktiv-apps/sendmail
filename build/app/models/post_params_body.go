// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// List of emails to send.
	// Required: true
	Emails []*PostParamsBodyEmailsItems `json:"emails"`

	// smtp
	// Required: true
	SMTP *PostParamsBodySMTP `json:"smtp"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) validateEmails(formats strfmt.Registry) error {

	if err := validate.Required("emails", "body", m.Emails); err != nil {
		return err
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostParamsBody) validateSMTP(formats strfmt.Registry) error {

	if err := validate.Required("smtp", "body", m.SMTP); err != nil {
		return err
	}

	if m.SMTP != nil {
		if err := m.SMTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post params body based on the context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emails); i++ {

		if m.Emails[i] != nil {
			if err := m.Emails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostParamsBody) contextValidateSMTP(ctx context.Context, formats strfmt.Registry) error {

	if m.SMTP != nil {
		if err := m.SMTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
