package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

type TupleParameters struct {
	Parameters []*Parameter `xml:"Parameter,omitempty"`
}

// Tuple Defines a tuple as a array of Parameter elements
// swagger:model Tuple
type Tuple struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp Tuple"`

	// id
	// Required: true
	ID *string `xml:"id"`

	// value
	// Required: true
	Value TupleParameters `xml:"value,omitempty"`
}

// Validate validates this tuple
func (m *Tuple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tuple) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Tuple) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	for i := 0; i < len(m.Value.Parameters); i++ {

		if swag.IsZero(m.Value.Parameters[i]) { // not required
			continue
		}

		if m.Value.Parameters[i] != nil {

			if err := m.Value.Parameters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}