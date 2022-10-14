// Code generated by go-swagger; DO NOT EDIT.

package promos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelNewsUpdated model news updated
//
// swagger:model model.NewsUpdated
type PromoCreated struct {

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this model news updated
func (m *PromoCreated) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model news updated based on context it is used
func (m *PromoCreated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PromoCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromoCreated) UnmarshalBinary(b []byte) error {
	var res PromoCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
