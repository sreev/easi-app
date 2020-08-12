// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Person person
// swagger:model Person
type Person struct {

	// common name
	CommonName string `json:"commonName,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// given name
	GivenName string `json:"givenName,omitempty"`

	// sur name
	SurName string `json:"surName,omitempty"`

	// telephone number
	TelephoneNumber string `json:"telephoneNumber,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this person
func (m *Person) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Person) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Person) UnmarshalBinary(b []byte) error {
	var res Person
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
